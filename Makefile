.PHONY: generate
generate:
	go run hack/generate/generate.go
	go fmt ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: generate-mocks
generate-mocks: ## Generate mocks for the controller-runtime.
	mockery --dir pkg/util/ --all --output pkg/util/mocks
