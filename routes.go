package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// function to validate the user
func validateCred(userList map[string]User, userToAuthorize User) User {
	user, found := userList[userToAuthorize.Username]
	if found && user.Password == userToAuthorize.Password {
		return user
	}

	return User{}
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
