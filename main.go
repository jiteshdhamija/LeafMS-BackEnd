package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	//uncomment the following method only when you want to insert public holidays of a country for a specific year
	// utils.PersistPublicHolidays(2024, "IN")

	routes := mux.NewRouter()
	routes.HandleFunc("/login", handleLogin).Methods("GET")
	authRoute := routes.NewRoute().Subrouter()
	authRoute.Use(handleAuth)
	authRoute.HandleFunc("/apply", handleApply).Methods("PUT")
	authRoute.HandleFunc("/leaves", handleViewLeaves).Methods("GET")
	authRoute.HandleFunc("/applications", handleViewLeaveApplications).Methods(("GET"))
	authRoute.HandleFunc("/approve", handleLeaveApproval).Methods("PATCH")

	http.ListenAndServe(":8080", routes)
}
