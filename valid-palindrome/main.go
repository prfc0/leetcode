package main

/*
https://leetcode.com/problems/valid-palindrome/
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.
Given a string s, return true if it is a palindrome, or false otherwise.
Example 1:
Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
*/

import "fmt"

func isPalindrome(s string) bool {
	str := []byte{}
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= 'A' && ch <= 'Z' {
			str = append(str, ch+32)
		} else if (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') {
			str = append(str, ch)
		}
	}
	p1, p2 := 0, len(str)-1
	for p1 < p2 {
		if str[p1] != str[p2] {
			return false
		}
		p1++
		p2--
	}
	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}
