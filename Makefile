
.PHONY: test
test:
	go test ./... -v

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test-docker
test-docker:
	docker build -t test -f test.Dockerfile . && docker run test
