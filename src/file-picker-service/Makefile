# Variables
APP_NAME = file-picker-service
DOCKER_IMAGE = file-picker-service:latest
DOCKER_CONTAINER = file-picker-service-container
PORT = 8080

# Build the Go application
build:
	@echo "Building the application..."
	go build -o ./bin/$(APP_NAME) ./cmd/main.go

# Run the Go application
run: build
	@echo "Running the application..."
	./bin/$(APP_NAME)

# Clean up the binary
clean:
	@echo "Cleaning up..."
	rm -f ./bin/$(APP_NAME)

# Build the Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

# Run the Docker container
docker-run:
	@echo "Running Docker container..."
	docker run --name $(DOCKER_CONTAINER) --env-file .env -p $(PORT):8080 $(DOCKER_IMAGE)

# Stop and remove the Docker container
docker-stop:
	@echo "Stopping and removing Docker container..."
	docker stop $(DOCKER_CONTAINER) || true
	docker rm $(DOCKER_CONTAINER) || true

# Clean up Docker images and containers
docker-clean: docker-stop
	@echo "Cleaning up Docker images..."
	docker rmi $(DOCKER_IMAGE)

# Run tests
test:
	@echo "Running tests..."
	go test ./...

# Format the​⬤