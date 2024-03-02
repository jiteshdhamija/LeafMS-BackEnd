package db

type User struct {
	Username    string `bson:"username" json:"username"`
	Password    string `bson:"password" json:"password"`
	Name        string `bson:"name" json:"name"`
	Team        string `bson:"team" json:"team"`
	Designation string `bson:"designation" json:"designation"`
	Approver    string `bson:"approver" json:"approver"`
}

type LeaveSpan struct {
	Start string `bson:"startTime" json:"startTime"`
	End   string `bson:"endTime" json:"endTime"`
}

type Leave struct {
	Username string      `bson:"username" json:"username"`
	Leaves   []LeaveSpan `bson:"leaves" json:"leaves"`
}
