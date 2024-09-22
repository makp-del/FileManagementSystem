# **File Downloader Service**

## **Overview**
The **File Downloader Service** is a microservice designed to handle file downloads from cloud storage providers like Google Drive and Dropbox. It receives requests from other services, fetches the files from cloud storage, and saves them to the server's local file system. The service also communicates with the **Notification Service** to inform users about the progress or completion of file downloads.

### **Key Features**:
- Downloads files from various cloud storage providers (e.g., Google Drive, Dropbox).
- Provides gRPC APIs for file download requests.
- Communicates with the **Notification Service** to send real-time updates about downloads.

---

## **Purpose**
The **File Downloader Service** is responsible for handling file downloads from cloud providers. It interacts with the **File Picker Service**, which sends requests when a user wants to download a file. The **File Downloader Service** retrieves the file and stores it locally. Additionally, it notifies users of the download status via the **Notification Service**.

---

## **APIs**

### **APIs Exposed**:

1. **gRPC API**:
   - **RPC Method**: `DownloadFile`
   - **Request**: 
     ```protobuf
     message DownloadFileRequest {
         string file_id = 1;       // File ID in the cloud provider
         string provider = 2;      // Cloud provider (e.g., "google_drive", "dropbox")
         string auth_token = 3;    // Access token for the cloud provider
     }
     ```
   - **Response**:
     ```protobuf
     message DownloadFileResponse {
         string file_path = 1;     // Local file path where the file is saved
     }
     ```
   - **Purpose**: Other services, like the **File Picker Service**, call this gRPC API to initiate file downloads from cloud providers.
   - **Response**: The response includes the local file path where the downloaded file is saved.

### **APIs Consumed**:
- **Notification Service**:
  - The **File Downloader Service** sends a notification to the **Notification Service** after a file is successfully downloaded to notify the user.
  - This is done using the `SendNotification` gRPC API exposed by the **Notification Service**.

---

## **Services Communicated With**

The **File Downloader Service** communicates with the following services:
- **File Picker Service**: Receives download requests and retrieves files from the cloud providers on behalf of the users.
- **Notification Service**: Sends notifications to users to update them on the progress or completion of a file download.

---

## **Folder Structure**

Here is a detailed breakdown of the folder structure for the **File Downloader Service**:

```
file-downloader-service/
│
├── cmd/
│   └── main.go                          # Entry point for starting the service (gRPC server)
│
├── internal/
│   ├── clients/                         # Clients for cloud storage providers
│   │   ├── google_drive_client.go       # Google Drive API client
│   │   ├── dropbox_client.go            # Dropbox API client
│   ├── handlers/                        # gRPC handler for processing file download requests
│   │   └── file_downloader_handler.go   # Handles gRPC file download requests
│   ├── services/                        # Core download logic
│   │   └── file_downloader_service.go   # Logic for downloading files and notifying the user
│   ├── notification/                    # gRPC client for the Notification Service
│   │   └── notification_client.go       # Sends notifications after the file is downloaded
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
│   └── file_downloader.proto            # Protobuf definition for the gRPC file downloader service
│
├── Dockerfile                           # Dockerfile for containerizing the file downloader service
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
   You can build the Docker image for the **file-downloader-service** using the provided **Dockerfile**:
   
   ```bash
   docker build -t downloader:v1 .
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
     name: file-downloader-service
   spec:
     type: ClusterIP
     selector:
       app: downloader
     ports:
       - protocol: TCP
         port: 8080       # Expose port 8080
         targetPort: 8080 # Target the internal port
   ```

   **Deployment YAML**:
   ```yaml
   apiVersion: apps/v1
   kind: Deployment
   metadata:
     labels:
       app: downloader
     name: downloader
   spec:
     replicas: 3  # Ensure high availability
     selector:
       matchLabels:
         app: downloader
     template:
       metadata:
         labels:
           app: downloader
       spec:
         containers:
         - image: downloader:v1
           name: downloader
           ports:
           - containerPort: 8080   # Expose internal container port
   ```

   Apply the **service** and **deployment** YAML files to the Kubernetes cluster:

   ```bash
   kubectl apply -f service.yaml
   kubectl apply -f deployment.yaml
   ```

4. **Helm Deployment** (Optional):
   If you're using Helm, you can deploy the **file-downloader-service** using a Helm chart with the provided `values.yaml` file.

   **Helm Commands**:
   ```bash
   helm install file-downloader-service ./file-downloader-service-chart
   ```

---

## **Environment Variables**

The service uses the following environment variables, defined in the **`.env`** file:

```bash
# Service ports
FILE_DOWNLOADER_SERVICE_PORT=8080

# Database URL (if using a database)
DATABASE_URL=postgres://admin:admin@localhost:5432/file_downloader_service_db?sslmode=disable

# Cloud storage API keys/tokens (optional for accessing cloud providers)
GOOGLE_DRIVE_API_KEY=your_google_drive_api_key
DROPBOX_TOKEN=your_dropbox_token

# Notification service gRPC address
NOTIFICATION_SERVICE_GRPC_ADDR=notification-service:50055
```

---

## **Conclusion**
The **File Downloader Service** is responsible for retrieving files from cloud providers, saving them locally, and notifying users of the download status. It is designed to work seamlessly with the **Notification Service** and is scalable in a Kubernetes environment. With proper configuration and resource management, the service ensures reliable file download operations and smooth integration with other services.

---