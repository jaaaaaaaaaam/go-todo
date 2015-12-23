package todo

import "time"

// Todo is a struct for todo items
type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos is
type Todos []Todo
