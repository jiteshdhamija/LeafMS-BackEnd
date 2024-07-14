package utils

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"

	db "LeafMS-BackEnd/database"
)

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

func ReturnHolidays(data []bson.Raw) []db.Holiday {
	var holidays []db.Holiday
	for _, entry := range data {
		var holiday db.Holiday
		if err := bson.Unmarshal(entry, &holiday); err != nil {
			log.Fatal(
				"The decoding of employee from raw bson document failed!\nError:-\n\n", err)
		}
		holidays = append(holidays, holiday)
	}
	return holidays
}
