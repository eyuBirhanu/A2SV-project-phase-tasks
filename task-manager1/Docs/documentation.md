# Task Management API - Clean Architecture Refactoring

## 1. Overview
This project represents a refactoring of the legacy Task Management API into a **Clean Architecture** structure. The primary goal was to decouple business logic from external frameworks (like Gin) and data storage mechanisms, ensuring high maintainability, testability, and scalability.

## 2. Architectural Design Decisions
The application follows the **Dependency Rule**, where dependencies only point inwards. The inner layers know nothing about the outer layers.

### Layers Overview

1.  **Domain Layer (`Domain/`)**
    *   **Role:** The core of the application.
    *   **Contents:** Pure struct definitions (`Task`, `User`).
    *   **Dependency:** No dependencies. It does not know about JSON, HTTP, or Databases.

2.  **Usecases Layer (`Usecases/`)**
    *   **Role:** Application-specific business rules.
    *   **Contents:** Methods like `Register`, `CreateTask`. It validates input and enforces rules (e.g., "Default status is Pending").
    *   **Dependency:** Depends only on `Domain` and `Repository Interfaces`. It relies on Dependency Injection to function.

3.  **Repositories Layer (`Repositories/`)**
    *   **Role:** Interface Adapters for data access.
    *   **Contents:** Interfaces defining *what* data operations are possible, and concrete implementations (currently In-Memory, adaptable to MongoDB).
    *   **Dependency:** Depends on `Domain`.

4.  **Infrastructure Layer (`Infrastructure/`)**
    *   **Role:** Frameworks and Drivers.
    *   **Contents:** JWT generation (`jwt-go`), Password hashing (`bcrypt`), and Auth Middleware.
    *   **Dependency:** Independent utilities injected into Usecases or Delivery.

5.  **Delivery Layer (`Delivery/`)**
    *   **Role:** External entry point.
    *   **Contents:** REST API Controllers (`Gin`), Routing, and `main.go` wiring.
    *   **Dependency:** Depends on all inner layers to orchestrate the flow.

## 3. How to Run

### Prerequisites
*   Go 1.18+

### Steps
1.  **Install Dependencies:**
    ```bash
    go mod tidy
    ```
2.  **Run the Server:**
    ```bash
    go run Delivery/main.go
    ```
3.  **API Endpoints:**
    *   `POST /api/register` - Create a user.
    *   `POST /api/login` - Get JWT Token.
    *   `GET /api/tasks` - List tasks (Requires Bearer Token).
    *   `POST /api/tasks` - Create task (Requires Bearer Token).

## 4. Future Improvements
*   **Persistence:** Replace `InMemoryTaskRepo` with `MongoTaskRepo` by simply implementing the Repository interface.
*   **Logging:** Introduce a structured logging interface in the Infrastructure layer.