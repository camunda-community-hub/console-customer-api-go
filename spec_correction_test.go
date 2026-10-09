package openapi

import (
	"encoding/json"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// Each spec correction in openapi-normalize.jq must hold in two directions:
//
//   - applied: the corrected spec (openapi.json) no longer has the defect.
//   - still needed: the upstream spec (openapi.upstream.json) still has it. When
//     Camunda fixes a defect upstream, the correction becomes dead code that
//     silently no-ops; this fails the weekly update PR instead, so it gets deleted.
//
// The drift check in CI guarantees openapi.json is exactly the corrections applied
// to openapi.upstream.json, so testing the committed files tests the correction.

const clusterStatusRef = "#/components/schemas/ClusterStatus"

var specCorrections = []struct {
	name        string
	stillNeeded func(upstream map[string]interface{}) bool
	applied     func(corrected map[string]interface{}) bool
}{
	{
		name: "rename ClusterStatus enum to ClusterComponentStatus",
		stillNeeded: func(s map[string]interface{}) bool {
			enum, ok := schemas(s)["ClusterStatus"].(map[string]interface{})
			return ok && enum["type"] == "string" && enum["enum"] != nil
		},
		applied: func(s map[string]interface{}) bool {
			_, hasOld := schemas(s)["ClusterStatus"]
			_, hasNew := schemas(s)["ClusterComponentStatus"]
			return !hasOld && hasNew && countRefs(s, clusterStatusRef) == 0
		},
	},
	{
		name: "open closed objects",
		stillNeeded: func(s map[string]interface{}) bool {
			return countClosedObjects(s) > 0
		},
		applied: func(s map[string]interface{}) bool {
			return countClosedObjects(s) == 0
		},
	},
}

func TestSpecCorrections(t *testing.T) {
	upstream := loadSpec(t, "openapi.upstream.json")
	corrected := loadSpec(t, "openapi.json")

	for _, c := range specCorrections {
		t.Run(c.name, func(t *testing.T) {
			if !c.stillNeeded(upstream) {
				t.Error("upstream spec no longer has this defect: delete the correction from openapi-normalize.jq")
			}
			if !c.applied(corrected) {
				t.Error("corrected spec still has this defect")
			}
		})
	}
}

// Spec correction 3 is a standing rule rather than a one-off fix, so it has no
// "still needed" side: in response-only schemas, the corrected spec requires
// exactly the fields in required-baseline.json, and every baseline entry must
// still be required upstream.
func TestRequiredBaseline(t *testing.T) {
	upstream := loadSpec(t, "openapi.upstream.json")
	corrected := loadSpec(t, "openapi.json")

	var baseline map[string][]string
	raw, err := os.ReadFile("required-baseline.json")
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(raw, &baseline); err != nil {
		t.Fatalf("decode required-baseline.json: %v", err)
	}

	t.Run("corrected spec requires exactly the baseline", func(t *testing.T) {
		got := requiredLocations(corrected)
		for key, fields := range got {
			if !sameFields(fields, baseline[key]) {
				t.Errorf("%s: required %v, baseline %v", key, fields, baseline[key])
			}
		}
		for key, fields := range baseline {
			if _, ok := got[key]; !ok {
				t.Errorf("%s: baseline %v, but nothing required", key, fields)
			}
		}
	})

	t.Run("baseline is still required upstream", func(t *testing.T) {
		up := requiredLocations(upstream)
		for key, fields := range baseline {
			for _, f := range fields {
				if !hasField(up[key], f) {
					t.Errorf("%s: upstream no longer requires %q: remove it from required-baseline.json", key, f)
				}
			}
		}
	})
}

func loadSpec(t *testing.T, path string) map[string]interface{} {
	t.Helper()
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	var spec map[string]interface{}
	if err := json.Unmarshal(raw, &spec); err != nil {
		t.Fatalf("decode %s: %v", path, err)
	}
	return spec
}

func schemas(spec map[string]interface{}) map[string]interface{} {
	components, _ := spec["components"].(map[string]interface{})
	s, _ := components["schemas"].(map[string]interface{})
	return s
}

func countClosedObjects(node interface{}) int {
	return countMatching(node, func(o map[string]interface{}) bool {
		return o["additionalProperties"] == false
	})
}

func countRefs(node interface{}, ref string) int {
	return countMatching(node, func(o map[string]interface{}) bool {
		return o["$ref"] == ref
	})
}

// countMatching walks the whole document, like jq's walk/1.
func countMatching(node interface{}, match func(map[string]interface{}) bool) int {
	n := 0
	switch v := node.(type) {
	case map[string]interface{}:
		if match(v) {
			n++
		}
		for _, child := range v {
			n += countMatching(child, match)
		}
	case []interface{}:
		for _, child := range v {
			n += countMatching(child, match)
		}
	}
	return n
}

// requiredLocations mirrors required_locations in openapi-normalize.jq: the
// required list of every object inside a response-only schema, keyed by its path
// within components.schemas joined by "/".
func requiredLocations(spec map[string]interface{}) map[string][]string {
	out := map[string][]string{}
	var walk func(key string, node interface{})
	walk = func(key string, node interface{}) {
		switch v := node.(type) {
		case map[string]interface{}:
			if required, ok := v["required"].([]interface{}); ok {
				for _, r := range required {
					out[key] = append(out[key], r.(string))
				}
			}
			for k, child := range v {
				walk(key+"/"+k, child)
			}
		case []interface{}:
			for i, child := range v {
				walk(key+"/"+strconv.Itoa(i), child)
			}
		}
	}
	for _, name := range responseOnlySchemas(spec) {
		walk(name, schemas(spec)[name])
	}
	return out
}

// responseOnlySchemas are the schemas reachable from a response but not from a
// request body.
func responseOnlySchemas(spec map[string]interface{}) []string {
	var responseRoots, requestRoots []string
	paths, _ := spec["paths"].(map[string]interface{})
	for _, item := range paths {
		ops, _ := item.(map[string]interface{})
		for _, op := range ops {
			op, ok := op.(map[string]interface{})
			if !ok {
				continue
			}
			responseRoots = append(responseRoots, schemaRefs(op["responses"])...)
			requestRoots = append(requestRoots, schemaRefs(op["requestBody"])...)
		}
	}
	responses := reachableSchemas(spec, responseRoots)
	requests := reachableSchemas(spec, requestRoots)
	var only []string
	for name := range responses {
		if !requests[name] {
			only = append(only, name)
		}
	}
	sort.Strings(only)
	return only
}

func reachableSchemas(spec map[string]interface{}, roots []string) map[string]bool {
	seen := map[string]bool{}
	todo := roots
	for len(todo) > 0 {
		name := todo[0]
		todo = todo[1:]
		if seen[name] {
			continue
		}
		seen[name] = true
		todo = append(todo, schemaRefs(schemas(spec)[name])...)
	}
	return seen
}

func schemaRefs(node interface{}) []string {
	var refs []string
	switch v := node.(type) {
	case map[string]interface{}:
		if ref, ok := v["$ref"].(string); ok {
			refs = append(refs, strings.TrimPrefix(ref, "#/components/schemas/"))
		}
		for _, child := range v {
			refs = append(refs, schemaRefs(child)...)
		}
	case []interface{}:
		for _, child := range v {
			refs = append(refs, schemaRefs(child)...)
		}
	}
	return refs
}

func sameFields(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for _, f := range a {
		if !hasField(b, f) {
			return false
		}
	}
	return true
}

// hasField is case-sensitive, unlike the generated contains in client.go.
func hasField(fields []string, field string) bool {
	for _, f := range fields {
		if f == field {
			return true
		}
	}
	return false
}
