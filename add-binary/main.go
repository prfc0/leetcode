package main

/*
https://leetcode.com/problems/add-binary/
Given two binary strings a and b, return their sum as a binary string.
Example 1:
Input: a = "11", b = "1"
Output: "100"
*/

import "fmt"

func addBinary(a string, b string) string {
	p1, p2 := len(a)-1, len(b)-1
	ans := ""
	carry := false
	for p1 >= 0 && p2 >= 0 {
		if a[p1] == '1' && b[p2] == '1' {
			if carry {
				ans = "1" + ans
			} else {
				ans = "0" + ans
				carry = true
			}
		} else if (a[p1] == '1' && b[p2] == '0') || (a[p1] == '0' && b[p2] == '1') {
			if carry {
				ans = "0" + ans
			} else {
				ans = "1" + ans
				carry = false
			}
		} else {
			if carry {
				ans = "1" + ans
				carry = false
			} else {
				ans = "0" + ans
			}
		}
		p1--
		p2--
	}
	for p1 >= 0 {
		if a[p1] == '1' {
			if carry {
				ans = "0" + ans
			} else {
				ans = "1" + ans
			}
		} else {
			if carry {
				ans = "1" + ans
				carry = false
			} else {
				ans = "0" + ans
			}
		}
		p1--
	}
	for p2 >= 0 {
		if b[p2] == '1' {
			if carry {
				ans = "0" + ans
			} else {
				ans = "1" + ans
			}
		} else {
			if carry {
				ans = "1" + ans
				carry = false
			} else {
				ans = "0" + ans
			}
		}
		p2--
	}
	if carry {
		ans = "1" + ans
	}
	return ans
}

func main() {
	// a := "1010"
	// b := "1011"
	a, b := "11", "1"
	fmt.Println(addBinary(a, b))
}
