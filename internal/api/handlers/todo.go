package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jaffyyyy/golang-api/internal/api/structs"
)

const URL string = "https://dummyjson.com/todos"

var todos structs.TodoResponse

// Define a handler function for the /users endpoint
func Todos(w http.ResponseWriter, _ *http.Request) {
	// Get the list of users from the database
	resp, err := http.Get(URL)

	if err != nil {
		// Handle error gracefully, e.g., return an error response with appropriate HTTP status code
		fmt.Printf("Error fetching todos: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(w, "Error reading response body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return the slice of Todo structs in the response body
	w.Header().Set("Content-Type", "application/json") // Set appropriate content type

	// Unmarshal the JSON response into a slice of Todo structs
	err = json.Unmarshal(body, &todos)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // Could be 400 for invalid JSON format
		return
	}

	// Marshal the slice of Todo structs back into JSON
	jsonResponse, err := json.Marshal(todos)
	if err != nil {
		fmt.Fprintf(w, "Error marshalling todos to JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write the JSON response to the ResponseWriter
	w.Write(jsonResponse)
	defer resp.Body.Close()
}