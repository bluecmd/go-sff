.PHONY: all build test clean example

all: example

test:
	go test ./...

test-verbose:
	go test -v ./...

test-coverage:
	go test -cover ./...

example:
	(cd example; CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build)

clean:
	\rm -f example/example

deps:
	go mod tidy
	go mod download

fmt:
	go fmt ./...

vet:
	go vet ./...
