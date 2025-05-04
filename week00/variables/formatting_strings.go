package main

import "fmt"

/*
Create a new variable called msg on line 9 and use the appropriate formatting function to return a string that contains following:

Hi NAME, your open rate is OPENRATE percentNEWLINE

1. Replace NAME with the variable name,
2. Replace OPENRATE with the variable openRate rounded to the nearest "tenths" place, e.g 10.523 should be rounded down to 10.5
3. The word percent should appear as part of the string following the open rate value
4. Replace NEWLINE with the newline \n escape sequence.
*/

func main() {
	const name = "Saul Goodman"
	const openRate = 30.5

	msg := fmt.Sprintf("Hi %v, your open rate is %.1f percent\n", name, openRate)

	fmt.Print(msg)
}