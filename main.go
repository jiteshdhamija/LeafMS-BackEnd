package main

import (
	"encoding/json"
	"fmt"
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

var payload []User
var content, _ = os.ReadFile("./users.json")

var usersMap = map[string]User{}

func main() {
	err := json.Unmarshal(content, &payload)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(payload); i++ {
		user := payload[i]
		usersMap[user.Username] = user
	}

	routes := mux.NewRouter()
	routes.HandleFunc("/login", handleLogin).Methods("GET")

	http.ListenAndServe(":8080", routes)
}

// handle login
func handleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Service is responding")

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request payload!!!"))
		return
	}

	//Authenticate the user credentials with the database
	user = validateCred(usersMap, user)
	if (User{} == user) {
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
