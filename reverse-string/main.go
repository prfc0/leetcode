package main

import "fmt"

func reverseString(s []byte) {
	p1, p2 := 0, len(s)-1
	for p1 < p2 {
		s[p1], s[p2] = s[p2], s[p1]
		p1++
		p2--
	}
}

func main() {
	S := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(S)
	fmt.Println(string(S))
}
