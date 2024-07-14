package main

import (
	"net/http"

	ctrlr "./"

	"github.com/gorilla/mux"
)

func main() {

	//uncomment the following method only when you want to insert public holidays of a country for a specific year
	// utils.PersistPublicHolidays(2024, "IN")

	routes := mux.NewRouter()
	routes.HandleFunc("/login", ctrlr.HandleLogin).Methods("GET")
	authRoute := routes.NewRoute().Subrouter()
	authRoute.Use(ctrlr.HandleAuth)
	authRoute.HandleFunc("/apply", ctrlr.HandleApply).Methods("PUT")
	authRoute.HandleFunc("/leaves", ctrlr.HandleViewLeaves).Methods("GET")
	authRoute.HandleFunc("/applications", ctrlr.HandleViewLeaveApplications).Methods(("GET"))
	authRoute.HandleFunc("/approve", ctrlr.HandleLeaveApproval).Methods("PATCH")
	authRoute.HandleFunc("/holidays", ctrlr.HandleViewHolidays).Methods("GET")

	http.ListenAndServe(":8080", routes)
}
