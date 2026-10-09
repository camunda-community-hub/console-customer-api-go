OPENAPI_GENERATOR_IMAGE = openapitools/openapi-generator-cli:v7.26.0
UPSTREAM_SPEC_FILE = openapi.upstream.json
CORRECTED_SPEC_FILE = openapi.json
UPSTREAM_SPEC_URL = https://console.cloud.camunda.io/customer-api/openapi/swagger.json

.PHONY: all fetch clean generate test

all:
	$(MAKE) fetch clean generate

# The only network step: store the upstream spec as published, keys sorted so
# diffs stay stable but otherwise untouched.
fetch:
	curl --fail --silent --show-error $(UPSTREAM_SPEC_URL) \
		| jq --sort-keys . \
		> $(UPSTREAM_SPEC_FILE)

# Camunda's published spec does not match what the API serves; openapi-normalize.jq
# applies the spec corrections and documents why for each one.
$(CORRECTED_SPEC_FILE): $(UPSTREAM_SPEC_FILE) openapi-normalize.jq required-baseline.json
	jq --sort-keys --slurpfile baseline required-baseline.json \
		--from-file openapi-normalize.jq $(UPSTREAM_SPEC_FILE) > $@

clean:
	cat .openapi-generator/FILES | xargs rm -f

generate: $(CORRECTED_SPEC_FILE)
	docker run --rm \
		--user $(shell id -u) \
		-v ${PWD}:/local \
		--workdir /local \
		$(OPENAPI_GENERATOR_IMAGE) generate --config openapi-generator.yaml
	go fmt .
	# go.mod and go.sum are generator-owned (see .openapi-generator/FILES), so
	# generation rewrites them. Restore the real dependency set.
	go mod tidy

test:
	go build -v ./...
	go vet ./...
	go test ./...
