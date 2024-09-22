# **Permission Service**

## **Overview**

The **Permission Service** is a microservice responsible for managing access control and file permissions within the Ema system. This service facilitates:
- Granting and revoking file access to specific users.
- Checking if a user has the appropriate permissions to access or modify a file (read, write, delete, etc.).
- Managing shared file access and ensuring proper security controls.

This service communicates with other services in the Ema system, including:
- **File Picker Service**: Requests permission checks when users attempt to access or modify files.
- **Notification Service**: Sends notifications when permission updates are made.

## **APIs**

### **gRPC APIs**

The Permission Service communicates with other services via gRPC. Below are the gRPC APIs it exposes:

### **Check Permission**
This endpoint checks whether a user has the required permission for a file.

- **Method**: `CheckPermission`
- **Request**:
    ```protobuf
    message CheckPermissionRequest {
      uint64 user_id = 1;
      string file_id = 2;
      string permission = 3;  // The permission to check (read, write, delete, etc.)
    }
    ```
- **Response**:
    ```protobuf
    message CheckPermissionResponse {
      bool has_permission = 1;
    }
    ```
- **Description**: This method checks if the user has the specific permission for the file in question.

### **Update Permission**
This endpoint is used to grant or revoke file permissions for a user.

- **Method**: `UpdatePermission`
- **Request**:
    ```protobuf
    message UpdatePermissionRequest {
      uint64 owner_id = 1;
      uint64 shared_user_id = 2;
      repeated string file_ids = 3;
      repeated string permissions = 4;  // List of permissions to grant (read, write, delete, etc.)
      bool is_owner = 5;  // Whether this is an owner permission
    }
    ```
- **Response**:
    ```protobuf
    message UpdatePermissionResponse {
      bool success = 1;
      string message = 2;
    }
    ```
- **Description**: This method is used to grant or revoke permissions for a particular file or set of files to another user. The `is_owner` field determines if the action involves owner permissions.

## **Architecture & Folder Structure**

Here is the directory structure of the service:

```
permission-service/
│
├── cmd/
│   └── main.go        # Entry point of the application
│
├── internal/
│   ├── handlers/      # gRPC handlers for permission-related functionality
│   │   └── permission_handler.go
│   ├── services/      # Core business logic for permissions management
│   │   └── permission_service.go
│   ├── models/        # Database models for permission storage and queries
│   │   └── permission.go
│   └── db/            # Database initialization and connection handling
│       └── database.go
│
├── proto/             # Protocol Buffers definition files
│   └── permission.proto
│
├── config/
│   └── config.go      # Configuration loading and environment management
│
├── pkg/               # Utility libraries (logging, helpers)
│   └── logger.go
│
├── .env               # Environment variables (for local development)
├── Dockerfile         # Docker configuration
├── Makefile           # Makefile for building and running the service
└── go.mod             # Go module dependencies
```

## **Deployment**

This service is containerized and deployed using Kubernetes. The deployment manifests and Helm chart configuration files are included.

### **Deployment YAML**
The **Deployment** and **Service** YAML for deploying the **Permission Service** are included:

- **permission-deployment.yaml**: Deploys the service as a Kubernetes Deployment with 3 replicas.
- **permission-service.yaml**: Exposes the service internally via a **ClusterIP** service.

### **Helm Values**

A Helm **values.yaml** is provided to customize the deployment. The key settings include:
- **Replica Count**: Number of replicas to run.
- **Resources**: Resource limits and requests for the containers.
- **Liveness & Readiness Probes**: Probes to check the health of the service.

### **Deployment Steps**

1. **Clone the repository**:
   ```bash
   git clone https://github.com/your-org/permission-service.git
   cd permission-service
   ```

2. **Build the Docker image**:
   ```bash
   make docker-build
   ```

3. **Run the Docker container locally**:
   ```bash
   make docker-run
   ```

4. **Deploy on Kubernetes**:
   Apply the deployment and service manifests:
   ```bash
   kubectl apply -f permission-deployment.yaml
   kubectl apply -f permission-service.yaml
   ```

5. **Using Helm**:
   Install the service using Helm:
   ```bash
   helm install permission-service ./helm/permission-service
   ```

## **Environment Variables**

Environment variables are managed via a `.env` file. The key variables include:
- `DATABASE_URL`: The URL of the PostgreSQL database where permission data is stored.
- `GRPC_PORT`: The gRPC server port (default: 8080).

## **Running the Service Locally**

To run the service locally, follow these steps:

1. **Install Dependencies**:
   Install the necessary Go dependencies:
   ```bash
   go mod download
   ```

2. **Run the Service**:
   Use the provided Makefile to run the service:
   ```bash
   make run
   ```

3. **Run Tests**:
   Run the unit tests for the service:
   ```bash
   make test
   ```

## **Database**

The **Permission Service** uses **PostgreSQL** for managing permissions. The `permission` table stores the owner and access control details for each file.

Here’s an example of the **permission** model:

```go
type Permission struct {
    ID            uint64 `gorm:"primaryKey"`
    FileID        string
    OwnerID       uint64
    SharedUserID  uint64
    Permission    string // e.g., "read", "write", "delete"
}
```

The **Database** folder contains the setup logic for the PostgreSQL connection. Ensure the **DATABASE_URL** environment variable is set correctly in your `.env` file.

## **Logging**

The service uses a custom logger defined in `pkg/logger.go`. All events and errors are logged to help with debugging and monitoring.

## **Probes**

The **Liveness** and **Readiness** probes are exposed on `/healthz` and `/readyz` respectively. These endpoints ensure that the service is running correctly and is ready to serve traffic.

---