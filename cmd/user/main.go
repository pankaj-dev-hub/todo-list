package main

import (
	"fmt"
	"log"
	"net/http"
	"pankaj-katyare/todo-list/internal/user"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello")

	router := mux.NewRouter()

	user.RegisterHandlers(router)

	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	log.Println("Starting User service1...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
