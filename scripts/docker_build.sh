#!/bin/bash

# Define the image name and tag (optional)
IMAGE_NAME="osh_registrar"
IMAGE_TAG="latest"

# Define the path to the Dockerfile (default is current directory)
DOCKERFILE_PATH="./registrar/"
DOCKERFILE_FILE="osh_reg.Dockerfile"

# Build the Docker image
echo "Building Docker image: build -t $IMAGE_NAME:$IMAGE_TAG $DOCKERFILE_FILE"
docker build -f "$DOCKERFILE_FILE" -t "$IMAGE_NAME:$IMAGE_TAG" "$DOCKERFILE_PATH"

# Check if the build was successful
if [ $? -eq 0 ]; then
  echo "Docker image $IMAGE_NAME:$IMAGE_TAG built successfully."
else
  echo "Failed to build Docker image."
  exit 1
fi
