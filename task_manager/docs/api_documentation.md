# Task Management API Documentation

This document provides documentation for the Task Management REST API.

## Base URL
`http://localhost:8080`

---

## Endpoints

### 1. Get All Tasks

*   **URL:** `/tasks`
*   **Method:** `GET`
*   **Description:** Retrieves a list of all tasks.
*   **Success Response:**
    *   **Code:** 200 OK
    *   **Content:**
        ```json
        [
            {
                "id": "f47ac10b-58cc-4372-a567-0e02b2c3d479",
                "title": "Finish Project Proposal",
                "description": "Write and submit the project proposal for Q4.",
                "due_date": "2025-12-15",
                "status": "In Progress"
            }
        ]
        ```

### 2. Get a Specific Task

*   **URL:** `/tasks/:id`
*   **Method:** `GET`
*   **Description:** Retrieves the details of a specific task by its ID.
*   **URL Params:**
    *   `id` (string, required): The ID of the task.
*   **Success Response:**
    *   **Code:** 200 OK
    *   **Content:**
        ```json
        {
            "id": "f47ac10b-58cc-4372-a567-0e02b2c3d479",
            "title": "Finish Project Proposal",
            "description": "Write and submit the project proposal for Q4.",
            "due_date": "2025-12-15",
            "status": "In Progress"
        }
        ```
*   **Error Response:**
    *   **Code:** 404 Not Found
    *   **Content:**
        ```json
        {
            "error": "Task not found"
        }
        ```

### 3. Create a New Task

*   **URL:** `/tasks`
*   **Method:** `POST`
*   **Description:** Creates a new task.
*   **Request Body:**
    ```json
    {
        "title": "Buy Groceries",
        "description": "Milk, Bread, Cheese",
        "due_date": "2025-11-10",
        "status": "To Do"
    }
    ```
*   **Success Response:**
    *   **Code:** 201 Created
    *   **Content:**
        ```json
        {
            "id": "a8c7b6e5-4d3c-2b1a-0f9e-8d7c6b5a4b3c",
            "title": "Buy Groceries",
            "description": "Milk, Bread, Cheese",
            "due_date": "2025-11-10",
            "status": "To Do"
        }
        ```

### 4. Update a Task

*   **URL:** `/tasks/:id`
*   **Method:** `PUT`
*   **Description:** Updates an existing task.
*   **URL Params:**
    *   `id` (string, required): The ID of the task to update.
*   **Request Body:**
    ```json
    {
        "title": "Buy Groceries and Fruits",
        "description": "Milk, Bread, Cheese, Apples",
        "due_date": "2025-11-11",
        "status": "In Progress"
    }
    ```
*   **Success Response:**
    *   **Code:** 200 OK
    *   **Content:**
        ```json
        {
            "id": "a8c7b6e5-4d3c-2b1a-0f9e-8d7c6b5a4b3c",
            "title": "Buy Groceries and Fruits",
            "description": "Milk, Bread, Cheese, Apples",
            "due_date": "2025-11-11",
            "status": "In Progress"
        }
        ```
*   **Error Response:**
    *   **Code:** 404 Not Found
    *   **Content:**
        ```json
        {
            "error": "Task not found"
        }
        ```

### 5. Delete a Task

*   **URL:** `/tasks/:id`
*   **Method:** `DELETE`
*   **Description:** Deletes a specific task.
*   **URL Params:**
    *   `id` (string, required): The ID of the task to delete.
*   **Success Response:**
    *   **Code:** 200 OK
    *   **Content:**
        ```json
        {
            "message": "Task deleted successfully"
        }
        ```
*   **Error Response:**
    *   **Code:** 404 Not Found
    *   **Content:**
        ```json
        {
            "error": "Task not found"
        }
        ```