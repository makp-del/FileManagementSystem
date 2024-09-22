# **File Picker Service**

## **Overview**
The **file-picker-service** is a crucial component in the microservices architecture, serving as the primary user-facing service for managing files within the system. It allows users to upload, list, download, share, and transform files. The service handles file operations and communicates with other microservices (e.g., permissions, file downloader, transformation, and notification services) via gRPC to perform complex actions like cloud-based file downloads, file transformations, and permission management.

---

## **Purpose**
The **file-picker-service** provides an interface for users to interact with their files, either through uploading new files, downloading existing ones, or sharing them with other users. Additionally, it supports performing transformations (like OCR) on files and notifying users of key events.

### **Key Features**:
- File upload and download support.
- Integration with cloud storage services (e.g., Google Drive, Dropbox).
- Sharing files with other users and managing file permissions.
- Triggering file transformations (e.g., OCR).
- Real-time notifications about file operations.

---

## **APIs**

### **1. File Upload API**

- **Endpoint**: `POST /api/upload`
- **Description**: Uploads a file to the server.
- **Request**:
  - **Headers**:
    - `Authorization`: Bearer token (JWT) for user authentication.
  - **Body (Multipart Form Data)**:
    - `file`: The file to be uploaded.
- **Response**: 
  - **200 OK**: File uploaded successfully.
  - **400 Bad Request**: Invalid file or missing file.
  - **500 Internal Server Error**: Failed to save the file.
  
- **Example**:
  ```bash
  curl -X POST -H "Authorization: Bearer <token>" -F "file=@/path/to/file.txt" http://<host>/api/upload
  ```

---

### **2. List Files API**

- **Endpoint**: `GET /api/files`
- **Description**: Lists all the files that the user has access to, including owned files and shared files.
- **Request**:
  - **Headers**:
    - `Authorization`: Bearer token (JWT) for user authentication.
- **Response**:
  - **200 OK**: Returns a list of files.
  - **500 Internal Server Error**: Failed to retrieve files.
  
- **Response Format**:
  ```json
  [
    {
      "id": "fileID1",
      "fileName": "file1.txt",
      "filePath": "./uploads/file1.txt",
      "isShared": false,
      "ownerId": 123
    },
    {
      "id": "fileID2",
      "fileName": "sharedfile.txt",
      "filePath": "./uploads/sharedfile.txt",
      "isShared": true,
      "ownerId": 124
    }
  ]
  ```

- **Example**:
  ```bash
  curl -X GET -H "Authorization: Bearer <token>" http://<host>/api/files
  ```

---

### **3. File Download API**

- **Endpoint**: `POST /api/files/download`
- **Description**: Downloads a file from cloud storage (Google Drive, Dropbox, etc.) and stores it on the local server.
- **Request**:
  - **Headers**:
    - `Authorization`: Bearer token (JWT) for user authentication.
  - **Body (JSON)**:
    ```json
    {
      "file_id": "fileID123",
      "provider": "google_drive",  // Supported providers: google_drive, dropbox
      "auth_token": "OAuthTokenFromProvider"
    }
    ```
- **Response**:
  - **200 OK**: Returns the local file path where the file is saved.
  - **400 Bad Request**: Invalid input.
  - **500 Internal Server Error**: Failed to download the file.
  
- **Response Format**:
  ```json
  {
    "file_path": "./uploads/fileID123.txt"
  }
  ```

- **Example**:
  ```bash
  curl -X POST -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d '{"file_id": "fileID123", "provider": "google_drive", "auth_token": "OAuthTokenFromProvider"}' http://<host>/api/files/download
  ```

---

### **4. Share File API**

- **Endpoint**: `POST /api/files/share`
- **Description**: Allows a file owner to share their file with another user by updating the file’s permissions.
- **Request**:
  - **Headers**:
    - `Authorization`: Bearer token (JWT) for user authentication.
  - **Body (JSON)**:
    ```json
    {
      "shared_user_id": 456,   // ID of the user the file is being shared with
      "file_ids": ["fileID1", "fileID2"],   // List of file IDs to share
      "permissions": ["read", "write"]   // Permissions granted to the shared user
    }
    ```
- **Response**:
  - **200 OK**: Permissions updated successfully.
  - **400 Bad Request**: Invalid input.
  - **500 Internal Server Error**: Failed to update permissions.
  
