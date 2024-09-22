# Database Schema and Microservice Interactions

This document describes the database schemas for each service in the system and outlines how different microservices interact with the database. Additionally, it discusses architectural considerations for deploying the database in a Kubernetes cluster.

---

## 1. **User Table**

### Table Structure: `users`

| Column      | Type      | Description                                   |
|-------------|-----------|-----------------------------------------------|
| `id`        | `BIGINT`  | Primary key, unique user ID                   |
| `username`  | `VARCHAR` | Unique username                               |
| `email`     | `VARCHAR` | User's email, unique                          |
| `password`  | `VARCHAR` | Hashed password for the user                  |
| `role`      | `VARCHAR` | Role of the user (e.g., admin, viewer)        |
| `created_at`| `TIMESTAMP` | When the user was created                   |
| `updated_at`| `TIMESTAMP` | When the user was last updated              |

### Service Interactions:
- **Auth-Service**:
  - **Read**: Verifies user credentials during login.
  - **Write**: Registers new users.
  - **Update**: Updates user details such as password or role.
  
- **File-Picker-Service**:
  - **Read**: Fetches the user ID and role from the database for file operations.

---

## 2. **Files Table**

### Table Structure: `files`

| Column       | Type      | Description                                   |
|--------------|-----------|-----------------------------------------------|
| `id`         | `VARCHAR` | Primary key, unique file ID                   |
| `file_name`  | `VARCHAR` | Name of the file                              |
| `file_path`  | `VARCHAR` | Path where the file is stored                 |
| `owner_id`   | `BIGINT`  | Foreign key, references `users(id)`           |
| `is_shared`  | `BOOLEAN` | Whether the file is shared with other users   |
| `created_at` | `TIMESTAMP` | When the file was uploaded                  |
| `updated_at` | `TIMESTAMP` | When the file metadata was last updated     |

### Service Interactions:
- **File-Picker-Service**:
  - **Read**: Lists files owned by or shared with a user.
  - **Write**: Adds metadata for newly uploaded files.
  - **Update**: Modifies file metadata when sharing the file with other users or changing file properties.
  
- **File-Downloader-Service**:
  - **Read**: Reads the file path before initiating a file download.
  
- **File-Transformation-Service**:
  - **Read**: Retrieves file paths for transformation tasks.
  - **Write**: Updates file metadata after transformation (if necessary).

---

## 3. **Permissions Table**

### Table Structure: `permissions`

| Column         | Type      | Description                                  |
|----------------|-----------|----------------------------------------------|
| `id`           | `BIGINT`  | Primary key, unique permission record ID      |
| `user_id`      | `BIGINT`  | Foreign key, references `users(id)`           |
| `file_id`      | `VARCHAR` | Foreign key, references `files(id)`           |
| `permission`   | `VARCHAR` | Permission type (e.g., read, write, delete)   |
| `is_owner`     | `BOOLEAN` | Indicates if the user is the file owner       |
| `granted_by`   | `BIGINT`  | Foreign key, references `users(id)` who granted the permission |
| `created_at`   | `TIMESTAMP` | When the permission was granted            |

### Service Interactions:
- **File-Picker-Service**:
  - **Read**: Checks if a user has the necessary permissions to perform operations like file upload, download, or sharing.
  - **Write**: Grants new permissions to shared users when a file is shared.
  
- **Permission-Service**:
  - **Read**: Verifies user permissions for file actions.
  - **Write**: Updates permissions when a file is shared with another user.
  - **Update**: Modifies permissions when a file is no longer shared or the permission level changes (e.g., from read to write).

---

## 4. **Transformation Jobs Table** *(Planned Improvement)*

This table will be introduced in the future to track file transformations (such as OCR, image resizing, etc.) requested by users. The table will hold information about transformation jobs.

### Planned Table Structure: `transformation_jobs`

| Column             | Type      | Description                                   |
|--------------------|-----------|-----------------------------------------------|
| `id`               | `BIGINT`  | Primary key, unique job ID                    |
| `file_id`          | `VARCHAR` | Foreign key, references `files(id)`           |
| `transformation_type` | `VARCHAR` | Type of transformation (e.g., OCR, resize)  |
| `status`           | `VARCHAR` | Status of the transformation (e.g., pending, complete) |
| `created_at`       | `TIMESTAMP` | When the job was created                    |
| `updated_at`       | `TIMESTAMP` | When the job status was last updated         |

