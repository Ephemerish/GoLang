package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Println(currentTime.Format(("Monday 01-02-2006 15:04:05")))

	createdDate := time.Date(2024, time.July, 24, 23, 0, 0, 0, time.Local)
	//	fmt.Println(createdDate)
	fmt.Println(createdDate.Format(("01-02-2006 Monday")))
}
