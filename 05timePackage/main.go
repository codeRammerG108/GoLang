package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	formattedTime := start.Format("Mon, Jan 02 2006 15:04:05")
	fmt.Println("This time is : ", formattedTime)

	createdDate := time.Date(2020, time.August, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println("This time in 2020 is : ", createdDate.Format("Monday, January 02 2006"))
}
