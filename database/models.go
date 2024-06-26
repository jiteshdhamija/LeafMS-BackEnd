package db

import "go.mongodb.org/mongo-driver/bson/primitive"

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

type LeaveData struct {
	Id       primitive.ObjectID `bson:"id" json:"id"`
	Start    string             `bson:"startTime" json:"startTime"`
	End      string             `bson:"endTime" json:"endTime"`
	Approved bool               `bson:"approved" json:"approved"`
}

type Leaves struct {
	Username string      `bson:"username" json:"username"`
	Approver string      `bson:"approver" json:"approver"`
	Leaves   []LeaveData `bson:"leaves" json:"leaves"`
}

type LeavesCount struct {
	CausalLeaves     int `bson:"casualLeaves" json:"casualLeaves"`
	MedicalLeaves    int `bson:"medicalLeaves" json:"medicalLeaves"`
	PrivilegedLeaves int `bson:"privilegedLeaves" json:"privilegedLeaves"`
	CompOff          int `bson:"compOff" json:"compOff"`
	TotalLeaveCount  int `bson:"totalLeaveCount" json:"totalLeaveCount"`
}
