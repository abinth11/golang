package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current time= ", currentTime)

	formattedTime := currentTime.Format("01-02-2006 15:04:05 Monday")
	fmt.Println("Formatted time= ", formattedTime)

	randomDate := time.Date(2020, time.November, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println("Current Date= ", randomDate)

	formattedDate := randomDate.Format("01-02-2006")
	fmt.Println("Formatted Date= ", formattedDate)
}
