# Stage 1: Build the Go application
FROM golang:1.20 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first to take advantage of Docker layer caching
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files haven't changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o file-picker-service ./cmd/main.go

# Stage 2: Build the final Docker image
FROM alpine:latest

# Create a directory for the app
WORKDIR /app

# Copy the binary from the build stage
COPY --from=builder /app/file-picker-service /app/file-picker-service

# Copy the .env file for configuration (if necessary)
COPY .env /app/.env

# Expose the port the app runs on
EXPOSE 8080

# Set environment variables
ENV PORT 8080

# Run the application
CMD ["/app/file-picker-service"]