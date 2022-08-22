package routes

import (
	"github.com/gorilla/mux"
	"github.com/q10242/go-dev-starter/pkg/controllers"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/", controllers.Home).Methods("GET")

}
