package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var userDatabase []User
var userDatabaseContent, _ = os.ReadFile("./usersDatabase.json")

var leaveDatabase []Leave
var leaveDatabaseContent, _ = os.ReadFile("./leaveDatabase.json")

// write the databases in these hashmaps
var usersMap = map[string]User{}
var leaveMap = map[string][]LeaveSpan{}

// app starts here
func main() {

	//(unwrap/deserialize) json file database content
	err := json.Unmarshal(userDatabaseContent, &userDatabase)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(userDatabase); i++ {
		user := userDatabase[i]
		usersMap[user.Username] = user
	}

	err = json.Unmarshal(leaveDatabaseContent, &leaveDatabase)
	if err != nil {
		log.Fatal("err")
	}

	for i := 0; i < len(leaveDatabase); i++ {
		user := leaveDatabase[i].Username
		leaveMap[user] = leaveDatabase[i].Leaves
	}

	//add routes and start the server
	routes := mux.NewRouter()
	routes.HandleFunc("/login", handleLogin).Methods("GET")
	routes.HandleFunc("/apply", handleApply).Methods("PUT")
	routes.HandleFunc("/aprrove", handleLeaves).Methods("PATCH")
	routes.HandleFunc("/approve", handleLeaves).Methods("GET")

	fmt.Println("Server is live!!! at localhost:8080")
	http.ListenAndServe(":8080", routes)
}
