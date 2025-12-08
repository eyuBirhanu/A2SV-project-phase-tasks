package domain

import "time"

// Task represents a task in the system
type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"` // e.g., "Pending", "In Progress", "Done"
	DueDate     time.Time `json:"due_date"`
	UserID      string    `json:"user_id"`
}

// User represents a user in the system
type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
