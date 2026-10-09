# Summarizes what changed between two versions of the upstream spec, as Markdown
# for the weekly update PR body. Run with:
#   jq -rn --slurpfile old OLD.json --slurpfile new NEW.json -f scripts/spec-changes.jq

def operations:
  [.paths // {} | to_entries[] | .key as $path | .value | to_entries[]
    | select(.key | IN("get", "put", "post", "delete", "patch", "head", "options"))
    | {key: "\(.key | ascii_upcase) \($path)", value}]
  | from_entries;

def schemas: .components.schemas // {};

def section($title; $items):
  if $items == [] then empty
  else "**\($title):** \($items | map("`\(.)`") | join(", "))" end;

def added($a; $b): [$b | keys[] | select(. as $k | $a | has($k) | not)];
def changed($a; $b): [$a | keys[] | select(. as $k | $b | has($k) and $a[$k] != $b[$k])];

$old[0] as $o | $new[0] as $n
| ($o | operations) as $oo | ($n | operations) as $no
| ($o | schemas) as $os | ($n | schemas) as $ns
| [
    section("Operations added"; added($oo; $no)),
    section("Operations removed"; added($no; $oo)),
    section("Operations changed"; changed($oo; $no)),
    section("Schemas added"; added($os; $ns)),
    section("Schemas removed"; added($ns; $os)),
    section("Schemas changed"; changed($os; $ns))
  ] as $sections
| (if $o.info.version == $n.info.version
   then "**API version:** \($n.info.version) (unchanged)"
   else "**API version:** \($o.info.version) → \($n.info.version)" end),
  "",
  (if $sections == [] then "No operation or schema changes; only other parts of the spec changed."
   else ($sections | join("\n\n")) end)
