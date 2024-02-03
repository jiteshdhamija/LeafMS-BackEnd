package main

type User struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	Team        string `json:"team"`
	Designation string `json:"designation"`
	Approver    string `json:"approver"`
}

type LeaveSpan struct {
	Start string `json:"startTime"`
	End   string `json:"endTime"`
}

type Leave struct {
	Username string      `json:"username"`
	Leaves   []LeaveSpan `json:"leaves"`
}
