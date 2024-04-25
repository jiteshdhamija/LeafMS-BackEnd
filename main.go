package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	routes := mux.NewRouter()
	routes.HandleFunc("/login", handleLogin).Methods("GET")
	authRoute := routes.NewRoute().Subrouter()
	authRoute.Use(handleAuth)
	authRoute.HandleFunc("/apply", handleApply).Methods("PUT")
	authRoute.HandleFunc("/leaves", handleViewLeaves).Methods("GET")
	authRoute.HandleFunc("/applications", handleViewLeaveApplications).Methods(("GET"))
	authRoute.HandleFunc("/aprrove", handleLeaveApproval).Methods("PUT")

	http.ListenAndServe(":8080", routes)
}
