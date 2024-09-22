# Stage 1: Build the Go application
FROM golang:1.18-alpine AS builder

# Install git to fetch Go modules
RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o file-transformation-service ./cmd/main.go

# Stage 2: A minimal runtime image
FROM alpine:latest

# Install necessary libraries for handling images
RUN apk add --no-cache libc6-compat libjpeg-turbo-utils libpng

WORKDIR /root/

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/file-transformation-service .

# Copy the .env file (optional for local environment)
COPY .env .env

# Expose the service port
EXPOSE 50051

# Command to run the executable
CMD ["./file-transformation-service"]