OPENAPI_GENERATOR_IMAGE = openapitools/openapi-generator-cli:v7.20.0
OPENAPI_SPEC_FILE = openapi.json
OPENAPI_SPEC_URL = https://console.cloud.camunda.io/customer-api/openapi/swagger.json

.PHONY: all $(OPENAPI_SPEC_FILE) clean generate test

all:
	$(MAKE) $(OPENAPI_SPEC_FILE) clean generate

# Camunda's published spec does not match what the API serves; openapi-normalize.jq
# corrects it before generation and documents why for each change.
$(OPENAPI_SPEC_FILE):
	curl --fail --silent --show-error $(OPENAPI_SPEC_URL) \
		| jq --sort-keys --from-file openapi-normalize.jq \
		> $@

clean:
	cat .openapi-generator/FILES | xargs rm -f

generate:
	docker run --rm \
		--user $(shell id -u) \
		-v ${PWD}:/local \
		--workdir /local \
		$(OPENAPI_GENERATOR_IMAGE) generate --config openapi-generator.yaml
	go fmt .
	# go.mod and go.sum are generator-owned (see .openapi-generator/FILES), so
	# generation rewrites them and drops testify -- which test/ imports -- down to an
	# indirect dependency. Restore the real dependency set.
	go mod tidy

test:
	go build -v ./...
	go get github.com/stretchr/testify/assert
	go test ./...
