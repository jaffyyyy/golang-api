package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jaffyyyy/golang-api/internal/api/structs"
)

const BASE_URL string = "https://dummyjson.com"

var todos structs.TodoResponse

// Define a handler function for the /users endpoint
func Todos(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set appropriate content type
	// Get the list of users from the database
	resp, err := http.Get(BASE_URL + "/todos")

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
	defer resp.Body.Close()

	// Write the JSON response to the ResponseWriter
	w.Write(jsonResponse)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	
	reqBody, err := io.ReadAll(r.Body)
	if err != nil { 
		w.WriteHeader(http.StatusInternalServerError)
		return
		} 
		
		req, err := http.NewRequest(http.MethodPut, BASE_URL+ "/todos/" + id, bytes.NewBuffer(reqBody)) 
		if err != nil{
			fmt.Fprintf(w, "Something went wrong: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("client: error making http request: %s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		
		
		// Read the response body
		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Fprintf(w, "Error reading response body: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
			} 
			
			var todo *structs.Todo
			// Unmarshal the JSON response into a slice of Todo structs
			err = json.Unmarshal(body, &todo)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest) // Could be 400 for invalid JSON format
				return
			}
			
			// Marshal the slice of Todo structs back into JSON
			jsonResponse, err := json.Marshal(todo)
			if err != nil {
				fmt.Fprintf(w, "Error marshalling todos to JSON: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			
			defer res.Body.Close()
			
			// Write the JSON response to the ResponseWriter
			w.Write(jsonResponse)
			
		} 
		
		func DeleteTodo(w http.ResponseWriter, r *http.Request){
			w.Header().Set("Content-Type", "application/json")
			vars := mux.Vars(r)
			id := vars["id"]

			req, err := http.NewRequest(http.MethodDelete, BASE_URL + "/todos/" + id, bytes.NewBuffer([]byte("")))
			if err != nil {
				fmt.Printf("Something went wrong: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			resp, err := http.DefaultClient.Do(req); if err != nil {
				fmt.Printf("Something went wrong: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// read response body
			bodyBytes, err := io.ReadAll(resp.Body); if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			} 

			// unmarshal and store to Todo struct
			jsonResponse, err := json.Marshal(bodyBytes); if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			defer resp.Body.Close()

			w.WriteHeader(http.StatusCreated)
			w.Write(jsonResponse)

		}