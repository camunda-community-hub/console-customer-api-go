# Applies the spec corrections: turns the upstream spec (openapi.upstream.json)
# into the corrected spec (openapi.json) that the generator reads. See GLOSSARY.md.
# spec_correction_test.go checks each correction is applied and still needed.
#
# Camunda's spec does not match what the Console API actually serves: it ships
# additive changes without bumping `info.version`, and it declares response
# objects as closed even though the server sends fields the spec omits. Three
# corrections are applied here so the generated Go client can decode real
# responses. All are deliberate deviations from the published spec -- see the
# rationale on each.

# 1. `ClusterStatus` is a string enum ("Healthy", "Unhealthy", ...). The status
#    object hanging off `Cluster.properties.status` is inline and unnamed, so the
#    generator also wants to call it `ClusterStatus`, and the two collide. Free up
#    the name by renaming the enum to `ClusterComponentStatus`.
#
#    Every `$ref` to the enum is rewritten, wherever it appears. The previous
#    approach (patch-001-openapi-cluster-status.diff) rewrote refs by line context
#    and `patch` applied it with fuzz, silently missing `connectorsStatus` -- which
#    left that field generated as a self-referential object while the API sends a
#    plain string. Rewriting refs structurally cannot miss one.
def rename_cluster_status_enum:
  walk(
    if type == "object" and .["$ref"] == "#/components/schemas/ClusterStatus"
    then .["$ref"] = "#/components/schemas/ClusterComponentStatus"
    else . end
  )
  | .components.schemas.ClusterComponentStatus = .components.schemas.ClusterStatus
  | del(.components.schemas.ClusterStatus);

# 2. Drop every `"additionalProperties": false`.
#
#    The generator emits `decoder.DisallowUnknownFields()` for any model that is
#    closed AND has at least one required property, which turns a backwards-
#    compatible server-side field addition into a hard decode failure. The API
#    demonstrably sends fields the spec does not declare, so `additionalProperties:
#    false` is simply untrue and we do not honor it.
#
#    With the objects open, the generator instead captures unknown fields in an
#    `AdditionalProperties map[string]interface{}` -- so new server fields survive a
#    decode/encode round trip instead of blowing it up.
def open_objects:
  walk(
    if type == "object" and .additionalProperties == false
    then del(.additionalProperties)
    else . end
  );

# 3. Only fields in the required baseline are required in responses.
#
#    The same generated UnmarshalJSON that rejects unknown fields also rejects a
#    *missing* required field. A spec that has already proven to drift from the
#    server is not a safe basis for a new hard decode requirement: when Camunda
#    marked `Cluster.encryption` and `CreatedClusterClient.audience` required, any
#    response omitting them would have failed to decode.
#
#    So in response-only schemas (and the objects nested in them), a field stays
#    required only if upstream requires it AND required-baseline.json lists it.
#    Anything upstream newly requires is relaxed: it generates as a pointer, still
#    fully readable, and existing typed fields consumers rely on stay unchanged.
#    Promote a field by adding it to the baseline. Request schemas are untouched,
#    so the New*() constructors of request bodies follow the spec.
#
#    Baseline keys are the object's path inside components.schemas, joined by "/":
#    "Cluster", "Cluster/properties/channel", "Cluster/properties/ipallowlist/items".

def schema_refs:
  [.. | objects | .["$ref"]? | strings | ltrimstr("#/components/schemas/")];

# Names of the schemas reachable from the given roots, following $refs.
def reachable_schemas($roots):
  .components.schemas as $schemas
  | {seen: [], todo: $roots}
  | until(.todo == [];
      (.todo - .seen) as $new
      | .seen += $new
      | .todo = ([$new[] | $schemas[.] | schema_refs[]] | unique))
  | .seen | unique;

def response_only_schemas:
  reachable_schemas([.paths[][] | objects | .responses // {} | .[] | .content // {} | .[] | .schema | schema_refs[]] | unique)
    - reachable_schemas([.paths[][] | objects | .requestBody // {} | .content // {} | .[] | .schema | schema_refs[]] | unique);

# Every [key, path] of an object with a required list inside the response-only schemas.
def required_locations:
  response_only_schemas[] as $name
  | .components.schemas[$name]
  | path(.. | select(type == "object" and (.required | type) == "array")) as $p
  | [([$name] + ($p | map(tostring)) | join("/")), (["components", "schemas", $name] + $p)];

def apply_required_baseline($baseline):
  reduce required_locations as [$key, $path] (.;
    ($baseline[$key] // []) as $allowed
    | ([getpath($path).required[] | select(. as $f | $allowed | index($f))]) as $kept
    | if $kept == [] then delpaths([$path + ["required"]])
      else setpath($path + ["required"]; $kept) end);

rename_cluster_status_enum
| open_objects
| apply_required_baseline($baseline[0])
