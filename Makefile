.PHONY: all build test test-container clean sfpdiag

all: sfpdiag

test:
	go test ./...

test-container:
	./run-test-container.sh

test-verbose:
	go test -v ./...

test-coverage:
	go test -cover ./...

sfpdiag:
	CGO_ENABLED=0 go build ./cmd/sfpdiag/

clean:
	\rm -f sfpdiag

deps:
	go mod tidy
	go mod download

fmt:
	go fmt ./...

vet:
	go vet ./...
