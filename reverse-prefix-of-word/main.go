package main

import "fmt"

func reversePrefix(word string, ch byte) string {
	idx := 0
	for i := 0; i < len(word); i++ {
		if word[i] == ch {
			idx = i
			break
		}
	}
	newWord := make([]byte, len(word))
	for i := 0; i <= idx; i++ {
		newWord[i] = word[idx-i]
	}
	for i := idx + 1; i < len(word); i++ {
		newWord[i] = word[i]
	}
	return string(newWord)
}

func main() {
	word := "abcdefd"
	ch := byte('d')
	fmt.Println(reversePrefix(word, ch))
}
