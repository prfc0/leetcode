package main

/*
https://leetcode.com/problems/excel-sheet-column-number/
Given a string columnTitle that represents the column title as appears in an Excel sheet, return its corresponding column number.
For example:
A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28
...
*/

import "fmt"

func titleToNumber(columnTitle string) int {
	N := len(columnTitle)
	digits := 1
	for i := 1; i < N; i++ {
		digits *= 26
	}
	ans := 0
	for i := 0; i < len(columnTitle); i++ {
		num := columnTitle[i] - 'A' + 1
		ans += (int(num) * digits)
		digits /= 26
	}
	return ans
}

func main() {
	S := "M"
	fmt.Println(titleToNumber(S))
}
