package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jaaaaaaaaaam/go-todo/models"
)

// Index shows the index
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

// TodoIndex shows the todo-list
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := todo.Todos{
		todo.Todo{Name: "Learn Golang"},
		todo.Todo{Name: "Write some awesome code"},
	}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

// TodoShow shows a single todo
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoID"]
	fmt.Fprintln(w, "Todo show:", todoID)
}
