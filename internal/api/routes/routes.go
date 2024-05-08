package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jaffyyyy/golang-api/internal/api/handlers"
)



func NewRouter(r *mux.Router){
	r.HandleFunc("/todos", handlers.Todos).Methods(http.MethodGet)
	r.HandleFunc("/todos/{id}", handlers.UpdateTodo).Methods(http.MethodPut)
	r.HandleFunc("/todos/{id}", handlers.DeleteTodo).Methods(http.MethodDelete)
	r.HandleFunc("/todos/add", handlers.CreateTodo).Methods(http.MethodPost) 

	http.ListenAndServe(":8000", r)
} 