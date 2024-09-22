# The name of the Docker image
IMAGE_NAME=notification-service

# The tag for the Docker image
TAG=latest

# Build the Docker image
build:
	@echo "Building the Docker image..."
	docker build -t $(IMAGE_NAME):$(TAG) .

# Run the Docker container
run:
	@echo "Running the Docker container..."
	docker run --env-file .env -p 50054:50054 -p 50055:50055 $(IMAGE_NAME):$(TAG)

# Clean up Docker images and containers
clean:
	@echo "Cleaning up Docker images and containers..."
	docker rm -f $(shell docker ps -aq --filter "ancestor=$(IMAGE_NAME)") || true
	docker rmi -f $(IMAGE_NAME):$(TAG) || true

# Build the Go binary (for local dev, if needed)
build-local:
	@echo "Building the Go binary locally..."
	CGO_ENABLED=0 GOOS=linux go build -o notification-service cmd/main.go

# Run the Go service locally (for local development)
run-local:
	@echo "Running the service locally..."
	./notification-service