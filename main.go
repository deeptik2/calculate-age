package main

import (
	"calculate-age/validation"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

func main() {
	fmt.Println("Please provide your birthday in ddmmyyyy format? e.g., 01121991")
	var dobDate string
	fmt.Scanln(&dobDate)
	log.Infof("your birthday is: %v", dobDate)

	// check date input
	day, month, year, err := validation.IsValidDate(dobDate)
	if err != nil {
		log.Errorf("Date of birthday invalid: %v", err)
	}

	// convert to date string
	dobDateInFormat, err := validation.ConvertStringToDate(day, month, year)
	if err != nil {
		log.Errorf("error converting birth date: %v", err)
	}

	fmt.Println(dobDateInFormat)

	// get future date
	fmt.Println("Enter a future dobDate in ddmmyyyy format to check your age: ")
	var futureDate string
	fmt.Scanln(&futureDate)

	// check future date input
	day, month, year, err = validation.IsValidDate(futureDate)
	if err != nil {
		log.Errorf("future date invalid: %v", err)
	}
	log.Infof("the future date is: %v", dobDate)

	// convert to date string
	futureDateInFormat, err := validation.ConvertStringToDate(day, month, year)
	if err != nil {
		log.Errorf("error converting future date: %v", err)
	}

	// get date difference
	age, err := validation.GetAge(dobDateInFormat, futureDateInFormat)
	if err != nil {
		log.Errorf("error calculating days: %v", err)
	} else {
		fmt.Printf("you will be %v days old \n", age)
	}
}
