package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// handle login
func handleLogin(w http.ResponseWriter, r *http.Request) {

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

	response, _ := json.MarshalIndent(user, "", "  ")
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

	for i := 0; i < len(leaveDatabase); i++ {
		if leaveDatabase[i].Username == leaveApplication.Username {
			leaveDatabase[i].Leaves = append(leaveDatabase[i].Leaves, leaveApplication.Leaves...)
			break
		}
	}
	leaveDatabaseContent, _ := json.MarshalIndent(leaveDatabase, "", "  ")
	if err := os.WriteFile("leaveDatabase.json", leaveDatabaseContent, 0666); err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	response, _ := json.MarshalIndent(leaveDatabaseContent, "", "  ")
	w.Write(response)
}

// approve one or multiple leave requests
func handleLeaves(w http.ResponseWriter, r *http.Request) {
	requestMethod := r.Method

}

func handleViewLeaves(w http.ResponseWriter, r *http.Request) {

}
