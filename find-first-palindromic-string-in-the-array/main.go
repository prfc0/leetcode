package main

import "fmt"

func isPalindrome(word string) bool {
	p1, p2 := 0, len(word)-1
	for p1 < p2 {
		if word[p1] != word[p2] {
			return false
		}
		p1++
		p2--
	}
	return true
}

func firstPalindrome(words []string) string {
	for _, word := range words {
		if isPalindrome(word) {
			return word
		}
	}
	return ""
}

func main() {
	S := []string{"abc", "car", "ada", "racecar", "cool"}
	fmt.Println(firstPalindrome(S))
}
