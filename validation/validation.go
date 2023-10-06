package validation

import (
	"fmt"
	"time"
)

func ConvertStringToDate(day, month, year string) (time.Time, error) {

	dateWithTZ := year + "-" + month + "-" + day + "T00:00:01Z"
	dateFormat, err := time.Parse(time.RFC3339, dateWithTZ)
	if err != nil {
		return dateFormat, err
	}
	return dateFormat, nil
}

func IsValidDate(date string) (string, string, string, error) {
	var day, month, year string
	if len(date) == 8 {
		day = date[0:2]
		month = date[2:4]
		year = date[4:8]
	} else {
		return "", "", "", fmt.Errorf("date should be ddmmyyyy format")
	}
	return day, month, year, nil
}

func GetAge(dob, futureDate time.Time) (age float64, err error) {
	diff := futureDate.Sub(dob).Hours()
	var days float64

	if diff < 0 {
		return 0.0, fmt.Errorf("future date is older than Date of Birth")
	}
	if diff == 0 {
		return 0.0, fmt.Errorf("future date is same as Date of Birth")
	}

	if diff > 24 {
		days = diff / 24
	}
	return days, nil
}
