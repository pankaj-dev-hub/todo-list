package auth

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	subRouter := r.PathPrefix("/todo").Subrouter()
	subRouter.Use(AuthenticateMiddleware)

	//cmd.RegisterHandlers(subRouter)
}
