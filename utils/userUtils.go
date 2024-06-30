package utils

import (
	db "LeafMS-BackEnd/database"
	"fmt"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func InterFaceToUser(user interface{}) db.User {
	var result db.User
	str := fmt.Sprintf("%v", user)
	for len(str) > 4 {
		i1 := strings.Index(str, "{")
		i2 := strings.Index(str, "}")
		temp := str[i1+1 : i2]
		str = str[i2+1:]
		temp2 := strings.Split(temp, " ")
		result = setUserVal(temp2, result)
	}
	return result
}

func setUserVal(temp2 []string, user db.User) db.User {
	switch temp2[0] {
	case "username":
		user.Username = temp2[1]
	case "password":
		user.Password = temp2[1]
	case "name":
		user.Name = temp2[1]
	case "team":
		user.Team = temp2[1]
	case "designation":
		user.Designation = temp2[1]
	case "approver":
		user.Approver = temp2[1]
	}
	return user
}

func ReturnLeaves(data []bson.Raw) []db.Leaves {
	var leaves []db.Leaves
	for _, entry := range data {
		var leave db.Leaves
		if err := bson.Unmarshal(entry, &leave); err != nil {
			log.Fatal(
				"The decoding of leaveApplication from raw bson document failed!\nError:-\n\n", err)
		}
		leaves = append(leaves, leave)
	}
	return leaves
}

func ReturnUsers(data []bson.Raw) []db.User {
	var employees []db.User
	for _, entry := range data {
		var employee db.User
		if err := bson.Unmarshal(entry, &employee); err != nil {
			log.Fatal(
				"The decoding of employee from raw bson document failed!\nError:-\n\n", err)
		}
		employees = append(employees, employee)
	}
	return employees
}
