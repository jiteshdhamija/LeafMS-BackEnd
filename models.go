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
	Start    string `json:"startTime"`
	End      string `json:"endTime"`
	Approval bool   `json:"approval"`
}

type Leave struct {
	Username string      `json:"username"`
	Approver string      `json:"approver"`
	Leaves   []LeaveSpan `json:"leaves"`
}

const (
	JuniorDev int = iota
	Dev
	SeniorDev
	TeamLead
	Manager
)
