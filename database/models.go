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
	Approved bool               `default:"false" bson:"approved" json:"approved"`
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

//structs for fetching public holidays
type Meta struct {
	Code int `bson:"code" json:"code"`
}
type Country struct {
	Id   string `bson:"id" json:"id"`
	Name string `bson:"name" json:"name"`
}
type Datetime struct {
	Year  int `bson:"year" json:"year"`
	Month int `bson:"month" json:"month"`
	Day   int `bson:"day" json:"day"`
}
type Date struct {
	Iso      string   `bson:"iso" json:"iso"`
	Datetime Datetime `bson:"datetime" json:"datetime"`
}
type Holiday struct {
	Name         string   `bson:"name" json:"name"`
	Description  string   `bson:"description" json:"description"`
	Country      Country  `bson:"country" json:"country"`
	Date         Date     `bson:"date" json:"date"`
	Type         []string `bson:"type" json:"type"`
	PrimaryType  string   `bson:"primary_type" json:"primary_type"`
	CanonicalUrl string   `bson:"canonical_url" json:"canonical_url"`
	UrlId        string   `bson:"urlid" json:"urlid"`
	Locations    string   `bson:"locations" json:"locations"`
	States       string   `bson:"states" json:"states"`
}
type HolidayResponse struct {
	Holidays []Holiday `bson:"holidays" json:"holidays"`
}
type HolidayApiResponse struct {
	Meta     Meta            `bson:"meta" json:"meta"`
	Response HolidayResponse `bson:"response" json:"response"`
}

type HolidayArgs struct {
	Country string `bson:"country" json :"country"`
	Year    int    `bson:"year" json :"year"`
	Month   int    `bson:"month" json:"month"`
}
