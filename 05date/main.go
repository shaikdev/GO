package main

import (
	"fmt"
	"time"
)

func main() {
	// current Date
	currentDate := time.Now()
	fmt.Println("currentDate", currentDate)

	// formatted Date
	formattedDate := currentDate.Format("01-02-2006")
	fmt.Println("formattedDate", formattedDate)

	// formatted by year/month/date/day/time/hour
	formatDate := currentDate.Format("01-02-2006 15:04:05 Monday")
	fmt.Println("formatDate", formatDate)

	// formatted by year/month/date/day/time/hour/sec
	formatDateWithSec := currentDate.Format("01-02-2006 15:04:05 Monday")
	fmt.Println("formatDateWithSec", formatDateWithSec)
}
