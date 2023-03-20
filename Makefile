.PHONY: generate
generate:
	go run hack/generate/generate.go
	go fmt ./...

.PHONY: test
test:
	go test ./... -v

.PHONY: fmt
fmt:
	go fmt ./...
