package main

import (
	"fmt"
	"time"
)

func main() {
	PresentTime := time.Now()
	fmt.Println(PresentTime)
	fmt.Println("The present time is: " + PresentTime.Format("15:04:05"))
	fmt.Println("The present day is: " + PresentTime.Format("Monday"))
	fmt.Println("The present year is: " + PresentTime.Format("01-02-2006"))
	createdDate := time.Date(2056, time.September, 17, 01, 34, 24, 546, time.UTC)
	fmt.Print("The created time is: ")
	fmt.Println(createdDate)
	fmt.Println("Formated created time is: " + createdDate.Format("01-02-2006 , Monday ,15:04:05"))
}
