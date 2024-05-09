package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jaffyyyy/golang-api/internal/api/structs"
)

const BASE_URL string = "https://dummyjson.com/"

var todos structs.TodoResponse

func HttpError(w http.ResponseWriter, m string, e int){
	http.Error(w, m, e)
}

// Define a handler function for the /users endpoint
func Todos(w http.ResponseWriter, _ *http.Request) { 
	// Set appropriate content type 
	w.Header().Set("Content-Type", "application/json") 

	// Get the list of users from the database
	resp, err := http.Get(BASE_URL + "todos")
	 
	if err != nil { 
		HttpError(w, "Error fetching todos", http.StatusInternalServerError)
		return
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil { 
		HttpError(w, "Error reading response body", http.StatusInternalServerError)
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
		HttpError(w, "Error marshalling todos to JSON", http.StatusInternalServerError)
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
	
	reqBody, err := io.ReadAll(r.Body); if err != nil { 
		w.WriteHeader(http.StatusInternalServerError)
		return
	} 
	
	url := BASE_URL + "todos/"+id;

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(reqBody)); if err != nil{
		HttpError(w, "Error calling NewRequest", http.StatusInternalServerError)
		return
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
 
	w.Write(content)
	defer resp.Body.Close()
			
} 
	
func DeleteTodo(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	req, err := http.NewRequest(http.MethodDelete, BASE_URL + "todos/" + id, bytes.NewBuffer([]byte("")))
	if err != nil { 
		HttpError(w, "Error calling NewRequest", http.StatusInternalServerError) 
		return
	}

	resp, err := http.DefaultClient.Do(req); if err != nil { 
		HttpError(w, "Something went wrong! Bad request!", http.StatusInternalServerError)  
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

	reqBody, err := io.ReadAll(r.Body); if err != nil{
		HttpError(w, "Error reading body", http.StatusInternalServerError)  
		return
	} 
	 
	client := &http.Client{} 
	resp, err := client.Post(BASE_URL+"todos/add", "application/json", bytes.NewBuffer(reqBody)); if err != nil {
		HttpError(w, "Error calling POST request", http.StatusInternalServerError)  
		return
	}
 
	
	// read
	content, err := io.ReadAll(resp.Body); if err != nil {
		HttpError(w, "Error reading body", http.StatusInternalServerError)  
		return
	} 
	
	w.Write(content)
	
	defer resp.Body.Close()
	  
}
