# **File Management System**

## **Overview**

The File Management System is a microservices-based solution designed to handle file management operations such as uploads, downloads, transformations, permissions management, and real-time notifications. The system is built for scalability and high availability (HA), with all services designed to handle large volumes of traffic and failover scenarios.

This system consists of multiple services that interact with each other via both REST and gRPC APIs. User interactions are primarily handled by the **File-Picker-Service** and **Auth-Service**, while the rest of the services communicate internally using gRPC for efficiency and security.

## **Services Developed**

### 1. **Auth-Service**
- **Role**: Handles user authentication and token issuance.
- **APIs**: 
  - `POST /auth/login`: Logs in users and provides JWT tokens.
  - `POST /auth/register`: Registers new users.
  - `POST /auth/logout`: Logs out the user by invalidating the token.

### 2. **File-Picker-Service**
- **Role**: Core service for file management, handling file uploads, file listing, file downloads, and file sharing.
- **User-Facing REST APIs**:
  - `POST /api/upload`: Upload files to the system.
  - `GET /api/files`: List files accessible to the user.
  - `POST /api/files/download`: Downloads files from cloud storage to the system (triggers gRPC to **File-Downloader-Service**).
  - `POST /api/transform/:id`: Trigger a file transformation.
  - `POST /api/files/share`: Share files by adding permissions.

- **Internal gRPC APIs**: 
  - Uploading, listing, downloading, and transforming files in coordination with **File-Downloader-Service**, **File-Transformation-Service**, and **Permission-Service**.

### 3. **File-Downloader-Service** (gRPC Only)
- **Role**: Downloads files from cloud providers (e.g., Google Drive, Dropbox).
- **gRPC API**: 
  - `DownloadFile`: Downloads files from a cloud provider and stores them locally.

### 4. **File-Transformation-Service** (gRPC Only)
- **Role**: Handles file transformations such as OCR or resizing.
- **gRPC API**: 
  - `TransformFile`: Performs transformations and sends results to the **File-Picker-Service**.

### 5. **Permission-Service** (gRPC Only)
- **Role**: Manages access control for files (read, write, delete).
- **gRPC API**: 
  - `CheckPermission`: Verifies user access to files.
  - `UpdatePermission`: Updates file permissions.

### 6. **Notification-Service** (WebSocket)
- **Role**: Sends real-time notifications to users for events such as file uploads, transformations, and sharing.
- **WebSocket API**: Provides notifications for file events.

---

## **High Availability and Scalability**

All services are designed for high availability and scalability:
- **Kubernetes** manages the deployment of services, ensuring auto-scaling and fault tolerance.
- **Load balancing** distributes traffic evenly across services to prevent overload.
- **Horizontal Pod Autoscaling** in Kubernetes ensures that services can scale automatically based on traffic and resource usage.
- **Istio** manages service mesh for secure, fault-tolerant inter-service communication.

---

## **Deployment Requirements**

- **Kubernetes and Helm**: The system is deployed on Kubernetes clusters using Helm charts for service management.
- **PostgreSQL**: Required for **Auth-Service**, **File-Picker-Service**, and **Permission-Service**.
- **Redis** (optional but recommended): For caching to improve performance.
- **Istio**: For traffic management, load balancing, and circuit breaking.
- **Prometheus** and **Grafana**: For monitoring performance and usage metrics.
- **ELK Stack**: For logging and system monitoring.

---

## **Documentation**

For detailed instructions on:
- **Setting up the environment** (Kubernetes, Helm, PostgreSQL).
- **Deploying services** via Helm charts.
- **Configuring Istio** for traffic management.
- **Monitoring the system** with Prometheus and Grafana.
- **Using the application** with detailed API documentation.

Please refer to the **docs** folder in the repository.

---

This README provides an overview of the services and system architecture. For all setup, deployment, and usage details, refer to the **docs** folder.