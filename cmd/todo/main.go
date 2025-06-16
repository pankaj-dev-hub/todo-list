package main

import (
	"fmt"
	"log"
	"net/http"
	"pankaj-dev-hub/todo-list/internal/todo"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello")

	router := mux.NewRouter()

	todo.RegisterHandlers(router)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Starting todo service1...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
