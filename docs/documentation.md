# Library Management System Documentation

## 1. Overview
This project is an advanced Library Management System written in Go. It demonstrates core language concepts including Structs, Interfaces, Maps, and specialized concurrency features: **Goroutines, Channels, and Mutexes**.

## 2. Concurrency Features
This version introduces concurrent handling of book reservations to simulate a real-world high-load environment.

### A. Mutex (Sync.Mutex)
*   **Location:** `services/library_service.go`
*   **Purpose:** The `Library` struct now includes a `sync.Mutex`.
*   **Usage:** Every method that accesses the `Books` or `Members` maps (Read or Write) locks the mutex (`l.mu.Lock()`) and defers unlocking (`defer l.mu.Unlock()`). This prevents **Race Conditions** where two users might try to modify a book's status at the exact same nanosecond.

### B. Goroutines & Channels (Worker Pool)
*   **Location:** `concurrency/reservation_worker.go`
*   **Mechanism:**
    1.  The `ReservationHandler` creates a buffered **Channel** (`chan ReservationRequest`).
    2.  When a user selects "Reserve a Book", the request is not processed immediately by the main thread. Instead, it is sent to the channel.
    3.  **Worker Goroutines:** We start multiple background workers (defined in `StartWorkers`). These workers constantly listen to the channel. When a message arrives, a worker picks it up and calls `ReserveBook`.
*   **Benefit:** This allows the UI to remain responsive even if reservation processing takes time, and handles bursts of requests efficiently.

### C. Auto-Cancellation (Timer Goroutine)
*   **Location:** `services/library_service.go` -> `ReserveBook`
*   **Mechanism:** When a book is successfully reserved, a dedicated anonymous Goroutine is spawned using `go func()`.
*   **Logic:**
    1.  The Goroutine sleeps for 5 seconds (`time.Sleep`).
    2.  It wakes up, locks the mutex, and checks if the book is still "Reserved".
    3.  If it hasn't been borrowed yet, it resets the status to "Available".

## 3. How to Run
1.  Navigate to root: `cd library_management`
2.  Run: `go run .`
3.  Test Concurrency:
    *   Select Option 7 to Reserve a book.
    *   Observe the "Reserved" status using Option 5.
    *   Wait 5 seconds without borrowing.
    *   Observe the console message indicating the reservation expired.
    *   Check Option 5 again; it should be "Available".