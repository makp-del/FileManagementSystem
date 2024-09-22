# **File Transformation Service**

## **Overview**

The **file-transformation-service** is responsible for performing various file transformations such as resizing images, converting file formats, and more. This service is built as a microservice and uses gRPC to communicate with other services in the system. It is designed to be flexible, scalable, and integrated into a larger system that involves uploading, downloading, and managing files.

This service communicates with:
- **file-picker-service**: To receive transformation requests.
- **notification-service**: To notify users of the transformation results.

## **Features**

- **Resize Images**: Supports resizing images based on width and height parameters.
- **Convert File Formats**: Allows converting images from one format (e.g., PNG, JPEG) to another.
- **gRPC Communication**: This service exposes APIs via gRPC for other services to consume.
- **Logging**: Detailed logging is implemented using a custom logger.
- **Health Probes**: Supports Kubernetes liveness and readiness probes for monitoring the health of the service.

---

## **APIs Exposed**

This service exposes the following gRPC methods for file transformation:

### **1. TransformFile**
Performs the requested transformation on a file (e.g., resize or format conversion).

- **gRPC Endpoint**: `/TransformFile`
- **Request**:
    ```protobuf
    message TransformFileRequest {
      string file_id = 1;           // Unique ID of the file to transform
      string file_path = 2;         // Path of the file on the server
      string transformation_type = 3; // Type of transformation (resize, convert)
      map<string, string> options = 4; // Key-value options for transformation (e.g., width, height)
    }
    ```
- **Response**:
    ```protobuf
    message TransformFileResponse {
      string file_path = 1;         // Path of the transformed file
      string status = 2;            // Status of the transformation (success, failure)
      string message = 3;           // Additional information (if any)
    }
    ```

- **Supported Transformation Types**:
  - `resize`: Resizes an image based on width and height.
  - `convert`: Converts the file format (e.g., from PNG to JPEG).

### **Example gRPC Request for Resizing**:

```json
{
  "file_id": "12345",
  "file_path": "/path/to/image.png",
  "transformation_type": "resize",
  "options": {
    "width": "800",
    "height": "600"
  }
}
```

### **Example gRPC Request for File Format Conversion**:

```json
{
  "file_id": "12345",
  "file_path": "/path/to/image.png",
  "transformation_type": "convert",
  "options": {
    "format": "jpeg"
  }
}
```

---

## **Folder Structure**

```bash
file-transformation-service/
│
├── cmd/
│   └── main.go                        # Entry point of the application
│
├── internal/
│   ├── handlers/                      # gRPC Handlers for managing file transformations
│   │   └── file_transformation_handler.go
│   ├── services/                      # Core services (business logic)
│   │   └── file_transformation_service.go
│   ├── utils/                         # Utility functions (logging, helpers)
│   └── pkg/                           # Packages like logger, errors
│       └── logger.go                  # Custom logger for logging service events
│
├── proto/
│   └── generated/                     # gRPC generated code
│       └── file_transformation.proto  # Proto file for the service
│
├── Dockerfile                         # Docker configuration file for containerization
├── Makefile                           # Makefile for building, running, and managing the service
└── README.md                          # Service documentation (this file)
```

---

## **Setup and Installation**

### **1. Clone the repository**
```bash
git clone https://github.com/your-username/file-transformation-service.git
cd file-transformation-service
```

### **2. Build the Service**

You can either build the service locally or using Docker.

#### **Local Build**
```bash
make build
```

#### **Docker Build**
```bash
make docker-build
```

### **3. Run the Service Locally**
You can run the service locally using the following command:
```bash
make run
```

---

## **Docker Deployment**

The **file-transformation-service** is containerized using Docker. Below is the relevant Dockerfile for the service:

### **Dockerfile**

```Dockerfile
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
```

---

## **Kubernetes Deployment**

### **1. Apply Kubernetes Deployment and Service Files**

To deploy the service in Kubernetes, run the following commands:

```bash
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```

### **2. Verify Deployment**

Check the pods to ensure they are running:
```bash
kubectl get pods
```

Check the logs of the pod for any errors or important information:
```bash
kubectl logs <pod-name>
```

---

## **Environment Variables**

This service relies on a set of environment variables that can be configured via a `.env` file or through Kubernetes Secrets.

- **SERVER_PORT**: The port on which the service runs (default: 50051).
- **LOG_LEVEL**: The logging level for the service (default: INFO).
- **DATABASE_URL**: The PostgreSQL connection string (if using a database).

---

## **Health Checks**

### **Liveness Probe**

```yaml
livenessProbe:
  httpGet:
    path: /
    port: http
```

### **Readiness Probe**

```yaml
readinessProbe:
  httpGet:
    path: /
    port: http
```

These health probes ensure Kubernetes can monitor and restart the service if necessary.

---

## **Helm Configuration**

The `values.yaml` file provides default values for deploying the service using Helm.

### **Values.yaml**

```yaml
replicaCount: 1
image:
  repository: file-transformation-service
  pullPolicy: IfNotPresent
  tag: "v1.0"

resources: {}
livenessProbe:
  httpGet:
    path: /
    port: 50051
readinessProbe:
  httpGet:
    path: /
    port: 50051
service:
  type: ClusterIP
  port: 50051
```

Use the following Helm commands to install the service:

```bash
helm install file-transformation-service .
```

---

## **Testing the Service**

You can test the service by sending gRPC requests using a tool like `grpcurl` or by integrating it with your system.

### **Test TransformFile Request**:
```bash
grpcurl -d '{"file_id":"12345","file_path":"/path/to/image.png","transformation_type":"resize","options":{"width":"800","height":"600"}}' localhost:50051 file_transformation_service.TransformFile
```

---

## **Contributions**

Feel free to submit issues or pull requests. All contributions are welcome!

---