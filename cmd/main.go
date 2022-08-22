package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/q10242/go-dev-starter/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9999", r))
}
