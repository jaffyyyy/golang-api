package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jaffyyyy/golang-api/internal/api/handlers"
)



func NewRouter(r *mux.Router){
	r.HandleFunc("/api/todos", handlers.Todos).Methods(http.MethodGet)
	r.HandleFunc("/api/todos/{id}", handlers.UpdateTodo).Methods(http.MethodPut)
	r.HandleFunc("/api/todos/{id}", handlers.DeleteTodo).Methods(http.MethodDelete)
	r.HandleFunc("/api/todos/add", handlers.CreateTodo).Methods(http.MethodPost) 

	http.ListenAndServe(":8000", r)
} 