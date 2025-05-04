package main

import "fmt"

/*
	Our current price to send a text message is 2 cents. However, it's likely that in the future the price will be a fraction of a penny, so we will need to use a float64 to store it.

	Edit the penniesPerText declaration so that it's inferred by the compiler as a float64.
*/

func main() {
	penniesPerText := 2.00

	fmt.Printf("The type of penniesPerText is %T\n", penniesPerText)
}