package main

import "fmt"

/*
	Textio uses basic authentication to log users in.

	The code on the right has a type error. Change the type of password to a string (but use the same numeric value) so that it can be concatenated with the username variable.
*/

func main() {
	var username string = "presidentSkroob"
	var password string = "12345"

	fmt.Println("Authorization: Basic", username+":"+password)
}