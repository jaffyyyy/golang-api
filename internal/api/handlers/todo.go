package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jaffyyyy/golang-api/internal/api/structs"
)

const BASE_URL string = "https://dummyjson.com/"

var todos structs.TodoResponse

// Define a handler function for the /users endpoint
func Todos(w http.ResponseWriter, _ *http.Request) { 
	// Set appropriate content type 
	w.Header().Set("Content-Type", "application/json") 

	// Get the list of users from the database
	resp, err := http.Get(BASE_URL + "todos")
	 
	if err != nil {
		// Handle error gracefully, e.g., return an error response with appropriate HTTP status code
		fmt.Println("Error fetching todos: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(w, "Error reading response body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} 

	// Unmarshal the JSON response into a slice of Todo structs
	err = json.Unmarshal(body, &todos)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // Could be 400 for invalid JSON format
		return
	}


	// Marshal the slice of Todo structs back into JSON
	jsonResponse, err := json.Marshal(todos)
	if err != nil {
		fmt.Println("Error marshalling todos to JSON: ", err)
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

	ct := r.Header.Get("Content-Type")
	if ct != "" {
			mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
			if mediaType != "application/json" {
					msg := "Content-Type header is not application/json"
					http.Error(w, msg, http.StatusUnsupportedMediaType)
					return
			}
	}
	
	reqBody, err := io.ReadAll(r.Body); if err != nil { 
		w.WriteHeader(http.StatusInternalServerError)
		return
	} 
	
	url := BASE_URL + "todos/"+id;

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(reqBody)); if err != nil{
		fmt.Println("Something went wrong:", err) 
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{} 
	resp, err := client.Do(req); if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} 
 
	
	// read
	content, err := io.ReadAll(resp.Body); if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println(string(content))
	fmt.Println(string(reqBody))
 
	w.Write(content)
	defer resp.Body.Close()
			
} 
	
func DeleteTodo(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	req, err := http.NewRequest(http.MethodDelete, BASE_URL + "todos/" + id, bytes.NewBuffer([]byte("")))
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
	bodyByte, err := io.ReadAll(resp.Body); if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} 

	defer resp.Body.Close()
	w.Write(bodyByte)
 
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	ct := r.Header.Get("Content-Type")
	if ct != "" {
			mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
			if mediaType != "application/json" {
					msg := "Content-Type header is not application/json"
					http.Error(w, msg, http.StatusUnsupportedMediaType)
					return
			}
	}

	reqBody, err := io.ReadAll(r.Body); if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	} 
	 
	client := &http.Client{} 
	resp, err := client.Post(BASE_URL+"todos/add", "application/json", bytes.NewBuffer(reqBody)); if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
 
	
	// read
	content, err := io.ReadAll(resp.Body); if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} 

	w.Write(content)
	
	defer resp.Body.Close()
	  
}
