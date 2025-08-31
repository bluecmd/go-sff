.PHONY: all build test test-container clean sfputil

all: sfputil

test:
	go test ./...

test-container:
	./run-test-container.sh

test-verbose:
	go test -v ./...

test-coverage:
	go test -cover ./...

sfputil:
	CGO_ENABLED=0 go build ./cmd/sfputil/

clean:
	\rm -f sfputil

deps:
	go mod tidy
	go mod download

fmt:
	go fmt ./...

vet:
	go vet ./...
