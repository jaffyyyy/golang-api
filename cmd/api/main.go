package main

import (
	"github.com/gorilla/mux"
	"github.com/jaffyyyy/golang-api/internal/api/routes"
)
 

func main() {
	r := mux.NewRouter() 
	routes.NewRouter(r)
}