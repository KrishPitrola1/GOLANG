package main

import (
	"fmt"
	"strings"
)

func capitalize(input string) string {
	return strings.ToUpper(input)
}

func main() {
	// Prompt the user for input
	fmt.Print("Enter a string: ")

	// Read input from the user
	var input string
	fmt.Scanln(&input)

	// Capitalize the input string
	capitalized := capitalize(input)

	// Display the capitalized string
	fmt.Println("Capitalized string:", capitalized)
}
