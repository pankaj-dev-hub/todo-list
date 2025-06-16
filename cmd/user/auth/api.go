package auth

import (
	"pankaj-dev-hub/todo-list/cmd/todo/auth"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	subRouter := r.PathPrefix("/user").Subrouter()
	subRouter.Use(auth.AuthenticateMiddleware)

	//cmd.RegisterHandlers(subRouter)
}