### Service Interactions:
- **File-Transformation-Service**:
  - **Read**: Retrieves job details for transformation processing.
  - **Write**: Logs new transformation jobs when a transformation is requested.
  - **Update**: Updates job status when the transformation is completed or fails.

---

## 5. **Notifications Table** *(Planned Improvement)*

This table will be added in future releases to store notifications sent to users regarding their file operations (such as file transformations, downloads, etc.).

### Planned Table Structure: `notifications`

| Column       | Type      | Description                                  |
|--------------|-----------|----------------------------------------------|
| `id`         | `BIGINT`  | Primary key, unique notification ID          |
| `user_id`    | `BIGINT`  | Foreign key, references `users(id)`          |
| `message`    | `VARCHAR` | Notification message                         |
| `status`     | `VARCHAR` | Status of the notification (e.g., sent, failed) |
| `created_at` | `TIMESTAMP` | When the notification was created          |
| `updated_at` | `TIMESTAMP` | When the notification status was updated   |

### Service Interactions:
- **Notification-Service**:
  - **Read**: Fetches pending notifications to be sent to users.
  - **Write**: Logs new notifications when actions (like file transformations) are completed.
  - **Update**: Updates the notification status (e.g., when successfully delivered).

---

## **Database Deployment in Kubernetes**

### Architectural Considerations

When deploying databases in a Kubernetes cluster, we need to carefully consider the architecture, replication, and failover mechanisms. Below are the architectural strategies and considerations for deploying the database.

### 1. **Separate DB Deployment with Multiple Pods**

The recommended approach is to deploy the database as a separate service in the Kubernetes cluster with its own **StatefulSet** or **Deployment**. This approach offers several advantages:
- **Separation of Concerns**: Microservices are stateless, whereas the database requires persistent storage. Separating them allows better management and scaling.
- **Scaling**: We can independently scale the microservices and database based on load.
- **Data Consistency**: Having one central database ensures data consistency across all services.

For the database, a **StatefulSet** is often used instead of a Deployment to maintain the stable network identity and storage of each replica.

### 2. **Clustered DB Setup**

For high availability, consider running multiple replicas of the database using tools like **PostgreSQL** or **MySQL** with clustering solutions such as:
- **Patroni** for PostgreSQL to ensure leader election and automatic failover.
- **Galera Cluster** for MySQL to enable synchronous multi-master replication.

Each pod in the StatefulSet will be assigned persistent volumes to ensure data persistence across failures. This setup ensures that:
- **Database Replicas**: There can be multiple database replicas for read scaling.
- **Failover and High Availability**: In case the primary database fails, another replica can take over as the primary.

### 3. **Database Deployment with Persistent Volumes (PV)**

For the database, Kubernetes **Persistent Volume Claims (PVCs)** should be used to allocate storage that persists beyond pod lifecycles. Each database pod will be mounted to its own persistent storage, ensuring:
- **Data Durability**: Data is not lost even if a pod is terminated or rescheduled.
- **Backup and Restore**: Persistent storage allows easier management of backups and disaster recovery plans.

### 4. **Single DB Instance per Pod (Not Recommended)**

Placing a separate database instance in each pod (i.e., each microservice has its own DB instance) can lead to:
- **Data Inconsistency**: Multiple DB instances mean maintaining data consistency between them can be difficult.
- **Replication Complexity**: Managing replication and synchronization between database instances becomes complex.
- **High Overhead**: Running multiple database instances adds unnecessary resource overhead.

### Recommended Database Architecture
Final recommendation is using a **single centralized database** with a **StatefulSet** configuration, where:
- The database is deployed as a separate service in the Kubernetes cluster.
- Multiple microservice pods communicate with this centralized database.
- Each database replica in the StatefulSet is backed by persistent volumes to ensure data persistence and durability.

---

## Conclusion

In this architecture, we ensure that:
- The database is deployed separately with high availability, scalability, and persistent storage.
- Microservices are stateless and scale independently.
- Data remains consistent, available, and durable across the system.