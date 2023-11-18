OPENAPI_GENERATOR_IMAGE = openapitools/openapi-generator-cli:v7.1.0
GIT_USER_ID = sijoma
GIT_REPO_ID = console-customer-api-go
OPENAPI_SPEC_FILE = openapi.json

.PHONY: all $(OPENAPI_SPEC_FILE) clean generate

all:
	$(MAKE) $(OPENAPI_SPEC_FILE) clean generate

$(OPENAPI_SPEC_FILE):
	curl https://console.cloud.camunda.io/customer-api/openapi/swagger.json | jq --sort-keys . > $@

clean:
	rm -rf go.mod go.sum *.go .openapi-generator/ api/ docs/ test/ README.md git_push.sh

generate:
	docker run --rm \
		--user $(shell id -u) \
		-v ${PWD}:/local $(OPENAPI_GENERATOR_IMAGE) \
		generate \
		--input-spec /local/$(OPENAPI_SPEC_FILE) \
		--output /local \
		--generator-name go \
		--git-user-id $(GIT_USER_ID) \
		--git-repo-id $(GIT_REPO_ID) \
		--skip-validate-spec
