package openapi

import (
	"encoding/json"
	"testing"
)

// The Camunda Console API ships additive changes without bumping the OpenAPI
// `info.version`, and its spec declares response objects closed even though the
// server sends fields the spec omits. The generated client used to call
// decoder.DisallowUnknownFields() on every model with a required property, so each
// new server-side field became a hard decode error.
//
// These payloads are shaped after what the live API actually returns. See
// openapi-normalize.jq for the spec corrections that keep them decodable.

const clusterPayload = `{
  "uuid": "1b0d1e0d-1111-2222-3333-444455556666",
  "name": "test-cluster",
  "ownerId": "owner-1",
  "created": "2026-01-15T10:00:00.000Z",
  "autoUpdate": true,
  "channel": {"uuid": "ch-1", "name": "Stable"},
  "generation": {"uuid": "gen-1", "name": "8.6.0"},
  "region": {"uuid": "reg-1", "name": "Europe West 1"},
  "planType": {"uuid": "plan-1", "name": "Professional"},
  "links": {"zeebe": "grpc://zeebe", "operate": "https://operate"},
  "status": {"ready": "Healthy", "zeebeStatus": "Healthy", "connectorsStatus": "Healthy"},
  "encryption": {"type": "Software", "status": "Ready"}
}`

const createdClusterClientPayload = `{
  "uuid": "client-uuid-1",
  "name": "my-client",
  "clientId": "abc123",
  "clientSecret": "s3cr3t",
  "permissions": ["Zeebe"],
  "audience": "zeebe.camunda.io"
}`

const backupPayload = `{
  "uuid": "backup-1",
  "name": "nightly",
  "created": "2026-01-15T10:00:00.000Z",
  "completed": "2026-01-15T10:05:00.000Z",
  "status": "Complete",
  "zeebeStatus": "Complete",
  "tasklistStatus": "Complete",
  "operateStatus": "Complete",
  "optimizeStatus": "Complete",
  "generationUuid": "gen-1",
  "generationName": "8.6.0"
}`

const clusterClientPayload = `{
  "clientId": "abc123",
  "name": "my-client",
  "permissions": ["Zeebe"]
}`

const parametersPayload = `{
  "channels": [],
  "clusterPlanTypes": [],
  "regions": []
}`

func TestDecodeLivePayloads(t *testing.T) {
	tests := []struct {
		name    string
		payload string
		target  interface{}
	}{
		{"Cluster", clusterPayload, &Cluster{}},
		{"CreatedClusterClient", createdClusterClientPayload, &CreatedClusterClient{}},
		{"BackupDto", backupPayload, &BackupDto{}},
		{"ClusterClient", clusterClientPayload, &ClusterClient{}},
		{"Parameters", parametersPayload, &Parameters{}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if err := json.Unmarshal([]byte(tc.payload), tc.target); err != nil {
				t.Fatalf("cannot decode live %s response: %v", tc.name, err)
			}
		})
	}
}

// The fields that were missing from the vendored spec must survive decoding, not
// just be tolerated -- terraform-provider-camunda needs to read them.
func TestNewFieldsAreDecoded(t *testing.T) {
	var c Cluster
	if err := json.Unmarshal([]byte(clusterPayload), &c); err != nil {
		t.Fatalf("decode Cluster: %v", err)
	}
	if c.Encryption == nil {
		t.Error("Cluster.Encryption was dropped")
	} else if got := c.Encryption.Type; got != CLUSTERENCRYPTIONKEY_SOFTWARE {
		t.Errorf("Cluster.Encryption.Type = %q, want Software", got)
	}

	var cc CreatedClusterClient
	if err := json.Unmarshal([]byte(createdClusterClientPayload), &cc); err != nil {
		t.Fatalf("decode CreatedClusterClient: %v", err)
	}
	if cc.Audience == nil || *cc.Audience != "zeebe.camunda.io" {
		t.Error("CreatedClusterClient.Audience was dropped")
	}
}

// `connectorsStatus` is a string enum like its sibling component statuses. The old
// spec patch rewrote the sibling $refs by line context and `patch` applied it with
// fuzz, silently leaving connectorsStatus pointing at the status *object* -- so a
// cluster running connectors failed to decode.
func TestConnectorsStatusIsAnEnumNotAnObject(t *testing.T) {
	var s ClusterStatus
	payload := `{"ready":"Healthy","zeebeStatus":"Healthy","connectorsStatus":"Healthy"}`
	if err := json.Unmarshal([]byte(payload), &s); err != nil {
		t.Fatalf("cannot decode ClusterStatus with connectorsStatus: %v", err)
	}
	if s.ConnectorsStatus == nil || *s.ConnectorsStatus != CLUSTERCOMPONENTSTATUS_HEALTHY {
		t.Errorf("ConnectorsStatus = %v, want Healthy", s.ConnectorsStatus)
	}
}

// The guarantee that keeps this from happening again: a field Camunda has not
// invented yet must not break decoding, and must round-trip rather than vanish.
func TestUnknownFutureFieldsDoNotBreakDecoding(t *testing.T) {
	payload := `{
	  "uuid": "client-uuid-1",
	  "name": "my-client",
	  "clientId": "abc123",
	  "clientSecret": "s3cr3t",
	  "permissions": ["Zeebe"],
	  "someFieldCamundaAddsNextQuarter": {"nested": true}
	}`

	var cc CreatedClusterClient
	if err := json.Unmarshal([]byte(payload), &cc); err != nil {
		t.Fatalf("an unknown field must not break decoding: %v", err)
	}
	if _, ok := cc.AdditionalProperties["someFieldCamundaAddsNextQuarter"]; !ok {
		t.Error("unknown field was silently dropped instead of captured in AdditionalProperties")
	}

	out, err := json.Marshal(&cc)
	if err != nil {
		t.Fatalf("re-encode: %v", err)
	}
	var round map[string]interface{}
	if err := json.Unmarshal(out, &round); err != nil {
		t.Fatalf("re-decode: %v", err)
	}
	if _, ok := round["someFieldCamundaAddsNextQuarter"]; !ok {
		t.Error("unknown field did not survive a decode/encode round trip")
	}
}
