package main

import (
	"fmt"
	"log"
	"net/http"
	"pankaj-katyare/todo-list/internal/user/handler"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello")

	router := mux.NewRouter()

	router.HandleFunc("/api/user/create", handler.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/get", handler.GetUser).Methods("GET")
	router.HandleFunc("/api/user/getAll", handler.GetAllUser).Methods("GET")
	router.HandleFunc("/api/user/update", handler.UpdateUser).Methods("POST")

	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	log.Println("Starting User service1...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
