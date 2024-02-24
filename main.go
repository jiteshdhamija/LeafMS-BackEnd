package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// var userDatabase []User
// var userDatabaseContent, _ = os.ReadFile("./usersDatabase.json")

// var leaveDatabase []Leave
// var leaveDatabaseContent, _ = os.ReadFile("./leaveDatabase.json")

// var usersMap = map[string]User{}
// var leaveMap = map[string][]LeaveSpan{}

func main() {

	routes := mux.NewRouter()
	routes.HandleFunc("/login", handleLogin).Methods("GET")
	routes.HandleFunc("/apply", handleApply).Methods("PUT")

	http.ListenAndServe(":8080", routes)
}
