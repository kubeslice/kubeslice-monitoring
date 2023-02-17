.PHONY: generate-events
generate-events: ## Generate event_names.go
	./event.sh

.PHONY: generate-mocks
generate-mocks: ## Generate mocks for the controller-runtime.
	mockery --dir pkg/util/ --all --output pkg/util/mocks