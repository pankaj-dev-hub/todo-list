package main

import (
	"fmt"
	"log"
	"net/http"
	"pankaj-katyare/todo-list/cmd/todo/auth"
	"pankaj-katyare/todo-list/internal/todo/handler"

	"github.com/gorilla/mux"
)

func RegisterHandlers(router *mux.Router) {

	router.Handle("/api/todo/create", auth.AuthenticateMiddleware(http.HandlerFunc(handler.CreateTask))).Methods("POST")
	router.Handle("/api/todo/get", auth.AuthenticateMiddleware(http.HandlerFunc(handler.GetTask))).Methods("GET")
	router.Handle("/api/todo/getAll", auth.AuthenticateMiddleware(http.HandlerFunc(handler.GetAllTask))).Methods("GET")
	router.Handle("/api/todo/update", auth.AuthenticateMiddleware(http.HandlerFunc(handler.UpdateTask))).Methods("POST")
	router.Handle("/api/todo/completed", auth.AuthenticateMiddleware(http.HandlerFunc(handler.CompletedTask))).Methods("GET")
	router.Handle("/api/todo/pending", auth.AuthenticateMiddleware(http.HandlerFunc(handler.PendingTask))).Methods("GET")

}

func main() {
	fmt.Println("Hello")

	router := mux.NewRouter()

	RegisterHandlers(router)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Starting todo service1...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
