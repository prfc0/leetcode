package main

import "fmt"

func isVowelString(word string) bool {
	first, last := word[0], word[len(word)-1]
	if (first == 'a' || first == 'e' || first == 'i' || first == 'o' || first == 'u') &&
		(last == 'a' || last == 'e' || last == 'i' || last == 'o' || last == 'u') {
		return true
	}
	return false
}

func vowelStrings(words []string, left int, right int) int {
	count := 0
	for i := left; i <= right; i++ {
		word := words[i]
		if isVowelString(word) {
			count++
		}
	}
	return count
}

type Input struct {
	words       []string
	left, right int
}

func main() {
	inputs := []Input{
		{
			words: []string{"are", "amy", "u"},
			left:  0,
			right: 2,
		},
		{
			words: []string{"hey", "aeo", "mu", "ooo", "artro"},
			left:  1,
			right: 4,
		},
	}
	for _, input := range inputs {
		fmt.Println(vowelStrings(input.words, input.left, input.right))
	}
}
