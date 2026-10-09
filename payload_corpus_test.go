package openapi

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"
)

// The Camunda Console API ships additive changes without bumping the OpenAPI
// `info.version`, and its spec declares response objects closed even though the
// server sends fields the spec omits. The generated client used to call
// decoder.DisallowUnknownFields() on every model with a required property, so each
// new server-side field became a hard decode error.
//
// testdata/payloads/<Model>/<case>.json holds response bodies shaped after what the
// live API returns. Every one must decode into its model and survive a re-encode
// with nothing dropped or changed. To cover a new response, or reproduce a decode
// bug, add a file. See openapi-normalize.jq for the spec corrections that keep
// these decodable.

const payloadCorpus = "testdata/payloads"

// responseModels maps a corpus directory to the model its payloads decode into:
// every struct model the corrected spec returns from a 200 response.
var responseModels = map[string]func() interface{}{
	"AuditDto":                       func() interface{} { return &AuditDto{} },
	"BackupDto":                      func() interface{} { return &BackupDto{} },
	"BackupScheduleDto":              func() interface{} { return &BackupScheduleDto{} },
	"Cluster":                        func() interface{} { return &Cluster{} },
	"ClusterClient":                  func() interface{} { return &ClusterClient{} },
	"ClusterClientConnectionDetails": func() interface{} { return &ClusterClientConnectionDetails{} },
	"CreatedClusterClient":           func() interface{} { return &CreatedClusterClient{} },
	"FailoverResult":                 func() interface{} { return &FailoverResult{} },
	"GenerationUpgradeForClusterDto": func() interface{} { return &GenerationUpgradeForClusterDto{} },
	"Member":                         func() interface{} { return &Member{} },
	"MetaDto":                        func() interface{} { return &MetaDto{} },
	"Parameters":                     func() interface{} { return &Parameters{} },
	"RestoreDto":                     func() interface{} { return &RestoreDto{} },
}

func TestPayloadCorpus(t *testing.T) {
	dirs, err := os.ReadDir(payloadCorpus)
	if err != nil {
		t.Fatalf("read corpus: %v", err)
	}
	for _, dir := range dirs {
		newModel, ok := responseModels[dir.Name()]
		if !ok {
			t.Errorf("%s/%s: no model registered in responseModels", payloadCorpus, dir.Name())
			continue
		}
		files, err := filepath.Glob(filepath.Join(payloadCorpus, dir.Name(), "*.json"))
		if err != nil {
			t.Fatal(err)
		}
		for _, file := range files {
			t.Run(dir.Name()+"/"+filepath.Base(file), func(t *testing.T) {
				raw, err := os.ReadFile(file)
				if err != nil {
					t.Fatal(err)
				}
				model := newModel()
				if err := json.Unmarshal(raw, model); err != nil {
					t.Fatalf("cannot decode: %v", err)
				}
				out, err := json.Marshal(model)
				if err != nil {
					t.Fatalf("re-encode: %v", err)
				}
				var in, round interface{}
				if err := json.Unmarshal(raw, &in); err != nil {
					t.Fatal(err)
				}
				if err := json.Unmarshal(out, &round); err != nil {
					t.Fatal(err)
				}
				assertPreserved(t, "$", in, round)
			})
		}
	}
}

// assertPreserved fails for every value in want that is missing from or different
// in got. got may carry extra keys (defaults the model fills in). Timestamps
// compare as instants, since time.Time re-encodes "…00.000Z" as "…00Z".
func assertPreserved(t *testing.T, path string, want, got interface{}) {
	t.Helper()
	switch w := want.(type) {
	case map[string]interface{}:
		g, ok := got.(map[string]interface{})
		if !ok {
			t.Errorf("%s: want object, got %v", path, got)
			return
		}
		for k, wv := range w {
			gv, ok := g[k]
			if !ok {
				t.Errorf("%s.%s: dropped", path, k)
				continue
			}
			assertPreserved(t, path+"."+k, wv, gv)
		}
	case string:
		if g, ok := got.(string); ok && sameInstant(w, g) {
			return
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("%s: want %v, got %v", path, want, got)
		}
	default:
		if !reflect.DeepEqual(want, got) {
			t.Errorf("%s: want %v, got %v", path, want, got)
		}
	}
}

func sameInstant(a, b string) bool {
	ta, errA := time.Parse(time.RFC3339Nano, a)
	tb, errB := time.Parse(time.RFC3339Nano, b)
	return errA == nil && errB == nil && ta.Equal(tb)
}

func decodePayload(t *testing.T, file string, model interface{}) {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join(payloadCorpus, file))
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(raw, model); err != nil {
		t.Fatalf("decode %s: %v", file, err)
	}
}

// The corpus round trip passes even when a field only survives in
// AdditionalProperties. terraform-provider-camunda reads these as typed fields,
// so check they land there.
func TestNewFieldsAreDecoded(t *testing.T) {
	var c Cluster
	decodePayload(t, "Cluster/encryption-and-connectors.json", &c)
	if c.Encryption == nil {
		t.Error("Cluster.Encryption was dropped")
	} else if got := c.Encryption.Type; got != CLUSTERENCRYPTIONKEY_SOFTWARE {
		t.Errorf("Cluster.Encryption.Type = %q, want Software", got)
	}

	var cc CreatedClusterClient
	decodePayload(t, "CreatedClusterClient/with-audience.json", &cc)
	if cc.Audience == nil || *cc.Audience != "zeebe.camunda.io" {
		t.Error("CreatedClusterClient.Audience was dropped")
	}
}

// `connectorsStatus` is a string enum like its sibling component statuses. The old
// spec patch rewrote the sibling $refs by line context and `patch` applied it with
// fuzz, silently leaving connectorsStatus pointing at the status *object* -- so a
// cluster running connectors failed to decode.
func TestConnectorsStatusIsAnEnumNotAnObject(t *testing.T) {
	var c Cluster
	decodePayload(t, "Cluster/encryption-and-connectors.json", &c)
	if s := c.Status.ConnectorsStatus; s == nil || *s != CLUSTERCOMPONENTSTATUS_HEALTHY {
		t.Errorf("ConnectorsStatus = %v, want Healthy", s)
	}
}
