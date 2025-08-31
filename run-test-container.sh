#!/bin/bash

# Script to run the test container using either podman or docker
# Prefers podman but falls back to docker if podman is not available

set -e

# Function to check if a command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Function to run container with the specified runtime
run_container() {
    local runtime=$1
    local image_name="local.local/go-sff-test"
    echo "Building container image using $runtime..."
    $runtime build -t "$image_name" -f Containerfile .
    echo "Running tests in container..."
    $runtime run --rm "$image_name"
}

# Check if podman is available
if command_exists podman; then
    echo "Using podman..."
    run_container "podman"
# Fall back to docker if podman is not available
elif command_exists docker; then
    echo "Using docker..."
    run_container "docker"
else
    echo "Error: Neither podman nor docker is available on this system."
    echo "Please install one of them to run the test container."
    exit 1
fi
