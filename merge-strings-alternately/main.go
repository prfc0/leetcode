package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	newWord := make([]byte, len(word1)+len(word2))
	p1, p2 := 0, 0
	for p1 < len(word1) && p1 < len(word2) {
		newWord[p2] = word1[p1]
		p2++
		newWord[p2] = word2[p1]
		p1++
		p2++
	}
	for i := p1; i < len(word1); i++ {
		newWord[p2] = word1[i]
		p2++
	}
	for i := p1; i < len(word2); i++ {
		newWord[p2] = word2[i]
		p2++
	}

	return string(newWord)
}

func main() {
	// word1 := "abc"
	// word2 := "pqr"
	// word1 := "ab"
	// word2 := "pqrs"
	word1 := "abcd"
	word2 := "pq"
	fmt.Println(mergeAlternately(word1, word2))
}
