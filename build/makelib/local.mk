# Override Docker command
DOCKER := podman
DOCKER_BUILDX := podman build

# Disable buildx features not supported by podman
USE_BUILDX := false

# Use local image store
BUILD_REGISTRY := localhost
