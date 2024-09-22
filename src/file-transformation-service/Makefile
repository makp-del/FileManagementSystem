# Variables
SERVICE_NAME := file-transformation-service
DOCKER_IMAGE := $(SERVICE_NAME):latest
DOCKER_CONTAINER := $(SERVICE_NAME)_container

# Default target: Build and run the Go application locally
.PHONY: run
run:
	@echo "Running the service locally..."
	go run ./cmd/main.go

# Build the Go binary
.PHONY: build
build:
	@echo "Building the Go application..."
	go build -o $(SERVICE_NAME) ./cmd/main.go

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	go test ./... -v

# Docker targets

# Build the Docker image
.PHONY: docker-build
docker-build:
	@echo "Building the Docker image..."
	docker build -t $(DOCKER_IMAGE) .

# Run the Docker container
.PHONY: docker-run
docker-run:
	@echo "Running the Docker container..."
	docker run -p 50051:50051 --name $(DOCKER_CONTAINER) $(DOCKER_IMAGE)

# Stop the Docker container
.PHONY: docker-stop
docker-stop:
	@echo "Stopping the Docker container..."
	docker stop $(DOCKER_CONTAINER) || true
	docker rm $(DOCKER_CONTAINER) || true

# Clean up generated files
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -f $(SERVICE_NAME)
	docker rmi $(DOCKER_IMAGE) || true