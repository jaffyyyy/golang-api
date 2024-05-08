package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jaffyyyy/golang-api/internal/api/handlers"
)



func NewRouter(r *mux.Router){
	r.HandleFunc("/todos", handlers.Todos).Methods(http.MethodGet)
	r.HandleFunc("/todos/{id}", handlers.UpdateTodo).Methods(http.MethodPut)


	http.ListenAndServe(":8000", r)
} 