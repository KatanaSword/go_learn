package main

import "fmt"

// Initialize the variables from the print statement to int, float64, bool and string with their zero values, respectively.

func main() {
	// initialize variables here
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}