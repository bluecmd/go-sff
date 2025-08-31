# Use Debian 11 (bullseye) as base image
FROM debian:bullseye-slim

# Set environment variables
ENV DEBIAN_FRONTEND=noninteractive
ENV GO_VERSION=1.19.13
ENV GOOS=linux
ENV GOARCH=amd64
ENV PATH=/usr/local/go/bin:$PATH
ENV GOPATH=/go
ENV GOCACHE=/go/cache

# Install required packages
RUN apt-get update && apt-get install -y \
    ca-certificates \
    curl \
    git \
    make \
    && rm -rf /var/lib/apt/lists/*

# Install Go 1.19
RUN curl -fsSL https://golang.org/dl/go${GO_VERSION}.${GOOS}-${GOARCH}.tar.gz | tar -C /usr/local -xzf -

# Create go directory
RUN mkdir -p /go/src /go/bin /go/cache

# Set working directory
WORKDIR /go/src/app

# Copy go mod files first for better caching
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Verify Go installation
RUN go version

# Default command
CMD ["make", "test"]
