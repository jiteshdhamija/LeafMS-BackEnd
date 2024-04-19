package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	db "LeafMS-BackEnd/database"
	util "LeafMS-BackEnd/utils"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
)

var database = db.ConnectDB()

// generate JWT token
func generateJWT(username string) (string, error) {
	secretKey := []byte("jiteshmc" + username)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string, username string) error {
	secretKey := []byte("jiteshmc" + username)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

// function to validate the db.user
func validateCred(userToAuthorize db.User) interface{} {
	var login db.UserLogin
	userInterface, err := database.Find("employees", bson.D{
		{Key: "username", Value: userToAuthorize.Username},
		{Key: "password", Value: userToAuthorize.Password}})
	if err != nil {
		log.Fatal("Failed authentication. Error:- \n\t", err)
		login.Login = false
		return login
	}
	var user = util.InterFaceToUser(userInterface)
	if user.Username == "" {
		login.Login = false
		return login
	} else {
		login.Username = user.Username
		login.Login = true
	}

	return login
}

// handle login
func handleLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var user db.User

	log.Println("started login api")
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request payload!!!"))
		return
	}

	//Authenticate the user credentials with the database
	result := validateCred(user).(db.UserLogin)
	log.Println("validated cred")

	jwtToken, err := generateJWT(user.Username)
	if err != nil {
		log.Printf("couldn't generate JWT auth token.\nError: %v", err)
	}
	w.Header().Add("Authorization", jwtToken)

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

	response, _ := json.MarshalIndent(result, "", "	")
	w.Write(response)
}

// handle leaves approval
func handleLeaveApproval(w http.ResponseWriter, r *http.Request) {

}

// MIDDLEWARES!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
func handleAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		jwtToken := r.Header.Get("Authorization")

		body, err := io.ReadAll(r.Body)

		var user db.User
		err = json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// Replace the body with a new reader after reading from the original
		r.Body = io.NopCloser(bytes.NewBuffer(body))

		err = verifyToken(jwtToken, user.Username)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

// func reuseBody(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		r.Body = io.NopCloser(util.ReusableReader(r.Body))
// 		next.ServeHTTP(w, r)
// 	})
// }
