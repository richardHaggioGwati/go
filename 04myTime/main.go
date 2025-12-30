package main

import (
	"time"
	"fmt"
)

func main() {
	fmt.Println("Time Management")

	presentTime := time.Now()
	
	fmt.Println("Present time is:", presentTime.String())
	fmt.Println("Year is:", presentTime.Year())
	fmt.Println("Month is:", presentTime.Month().String())
	fmt.Println("Day is:", presentTime.Day())
	fmt.Println("Hour is:", presentTime.Hour())
	fmt.Println("Minute is:", presentTime.Minute())
	fmt.Println("Second is:", presentTime.Second())
	fmt.Println("Nanosecond is:", presentTime.Nanosecond())
	fmt.Println("Location is:", presentTime.Location().String())
	fmt.Println("Weekday is:", presentTime.Weekday().String())
}