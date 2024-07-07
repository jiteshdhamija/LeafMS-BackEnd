package utils

import (
	db "LeafMS-BackEnd/database"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var daysInAMonth = map[int]int{
	1:  31,
	2:  28,
	3:  31,
	4:  30,
	5:  31,
	6:  30,
	7:  31,
	8:  31,
	9:  30,
	10: 31,
	11: 30,
	12: 31,
}

func adjustForLeapYear(year int) {

}

func rollBackLeaveOneDay(date db.Datetime) db.Datetime {
	if date.Day == 1 {
		date.Day = daysInAMonth[date.Month-1]
		if date.Month == 1 {
			date.Year -= 1
			date.Month = 12
		} else {
			date.Month -= 1
		}
	} else {
		date.Day -= 1
	}
	return date
}
func rollForwardLeaveOneDay(date db.Datetime) db.Datetime {
	if date.Day == daysInAMonth[date.Month] {
		date.Day = 1
		if date.Month == 12 {
			date.Year += 1
			date.Month = 1
		} else {
			date.Month += 1
		}
	} else {
		date.Day += 1
	}
	return date
}

func RemoveHolidayFromLeaveData(leave db.LeaveData) ([]db.LeaveData, error) {
	var splitLeaves []db.LeaveData
	leaveStartDate, err := ParseStringToDate(leave.Start)
	if err != nil {
		log.Println("There was problem parsing the starting date of a leave Err:", err)
		return splitLeaves, err
	}
	leaveEndDate, err := ParseStringToDate(leave.End)
	if err != nil {
		log.Println("There was problem parsing the ending date of a leave Err:", err)
		return splitLeaves, err
	}

	holidaysRaw, err := database.Find("publicHolidays", bson.D{
		{Key: "$and", Value: bson.M{
			"date.year": bson.M{
				"$gte": leaveStartDate.Year,
				"$lte": leaveEndDate.Year,
			},
			"date.month": bson.M{
				"$gte": leaveStartDate.Month,
				"$lte": leaveEndDate.Month,
			},
			"date.day": bson.M{
				"$gte": leaveStartDate.Day,
				"$lte": leaveEndDate.Day,
			},
		}},
	})
	if err != nil {
		errMessage := "For fuck's sake, there was a problem, "
		errMessage += "while trying to find a holiday conflicting with the applied leave in the database. Err:"
		log.Println(errMessage, err)
		return splitLeaves, err
	}
	holidays := ReturnHolidays(holidaysRaw)

	//	this part needs to be modified to include the case
	//	when there is no holiday conflict
	startDate := leave.Start
	for _, holiday := range holidays {
		var leaveSpan db.LeaveData
		leaveSpan.Id = primitive.NewObjectID()
		leaveSpan.Start = startDate
		leaveSpan.End = ParseDateToString(rollBackLeaveOneDay(holiday.Date.Datetime))
		splitLeaves = append(splitLeaves, leaveSpan)
	}
	return splitLeaves, nil
}
