package main

import "fmt"

/*
1. Complete the main function. It should print: "Happy birthday! You are now 21 years old!".
	- Create a string variable messageStart with the text "Happy birthday! You are now"
	- Create an integer variable age set to 21
	- Create another string variable messageEnd with the text "years old!"

2. The provided fmt.Println statement will print the full message on a single line separated by spaces.
*/

func main() {
	messageStart:= "Happy birthday!"
	age:= 27
	messageEnd:= "years old"

	fmt.Println(messageStart, age, messageEnd)
}