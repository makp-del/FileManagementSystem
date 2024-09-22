# Use an official Go image to create a build stage.
FROM golang:1.19 as builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all Go module dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o permission-service cmd/main.go

# Use a smaller base image for the final output
FROM alpine:latest  

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/permission-service .

# Expose the port the service will run on
EXPOSE 50053

# Command to run the binary
CMD ["./permission-service"]