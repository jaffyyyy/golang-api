package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jaffyyyy/golang-api/internal/api/handlers"
	"github.com/rs/cors"
)



func NewRouter(r *mux.Router){
	r.HandleFunc("/api/todos", handlers.Todos).Methods(http.MethodGet)
	r.HandleFunc("/api/todos/{id}", handlers.UpdateTodo).Methods(http.MethodPut)
	r.HandleFunc("/api/todos/{id}", handlers.DeleteTodo).Methods(http.MethodDelete)
	r.HandleFunc("/api/todos/add", handlers.CreateTodo).Methods(http.MethodPost) 

	origins := []string{"*"} 
	
	c := cors.New(cors.Options{
		AllowedOrigins:   origins,
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"X-Execution-ID"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	http.ListenAndServe(":8000", handler)
} 