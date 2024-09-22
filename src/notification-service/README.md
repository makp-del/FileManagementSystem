# **Notification Service**

## **Overview**
The **Notification Service** is a microservice responsible for handling real-time notifications via WebSockets. It allows other microservices to send notifications to connected clients through a gRPC interface. Notifications are sent to clients using WebSocket connections, ensuring real-time communication.

### **Key Features**:
- Real-time notifications via WebSockets.
- gRPC interface for other services to send notifications.
- Scalable and containerized for Kubernetes deployments.
  
---

## **Purpose**
This service handles real-time notifications for users connected via WebSockets. Other microservices (such as **file-picker-service** or **file-downloader-service**) send notification requests to the **notification-service** through gRPC. The notification-service then forwards these notifications to connected WebSocket clients.

---

## **APIs**

### **APIs Exposed**:

1. **WebSocket Endpoint**:
   - **Endpoint**: `/ws`
   - **Method**: `GET`
   - **Purpose**: This endpoint allows clients to establish a WebSocket connection for receiving notifications.
   - **Response**: A WebSocket connection for real-time notifications.

2. **gRPC API**:
   - **RPC Method**: `SendNotification`
   - **Request**: 
     ```protobuf
     message SendNotificationRequest {
         string user_id = 1;   // User ID to send the notification to
         string message = 2;   // The notification message
     }
     ```
   - **Response**:
     ```protobuf
     message SendNotificationResponse {
         bool success = 1;     // Whether the notification was sent successfully
     }
     ```
   - **Purpose**: Other services call this gRPC API to send a notification to a WebSocket client.

### **APIs Consumed**:
- **gRPC Calls from Other Services**: 
  - Services like **file-picker-service** or **file-downloader-service** will use the `SendNotification` gRPC API to trigger notifications to clients connected via WebSockets.

---

## **Services Communicated With**

The **notification-service** communicates with the following services:
- **file-picker-service**: To notify users when a file operation (e.g., upload, share) is completed.
- **file-downloader-service**: To notify users when a file is downloaded or processed.

---

## **Folder Structure**

Here is a detailed breakdown of the folder structure for the **notification-service**:

```
notification-service/
│
├── cmd/
│   └── main.go                          # Entry point for starting the service (WebSocket and gRPC server)
│
├── internal/
│   ├── handlers/                        # WebSocket and gRPC handlers
│   │   ├── websocket_handler.go         # WebSocket connection handler
│   │   ├── notification_handler.go      # gRPC handler to handle SendNotification requests
│   ├── services/                        # Core notification business logic
│   │   └── notification_service.go      # Logic for sending notifications via WebSocket
│   ├── websocket/                       # WebSocket connection management
│   │   ├── connection.go                # WebSocket connection management logic
│   │   └── hub.go                       # WebSocket hub for managing multiple clients
│   ├── db/                              # (Optional) Database connection setup
│   │   └── database.go                  # Initializes and manages the database connection
│   ├── logger/                          # Logging utilities
│   │   └── logger.go                    # Handles logging setup
│   ├── errors/                          # Custom error handling
│   │   └── errors.go                    # Defines custom errors for the service
│   └── proto/                           # Protobuf-generated files for gRPC
│       └── generated/                   # Generated protobuf Go files
│
├── config/
│   └── config.go                        # Configuration loader for environment variables
│
├── proto/
│   └── notification_service.proto       # Protobuf definition for the gRPC notification service
│
├── Dockerfile                           # Dockerfile for containerizing the notification service
├── Makefile                             # Makefile for building and running the service
├── .env                                 # Environment variables for local development (WebSocket and gRPC ports, DB URL)
├── go.mod                               # Go module dependencies
└── README.md                            # Project documentation
```

---

## **Deployment**

### **Prerequisites**:
- Docker
- Kubernetes cluster (for deployment)
- Helm (optional, if using Helm for deployment)

### **Deployment Steps**:

1. **Build Docker Image**:
   You can build the Docker image for the **notification-service** using the provided **Dockerfile**:
   
   ```bash
   docker build -t notification-service:v1 .
   ```

2. **Run Locally with Docker**:
   If you want to run the service locally with Docker, you can use the **Makefile**:
   
   ```bash
   make run
   ```

3. **Kubernetes Deployment**:

   **Service YAML**:
   ```yaml
   apiVersion: v1
   kind: Service
   metadata:
     name: notification-service
   spec:
     type: ClusterIP
     selector:
       app: notification
     ports:
       - protocol: TCP
         port: 50054   # WebSocket port
         targetPort: 50054
       - protocol: TCP
         port: 50055   # gRPC port
         targetPort: 50055
   ```

   **Deployment YAML**:
   ```yaml
   apiVersion: apps/v1
   kind: Deployment
   metadata:
     labels:
       app: notification
     name: notification
   spec:
     replicas: 3
     selector:
       matchLabels:
         app: notification
     template:
       metadata:
         labels:
           app: notification
       spec:
         containers:
         - image: notification:v1
           name: notification
           ports:
           - containerPort: 50054   # WebSocket port
           - containerPort: 50055   # gRPC port
   ```

   Apply the **service** and **deployment** YAML files to the Kubernetes cluster:

   ```bash
   kubectl apply -f service.yaml
   kubectl apply -f deployment.yaml
   ```

4. **Helm Deployment** (Optional):
   If you're using Helm, you can deploy the **notification-service** using a Helm chart with the provided `values.yaml` file.

   **Helm Commands**:
   ```bash
   helm install notification-service ./notification-service-chart
   ```

---

## **Environment Variables**

The service uses the following environment variables, defined in the **`.env`** file:

```bash
# WebSocket server port
NOTIFICATION_SERVICE_PORT=50054

# gRPC server port
NOTIFICATION_SERVICE_GRPC_PORT=50055

# Database URL (if using a database)
DATABASE_URL=postgres://admin:admin@localhost:5432/notification_service_db?sslmode=disable
```

---

## **Conclusion**
The **notification-service** is a real-time WebSocket-based service that integrates with other microservices to provide notifications to connected clients. It is built with scalability in mind and is designed to run efficiently in a Kubernetes environment. You can scale the service horizontally by increasing the number of replicas, ensuring high availability and fault tolerance.

---