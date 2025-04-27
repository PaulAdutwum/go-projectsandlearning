package main

import (
	"fmt"
)

// reverseString reverses the characters of the input string using two pointers
func reverseString(s string) string {
	// Convert string to slice of runes (important for Unicode support!)
	runes := []rune(s)
	left := 0
	right := len(runes) - 1

	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}

func main() {
	fmt.Println("🔄 Welcome to the Reverse String Tool!")

	var input string
	fmt.Print("Enter a word or sentence to reverse: ")
	fmt.Scanln(&input)

	reversed := reverseString(input)

	fmt.Println("\n✨ Reversed Result:", reversed)
}
