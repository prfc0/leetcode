package main

/*
https://leetcode.com/problems/palindrome-number/
Given an integer x, return true if x is palindrome integer.
An integer is a palindrome when it reads the same backward as forward.
For example, 121 is a palindrome while 123 is not.
*/

import "fmt"

func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	numDigits := -1
	for i := x; i > 0; i = i / 10 {
		numDigits++
	}
	digits := 1
	for j := 1; j <= numDigits; j++ {
		digits *= 10
	}
	reverseNum := 0
	for i := x; i > 0; i = i / 10 {
		rem := i % 10
		reverseNum = rem*digits + reverseNum
		digits = digits / 10
	}
	fmt.Println(reverseNum)
	if x == reverseNum {
		return true
	}
	return false
}

func main() {
	fmt.Println(isPalindrome(121))
}
