package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type User struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	Team        string `json:"team"`
	Designation string `json:"designation"`
	Approver    string `json:"approver"`
}

type LeaveSpan struct {
	Start string `json:"startTime"`
	End   string `json:"endTime"`
}

type Leave struct {
	Username string      `json:"username"`
	Leaves   []LeaveSpan `json:"leaves"`
}

var userDatabase []User
var userDatabaseContent, _ = os.ReadFile("./usersDatabase.json")

var leaveDatabase []Leave
var leaveDatabaseContent, _ = os.ReadFile("./leaveDatabase.json")

var usersMap = map[string]User{}
var leaveMap = map[string][]LeaveSpan{}

func main() {

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

	routes := mux.NewRouter()
	routes.HandleFunc("/login", handleLogin).Methods("GET")
	routes.HandleFunc("/apply", handleApply).Methods("PUT")

	http.ListenAndServe(":8080", routes)
}

// handle login
func handleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Service is responding")

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request userDatabase!!!"))
		return
	}

	//Authenticate the user credentials with the database
	user = validateCred(usersMap, user)
	if (User{} == user) {
		log.Fatal(err)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid Credentials!!"))
		return
	}

	response, _ := json.Marshal(user)
	w.Write(response)
}

// function to validate the user
func validateCred(userList map[string]User, userToAuthorize User) User {
	user, found := userList[userToAuthorize.Username]
	if found && user.Password == userToAuthorize.Password {
		return user
	}

	return User{}
}

// handle leave application
func handleApply(w http.ResponseWriter, r *http.Request) {
	var leaveApplication Leave
	err := json.NewDecoder(r.Body).Decode(&leaveApplication)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	leaveToAppend := LeaveSpan{
		Start: leaveApplication.Leaves[0].Start,
		End:   leaveApplication.Leaves[0].End,
	}

	for i := 0; i < len(leaveDatabase); i++ {
		if leaveDatabase[i].Username == leaveApplication.Username {
			leaveDatabase[i].Leaves = append(leaveDatabase[i].Leaves, leaveToAppend)
			break
		}
	}
	leaveDatabaseContent, _ := json.Marshal(leaveDatabase)
	if err := os.WriteFile("leaveDatabase.json", leaveDatabaseContent, 0666); err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	response, _ := json.Marshal(leaveDatabaseContent)
	w.Write(response)
}
