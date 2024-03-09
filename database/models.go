package db

import (
	"fmt"
)

type collection interface {
	find() interface{}
}

func (u User) find() {
	fmt.Println(u.Username)
}
func (l Leave) find() {
	fmt.Println(l.Leaves)
}

type User struct {
	Username    string `bson:"username" json:"username"`
	Password    string `bson:"password" json:"password"`
	Name        string `bson:"name" json:"name"`
	Team        string `bson:"team" json:"team"`
	Designation string `bson:"designation" json:"designation"`
	Approver    string `bson:"approver" json:"approver"`
}

type UserLogin struct {
	Username string `bson:"username" json:"username"`
	Login    bool   `bson:"isLogin" json:"isLogin"`
}

type LeaveSpan struct {
	Start string `bson:"startTime" json:"startTime"`
	End   string `bson:"endTime" json:"endTime"`
}

type Leave struct {
	Username string      `bson:"username" json:"username"`
	Leaves   []LeaveSpan `bson:"leaves" json:"leaves"`
}
