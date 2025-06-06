IS_IN_PROGRESS = "is in progress ..."

## gen: will generate mock for usecases & repositories interfaces
.PHONY: gen
gen:
	@echo "make gen ${IS_IN_PROGRESS}"

.PHONY: lint-fix
lint-fix:
	@echo "make lint ${IS_IN_PROGRESS}"
	@golangci-lint run --fix

.PHONY: lint-check
lint-check:
	@echo "make lint ${IS_IN_PROGRESS}"
	@golangci-lint run


## e2e-test: will test with e2e tags
.PHONY: e2e-test
e2e-test:
	@echo "make e2e-test ${IS_IN_PROGRESS}"
	@go clean -testcache
	@go test --race -timeout=90s -failfast \
		-vet= -cover -covermode=atomic -coverprofile=./build/coverage/e2e.out \
		-tags=e2e ./test/...\

## tests: run tests and any dependencies
.PHONY: tests
tests: e2e-test