how-to:
#	"Remove everything except openapi.json and run"
# 	TODO improve
	openapi-generator generate -i openapi.json -g go --skip-validate-spec
#	Remove roles from .go files

openapi.json:
	curl https://console.cloud.camunda.io/customer-api/openapi/swagger.json | jq --sort-keys . > $@

.PHONY: openapi.json
