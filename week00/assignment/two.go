package main

import "fmt"

// The code that prints server logs won't even compile!

func main() {
	var startup string = "Textio SMS service booting up..."
	var message string = "Sending test message"
	var confirmation string = "Message sent!"

	fmt.Println(startup)
	fmt.Println(message)
	fmt.Println(confirmation)
}