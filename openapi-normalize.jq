# Normalizes Camunda's published OpenAPI spec before code generation.
#
# Camunda's spec does not match what the Console API actually serves: it ships
# additive changes without bumping `info.version`, and it declares response
# objects as closed even though the server sends fields the spec omits. Two
# corrections are applied here so the generated Go client can decode real
# responses. Both are deliberate deviations from the published spec -- see the
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

# 3. Do not treat newly-introduced fields as required.
#
#    The same generated UnmarshalJSON that rejects unknown fields also rejects a
#    *missing* required field. Camunda recently marked `Cluster.encryption` and
#    `CreatedClusterClient.audience` required, but a spec that has already proven to
#    drift from the server is not a safe basis for a hard decode requirement -- if
#    any cluster omits them, decoding fails exactly as before, just on a different
#    field. Relaxing them costs nothing (they generate as pointers and are still
#    fully readable) and keeps the New*() constructor signatures stable.
def relax_required($schema; $field):
  if .components.schemas[$schema].required
  then .components.schemas[$schema].required -= [$field]
  else . end;

rename_cluster_status_enum
| open_objects
| relax_required("Cluster"; "encryption")
| relax_required("CreatedClusterClient"; "audience")
