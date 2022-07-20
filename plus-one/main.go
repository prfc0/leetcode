package main

/*
https://leetcode.com/problems/plus-one/
You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.
Increment the large integer by one and return the resulting array of digits.
*/

import "fmt"

func plusOne(digits []int) []int {
	N := len(digits)
	carry := 1
	for i := N - 1; i >= 0; i-- {
		if carry == 1 {
			if digits[i] == 9 {
				digits[i] = 0
				carry = 1
			} else {
				digits[i]++
				carry = 0
			}
		}
	}
	if carry == 1 {
		tmp := []int{1}
		tmp = append(tmp, digits...)
		return tmp
	} else {
		return digits
	}
}

func main() {
	digits := []int{9, 9, 9}
	fmt.Println(plusOne(digits))
}
