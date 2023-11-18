OPENAPI_GENERATOR_IMAGE = openapitools/openapi-generator-cli:v7.1.0
OPENAPI_SPEC_FILE = openapi.json

.PHONY: all $(OPENAPI_SPEC_FILE) clean generate test

all:
	$(MAKE) $(OPENAPI_SPEC_FILE) clean generate

$(OPENAPI_SPEC_FILE):
	curl https://console.cloud.camunda.io/customer-api/openapi/swagger.json | jq --sort-keys . > $@

clean:
	cat .openapi-generator/FILES | xargs rm -f

generate:
	docker run --rm \
		--user $(shell id -u) \
		-v ${PWD}:/local \
		--workdir /local \
		$(OPENAPI_GENERATOR_IMAGE) generate --config openapi-generator.yaml
	go fmt .

test:
	go build -v ./...
	go get github.com/stretchr/testify/assert
	go test ./...