- **Example**:
  ```bash
  curl -X POST -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d '{"shared_user_id": 456, "file_ids": ["fileID1", "fileID2"], "permissions": ["read", "write"]}' http://<host>/api/files/share
  ```

---

### **5. File Transformation API**

- **Endpoint**: `POST /api/transform/:id`
- **Description**: Initiates a transformation (e.g., OCR) for the specified file.
- **Request**:
  - **Headers**:
    - `Authorization`: Bearer token (JWT) for user authentication.
  - **Path Parameters**:
    - `id`: The ID of the file to transform.
  - **Body (JSON)**:
    ```json
    {
      "transformation_type": "ocr"  // Example: "ocr", could also be "image_recognition", etc.
    }
    ```
- **Response**:
  - **200 OK**: Transformation initiated successfully.
  - **400 Bad Request**: Invalid input or permissions.
  - **500 Internal Server Error**: Failed to initiate transformation.
  
- **Example**:
  ```bash
  curl -X POST -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d '{"transformation_type": "ocr"}' http://<host>/api/transform/fileID123
  ```

---

## **Inter-Service Communication**

The **file-picker-service** communicates with the following microservices via gRPC:

- **File Downloader Service**: To download files from external providers like Google Drive or Dropbox.
- **Permissions Service**: To verify and update file permissions.
- **Transformation Service**: To request file transformations such as OCR or image recognition.
- **Notification Service**: To send notifications to users about important events such as uploads, downloads, or transformations.

---

## **Folder Structure**

Here is a detailed breakdown of the folder structure for the **file-picker-service**:

```
file-picker-service/
│
├── cmd/
│   └── main.go                          # Entry point for starting the service
│
├── internal/
│   ├── db/                              # Database initialization and connection
│   │   └── database.go
│   ├── handlers/                        # HTTP handlers for API routes
│   │   └── file_handler.go
│   ├── services/                        # Core logic for file operations
│   │   └── file_service.go
│   ├── models/                          # Database models for files and users
│   │   ├── file.go
│   │   ├── user.go
│   ├── pkg/                             # Utility packages (e.g., logger, error handling)
│   │   └── logger.go
│   ├── proto/                           # gRPC-generated files for inter-service communication
│       └── generated/
│
├── config/                              # Configuration handling (loading .env variables)
│   └── config.go
│
├── proto/                               # Protobuf definitions for gRPC
│   └── file_picker.proto
│
├── Dockerfile                           # Docker configuration for building the image
├── Makefile                             # Makefile for building and running the service
├── .env                                 # Environment variables for local development
├── go.mod                               # Go module dependencies
└── README.md                            # Project documentation
```

---

## **Deployment Instructions**

### **Prerequisites**:
- **Kubernetes Cluster**: The service is designed to run in a Kubernetes environment.
- **Docker**: Required for containerizing the service.
- **Helm** (Optional): For deploying the service with Helm.

### **Deployment Steps**:

1. **Build Docker Image**:
   Build the Docker image using the provided **Dockerfile**.
   ```bash
   docker build -t file-picker:v1 .
   ```

2. **Deploy to Kubernetes**:
   - Apply the **deployment.yaml** and **service.yaml** to deploy the service on a Kubernetes cluster.
   - Example:
     ```bash
     kubectl apply -f deployment.yaml
     kubectl apply -f service.yaml
     ```

3. **Verify Deployment**:
   Use `kubectl get pods` to ensure the pods are running, and use `kubectl logs <pod-name>` to inspect logs if necessary.

4. **Test the APIs**:
   Once the service is deployed and running, test the APIs using tools like **cURL** or **Postman**.

---

## **Environment Variables**

- **SERVER_PORT**: The port on which the service runs (default: 8080).
- **DATABASE_URL**: PostgreSQL connection string for the database.
- **GRPC_FILE_DOWNLOADER_ADDRESS**: Address of the file-downloader-service.
- **GRPC_PERMISSIONS_ADDRESS**: Address of the permissions-service.
- **GRPC_NOTIFICATION_ADDRESS**: Address of the notification-service.
- **GRPC_TRANSFORMATIONS_ADDRESS**: Address of the transformation-service.

---

## **Logging**

The service logs all critical events using a custom logger. The logs provide detailed insights into file uploads, downloads, permission changes, and transformation processes.

---

## **Contributions**

Feel free to open an issue or submit a pull request to contribute to this project. We welcome all contributions!

---