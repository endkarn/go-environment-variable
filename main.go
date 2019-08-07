package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	currentTime := getCurrentTime()
	fmt.Printf("Current time is : %s", currentTime.String())
	fmt.Printf("\nMy user is %s", os.Getenv("username"))
}

func getCurrentTime() time.Time {
	if os.Getenv("MYTIME") != "" {
		myTime := os.Getenv("MYTIME")
		parsedMyTime := convertDateStringToTime(myTime)
		return parsedMyTime
	}
	return time.Now().In(time.UTC)
}

func convertDateStringToTime(dateFormatted string) time.Time {
	dateLayout := "20060102"
	timeNow, err := time.Parse(dateLayout, dateFormatted)
	if err != nil {
		log.Printf("Error pasing date from env , %s", err)
	}
	return timeNow
}
