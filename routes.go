package main

import (
	"encoding/json"
	"log"
	"net/http"

	db "github.com/jiteshdhamija/LeafMS-BackEnd/database"
	"go.mongodb.org/mongo-driver/bson"
)

// type User db.User
// type Leave db.Leave
// type LeaveSpan db.LeaveSpan
// type Database db.Database

var database = db.ConnectDB()

// function to validate the db.user
func validateCred(userToAuthorize db.User) []db.User {
	user, err := database.Find("employees", bson.D{
		{Key: "username", Value: userToAuthorize.Username},
		{Key: "password", Value: userToAuthorize.Password}})

	if err != nil {
		log.Fatal("Failed authentication. Error:- \n\t", err)
		return nil
	}

	return user
}

// handle login
func handleLogin(w http.ResponseWriter, r *http.Request) {

	var user db.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request userDatabase!!!"))
		return
	}

	//Authenticate the user credentials with the database
	result := validateCred(user)
	if result == nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid Credentials!!"))
		return
	}

	response, _ := json.Marshal(result)
	w.Write(response)
}

// handle leave application
func handleApply(w http.ResponseWriter, r *http.Request) {
	var leaveApplication db.Leave
	err := json.NewDecoder(r.Body).Decode(&leaveApplication)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// leaveToAppend := LeaveSpan{
	// 	Start: leaveApplication.Leaves[0].Start,
	// 	End:   leaveApplication.Leaves[0].End,
	// }

	// for i := 0; i < len(leaveDatabase); i++ {
	// 	if leaveDatabase[i].db.Username == leaveApplication.db.Username {
	// 		leaveDatabase[i].Leaves = append(leaveDatabase[i].Leaves, leaveToAppend)
	// 		break
	// 	}
	// }
	// leaveDatabaseContent, _ := json.Marshal(leaveDatabase)
	// if err := os.WriteFile("leaveDatabase.json", leaveDatabaseContent, 0666); err != nil {
	// 	log.Fatal(err)
	// }

	w.WriteHeader(http.StatusCreated)
	response, _ := json.Marshal(leaveApplication)
	w.Write(response)
}
