package main

import (
	"fmt"
	"log"
	"net/http"
	"pankaj-katyare/todo-list/internal/todo/handler"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello")

	router := mux.NewRouter()

	router.HandleFunc("/api/todo/create", handler.CreateTask).Methods("POST")
	router.HandleFunc("/api/todo/get", handler.GetTask).Methods("GET")
	router.HandleFunc("/api/todo/getAll", handler.GetAllTask).Methods("GET")
	router.HandleFunc("/api/todo/update", handler.UpdateTask).Methods("POST")
	router.HandleFunc("/api/todo/completed", handler.CompletedTask).Methods("GET")
	router.HandleFunc("/api/todo/pending", handler.PendingTask).Methods("GET")

	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	log.Println("Starting todo service1...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
