package openapi

import (
	"encoding/json"
	"os"
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
	{
		name:        "relax Cluster.encryption",
		stillNeeded: func(s map[string]interface{}) bool { return isRequired(s, "Cluster", "encryption") },
		applied:     func(s map[string]interface{}) bool { return !isRequired(s, "Cluster", "encryption") },
	},
	{
		name:        "relax CreatedClusterClient.audience",
		stillNeeded: func(s map[string]interface{}) bool { return isRequired(s, "CreatedClusterClient", "audience") },
		applied:     func(s map[string]interface{}) bool { return !isRequired(s, "CreatedClusterClient", "audience") },
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

func isRequired(spec map[string]interface{}, schema, field string) bool {
	s, _ := schemas(spec)[schema].(map[string]interface{})
	required, _ := s["required"].([]interface{})
	for _, r := range required {
		if r == field {
			return true
		}
	}
	return false
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
