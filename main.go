package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

var loggedInToken map[string]string

func main() {

	routes := mux.NewRouter()
	routes.HandleFunc("/login", handleLogin).Methods("GET")
	routes.Use(handleAuth)
	routes.HandleFunc("/apply", handleApply).Methods("PUT")
	routes.HandleFunc("/leaves", handleViewLeaves).Methods("GET")
	routes.HandleFunc("/aprrove", handleLeaveApproval).Methods("PUT")

	http.ListenAndServe(":8080", routes)
}
