package main

import "fmt"

func isVowel(char byte) bool {
	if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
		return true
	}
	return false
}

func isWordStartingAndEndingWithVowel(word string) bool {
	if isVowel(word[0]) && isVowel(word[len(word)-1]) {
		return true
	}
	return false
}

func vowelStrings(words []string, queries [][]int) []int {
	n := len(words)
	prefixSum := make([]int, n)

	for i, word := range words {
		if isWordStartingAndEndingWithVowel(word) {
			prefixSum[i] = 1
		}
	}

	for i := 1; i < len(prefixSum); i++ {
		prefixSum[i] += prefixSum[i-1]
	}

	fmt.Println("PrefixSum:", prefixSum)
	ans := make([]int, len(queries))
	for i, query := range queries {
		l, r := query[0], query[1]
		fmt.Println("query:", l, r, prefixSum[r], prefixSum[l])
		if l == 0 {
			ans[i] = prefixSum[r]
		} else {
			ans[i] = prefixSum[r] - prefixSum[l-1]
		}
	}

	return ans
}

type Input struct {
	words   []string
	queries [][]int
}

func main() {
	inputs := []Input{
		{
			words:   []string{"aba", "bcb", "ece", "aa", "e"},
			queries: [][]int{{0, 2}, {1, 4}, {1, 1}},
		},
		{
			words:   []string{"a", "e", "i"},
			queries: [][]int{{0, 2}, {0, 1}, {2, 2}},
		},
	}
	for _, input := range inputs {
		fmt.Println(vowelStrings(input.words, input.queries))
	}
}
