package main

import "fmt"

/*
	At the top of the main function, declare a float called averageOpenRate and string called displayMessage on the same line.

	Initialize them to values:

	1. .23
	2. is the average open rate of your messages

	before they're printed.
*/

func main() {
	averageOpenRate, displayMessage := 0.23, "is the average open rate of your messages"
	fmt.Println(averageOpenRate, displayMessage)
}