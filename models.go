package main

type User struct {
	Username    string `json:"username" bson:"username"`
	Password    string `json:"password" bson:"password"`
	Name        string `json:"name" bson:"name"`
	Team        string `json:"team" bson:"team"`
	Designation string `json:"designation" bson:"designation"`
	Approver    string `json:"approver" bson:"approver"`
}

type LeaveSpan struct {
	Start string `json:"startTime" bson:"startTime"`
	End   string `json:"endTime" bson:"endTime"`
}

type Leave struct {
	Username string      `json:"username"`
	Leaves   []LeaveSpan `json:"leaves"`
}
