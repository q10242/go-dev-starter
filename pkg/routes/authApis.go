package routes

import (
	"github.com/gorilla/mux"
	"github.com/q10242/go-dev-starter/pkg/controllers"
)

var RegisterAuthApiRoutes = func(router *mux.Router) {
	prefix := "/auth"
	router.HandleFunc(prefix+"/register", controllers.Register).Methods("POST")
	router.HandleFunc(prefix+"/login", controllers.Login).Methods("POST")
	router.HandleFunc(prefix+"/logout", controllers.Logout).Methods("POST")
	router.HandleFunc(prefix+"/listAllUsers", controllers.ListAllUsers).Methods("GET")
	router.HandleFunc(prefix+"/issueToken", controllers.IssueToken).Methods("POST")
	router.HandleFunc(prefix+"/renewToken", controllers.RenewToken).Methods("POST")
	router.HandleFunc(prefix+"/user/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc(prefix+"/user/{id}", controllers.EditUser).Methods("PATCH")

}
