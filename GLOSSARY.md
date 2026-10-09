# Glossary

**Upstream spec**: Camunda's published OpenAPI spec for the Console API, stored in
`openapi.upstream.json` exactly as fetched (keys sorted, nothing else changed). Only
`make fetch` writes it.

**Corrected spec**: the generator's input, `openapi.json`. Always derived from the
upstream spec by applying the spec corrections; never edited by hand.

**Spec correction**: a deliberate, documented deviation from the upstream spec,
made because the published spec does not match what the API actually serves. Each
lives in `openapi-normalize.jq` with its rationale, and `spec_correction_test.go`
checks it is applied and, for one-off fixes, still needed.

**Required baseline**: the fields we accept as required in response schemas,
listed in `required-baseline.json` by location (`Cluster`,
`Cluster/properties/channel`). Any other field upstream requires is relaxed, so a
newly required field can't break decoding. Promote a field by adding it; an entry
upstream no longer requires fails the tests until it is removed.
