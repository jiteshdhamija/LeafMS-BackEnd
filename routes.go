package main

import (
	"encoding/json"
	"log"
	"net/http"

	db "LeafMS-BackEnd/database"
	util "LeafMS-BackEnd/utils"

	"go.mongodb.org/mongo-driver/bson"
)

// type User db.User
// type Leave db.Leave
// type LeaveSpan db.LeaveSpan
// type Database db.Database

var database = db.ConnectDB()

// function to validate the db.user
func validateCred(userToAuthorize db.User) interface{} {
	var login db.UserLogin
	user, err := database.Find("employees", bson.D{
		{Key: "username", Value: userToAuthorize.Username},
		{Key: "password", Value: userToAuthorize.Password}})
	if err != nil {
		log.Fatal("Failed authentication. Error:- \n\t", err)
		login.Login = false
		return login
	}
	var user2 = util.InterFaceToUser(user)
	if user2.Username == "" {
		login.Login = false
		return login
	} else {
		login.Username = user2.Username
		login.Login = true
	}
	return login
}

// handle login
func handleLogin(w http.ResponseWriter, r *http.Request) {
	var user db.User
	log.Println("started login api")
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request userDatabase!!!"))
		return
	}

	//Authenticate the user credentials with the database
	result := validateCred(user).(db.UserLogin)
	log.Println("validated cred")

	response, _ := json.MarshalIndent(result, "", "	")
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

	result, err := database.UpdateOne("leaves", bson.D{
		{Key: "username", Value: leaveApplication.Username},
	}, bson.D{
		{Key: "$push", Value: bson.D{
			{Key: "leaves", Value: bson.D{
				{Key: "$each", Value: leaveApplication.Leaves},
			}},
		}},
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		response, _ := json.Marshal("No User with the username: " + leaveApplication.Username + " exists.")
		w.Write(response)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}

	response, _ := json.MarshalIndent(result, "", "	")
	w.Write(response)
}

// handle `view leaves`
func handleViewLeaves(w http.ResponseWriter, r *http.Request) {
	var user db.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := database.Find("leaves", bson.D{
		{Key: "username", Value: user.Username},
	})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// var leaves []db.Leave
	// for i := 0; i < len(result); i++ {
	// 	leave, ok := result[i].(db.Leave)
	// 	if !ok {
	// 		log.Fatal("Interface to Leave struct conversion failed")
	// 	} else {
	// 		leaves = append(leaves, leave)
	// 	}
	// }

	response, _ := json.MarshalIndent(result, "", "	")
	w.Write(response)

}
