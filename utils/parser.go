package utils

import (
	"errors"
	"log"
	"strconv"

	db "LeafMS-BackEnd/database"
)

func ParseStringToDate(date string) (db.Datetime, error) {
	const (
		Date = iota
		Month
	)

	var valToInsert = Date
	var parsedDate db.Datetime
	var currParserVal = ""
	var index = 0

	for index < len(date) {
		if date[index] != '/' {
			currParserVal += string(date[index])
		} else {
			parsedInt, err := strconv.Atoi(currParserVal)
			if err != nil {
				log.Println("Encountered error while parsing the date. Error:	", err)
				return parsedDate, err
			}

			if valToInsert == Date {
				parsedDate.Day = parsedInt
				valToInsert = Month
			} else {
				parsedDate.Month = parsedInt
			}
			currParserVal = ""
		}
		index++
	}

	parsedInt, err := strconv.Atoi(currParserVal)
	if err != nil {
		log.Println("Encountered error while parsing the date. Error:	", err)
		return parsedDate, err
	}
	parsedDate.Year = parsedInt
	if parsedDate.Day > 31 || parsedDate.Month > 12 {
		errStr := "The date provided is more than what is practically possible."
		errStr += "For ex - the day of month is either 32nd day or the month is 13th, which is not possible."
		err = errors.New(errStr)
		return parsedDate, err
	}
	return parsedDate, nil
}

func ParseDateToString(date db.Datetime) string {
	dateInStr := ""
	dateInStr += strconv.Itoa(date.Day) + string('/')
	dateInStr += strconv.Itoa(date.Month) + string('/')
	dateInStr += strconv.Itoa(date.Year)
	return dateInStr
}
