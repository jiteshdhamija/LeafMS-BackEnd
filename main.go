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
	// err := json.Unmarshal(userDatabaseContent, &userDatabase)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for i := 0; i < len(userDatabase); i++ {
	// 	user := userDatabase[i]
	// 	usersMap[user.Username] = user
	// }

	// err = json.Unmarshal(leaveDatabaseContent, &leaveDatabase)
	// if err != nil {
	// 	log.Fatal("err")
	// }

	// for i := 0; i < len(leaveDatabase); i++ {
	// 	user := leaveDatabase[i].Username
	// 	leaveMap[user] = leaveDatabase[i].Leaves
	// }

	routes := mux.NewRouter()
	routes.HandleFunc("/login", handleLogin).Methods("GET")
	routes.HandleFunc("/apply", handleApply).Methods("PUT")

	http.ListenAndServe(":8080", routes)
}
