package user

import (
	"net/http"
	"pankaj-katyare/todo-list/cmd/todo/auth"
	"pankaj-katyare/todo-list/internal/user/handler"

	"github.com/gorilla/mux"
)

func RegisterHandlers(router *mux.Router) {

	router.HandleFunc("/api/user/login", handler.LoginUser).Methods("POST")
	router.HandleFunc("/api/user/create", handler.CreateUser).Methods("POST")
	router.Handle("/api/user/get", auth.AuthenticateMiddleware(http.HandlerFunc(handler.GetUser))).Methods("GET")
	router.Handle("/api/user/getAll", auth.AuthenticateMiddleware(http.HandlerFunc(handler.GetAllUser))).Methods("GET")
	router.Handle("/api/user/update", auth.AuthenticateMiddleware(http.HandlerFunc(handler.UpdateUser))).Methods("POST")
}
