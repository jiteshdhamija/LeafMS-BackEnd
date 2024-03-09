package utils

import (
	db "LeafMS-BackEnd/database"
	"fmt"
	"strings"
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
		break
	case "password":
		user.Password = temp2[1]
		break
	case "name":
		user.Name = temp2[1]
		break
	case "team":
		user.Team = temp2[1]
		break
	case "designation":
		user.Designation = temp2[1]
		break
	case "approver":
		user.Approver = temp2[1]
		break
	}
	return user
}
