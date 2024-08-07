package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	length := len(s)
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var input string
	fmt.Println("Enter a word or phrase:")
	fmt.Scanln(&input)

	if isPalindrome(input) {
		fmt.Println(input, "is a palindrome.")

	} else {
		fmt.Println(input, "is not a palindrome.")
	}
}
