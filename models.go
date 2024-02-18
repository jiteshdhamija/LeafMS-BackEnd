package main

type User struct {
	Username    string `bson:"username"`
	Password    string `bson:"password"`
	Name        string `bson:"name"`
	Team        string `bson:"team"`
	Designation string `bson:"designation"`
	Approver    string `bson:"approver"`
}

type LeaveSpan struct {
	Start string `json:"startTime" bson:"startTime"`
	End   string `json:"endTime" bson:"endTime"`
}

type Leave struct {
	Username string      `json:"username"`
	Leaves   []LeaveSpan `json:"leaves"`
}
