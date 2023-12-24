package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Username string `json:"username"`
	PassWord string `json:"password"`
}

func main() {
	routes := mux.NewRouter()
	routes.HandleFunc("/login", handleLogin).Methods("GET")

	http.ListenAndServe(":8080", routes)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	var user User
	fmt.Println("Service is responding")
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(errors.New(
			"Could not parse/decode the json request. Make sure that the json is properly formatted and has all the required fields."))
	}

	//Authenticate the user credentials with the database
	//validateCred(user);

	response, _ := json.Marshal(user)
	w.Write(response)
}
