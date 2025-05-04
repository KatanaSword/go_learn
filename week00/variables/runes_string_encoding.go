package main

import (
	"fmt"
	"unicode/utf8"
)

/*
	Boots is a bear. (Not a dog, haters).

	1. Run the code as-is. Notice that the simple string "boots" has 5 bytes, and 5 runes (characters).
	2. Update the name variable to be the bear emoji instead of the word "boots".

	🐻

	If you've got it right, you should notice that the emoji is only one rune, but it takes up 4 bytes.
*/

func main() {
	const name = "🐻"
	fmt.Printf("constant 'name' byte length: %d\n", len(name))
	fmt.Printf("constant 'name' rune length: %d\n", utf8.RuneCountInString(name))
	fmt.Println("=====================================")
	fmt.Printf("Hi %s, so good to have you back in the arcanum\n", name)
}
