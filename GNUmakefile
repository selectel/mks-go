OAPI_CODEGEN_VERSION=v2.6.0
MOCKERY_VERSION=v2.53.6
GOLANGCI_LINT_VERSION=v2.11.4

default: tests

tests: golangci-lint unittest

unittest:
	@sh -c "'$(CURDIR)/scripts/gotest.sh'"

golangci-lint:
	@VERSION=$(GOLANGCI_LINT_VERSION) sh -c "'$(CURDIR)/scripts/golangci_lint_check.sh'"

generate-mocks:
	go run github.com/vektra/mockery/v2@$(MOCKERY_VERSION)

oapi-codegen:
	@cd swagger && go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@$(OAPI_CODEGEN_VERSION) --config=oapi-codegen-models.yaml managed-kubernetes.swagger.yaml
	@cd swagger && go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@$(OAPI_CODEGEN_VERSION) --config=oapi-codegen-client.yaml managed-kubernetes.swagger.yaml

ci-check-mocks:
	@VERSION=$(MOCKERY_VERSION) sh -c "'$(CURDIR)/scripts/check_mockery.sh'"

ci-check-oapi:
	@VERSION=$(OAPI_CODEGEN_VERSION) sh -c "'$(CURDIR)/scripts/check_oapi.sh'"

.PHONY: tests unittest golangci-lint generate-mocks oapi-codegen ci-check-mocks ci-check-oapi
