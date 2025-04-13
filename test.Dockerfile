FROM golang:1.22.6 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
ADD vendor vendor
COPY Makefile Makefile


# Copy the go source
COPY hack/ hack/
COPY config/ config
COPY pkg/ pkg/

CMD ["make", "test"]
