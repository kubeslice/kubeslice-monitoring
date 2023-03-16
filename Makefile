.PHONY: generate
generate:
	go run hack/generate/generate.go
	go fmt ./...

.PHONY: fmt
fmt:
	go fmt ./...
