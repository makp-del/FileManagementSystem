# Variables
SERVICE_NAME = file-downloader-service
DOCKER_IMAGE = file-downloader-service-image
DOCKER_CONTAINER = file-downloader-container

# Build the Go application
build:
	@echo "Building the Go application..."
	go build -o $(SERVICE_NAME) ./cmd/main.go

# Run the Go application
run:
	@echo "Running the Go application..."
	./$(SERVICE_NAME)

# Build the Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

# Run the Docker container
docker-run:
	@echo "Running Docker container..."
	docker run --rm -p 50052:50052 --name $(DOCKER_CONTAINER) $(DOCKER_IMAGE)

# Stop the Docker container
docker-stop:
	@echo "Stopping Docker container..."
	docker stop $(DOCKER_CONTAINER)

# Clean the built files
clean:
	@echo "Cleaning up..."
	rm -f $(SERVICE_NAME)

# Build and run the Go application
build-run: build run

# Build and run the Docker container
docker-build-run: docker-build docker-run