package main

import "fmt"

func main() {
	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondsInMinute*minutesInHour

	fmt.Println("Number of seconds in an hour:", secondsInHour)
}