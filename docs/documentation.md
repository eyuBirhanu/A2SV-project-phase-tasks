# Library Management System Documentation

## 1. Overview

This project is a simple, console-based library management system built using the Go programming language. It serves as a practical demonstration of core Go concepts, including structs, interfaces, methods, maps, and slices.

The system allows a user to perform basic library operations through a command-line interface, such as adding, removing, borrowing, and returning books.

---

## 2. Folder Structure

The project is organized into a clean, modular structure to separate concerns, making the code easier to maintain and understand.
library_management/
├── main.go # The entry point of the application.
├── go.mod # Manages project dependencies.
├── controllers/
│ └── library_controller.go # Handles user input and console interaction.
├── models/
│ ├── book.go # Defines the Book struct.
│ └── member.go # Defines the Member struct.
├── services/
│ └── library_service.go # Contains all the core business logic.
└── docs/
└── documentation.md # Contains this documentation file.
code
Code
-   **`main.go`**: Initializes the services and controllers and starts the application.
-   **`models/`**: Defines the data structures (`Book`, `Member`) used throughout the application.
-   **`services/`**: Implements the business logic for managing the library's data. It is completely decoupled from the user interface.
-   **`controllers/`**: Acts as the bridge between the user and the application's services. It reads console input and calls the appropriate service methods.
-   **`docs/`**: Contains project documentation.

---

## 3. How to Run the Application

To run the application, you need to have Go installed on your system.

1.  **Navigate to the Project Directory**:
    Open your terminal and navigate to the root folder of the project.
    ```sh
    cd path/to/library_management
    ```

2.  **Run the Application**:
    Execute the following command:
    ```sh
    go run .
    ```

This command will compile and run the `main.go` file, starting the application and displaying the interactive menu in your terminal.