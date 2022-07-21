package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	newWords := make([]string, len(words))
	for i := 0; i < len(words); i++ {
		word := words[i]
		revWord := make([]byte, len(word))
		p1, p2 := 0, len(word)-1
		for p1 < len(word) && p2 >= 0 {
			revWord[p1] = word[p2]
			p1++
			p2--
		}
		newWords[i] = string(revWord)
	}
	revS := strings.Join(newWords, " ")
	return revS
}

func main() {
	S := "Let's take LeetCode contest"
	fmt.Println(reverseWords(S))
}
