# Use the official Golang image as the base image
FROM golang:1.19 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application
RUN go build -o file-downloader-service ./cmd/main.go

# Use a minimal image for the final stage
FROM gcr.io/distroless/base-debian10

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the previous stage
COPY --from=builder /app/file-downloader-service /app/file-downloader-service

# Copy .env if needed (for local testing)
COPY .env /app/.env

# Expose the application port
EXPOSE 50052

# Run the binary
CMD ["/app/file-downloader-service"]