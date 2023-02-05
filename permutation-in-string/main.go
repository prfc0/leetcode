package main

import "fmt"

func check(s1, s2 []int) bool {
	for i := 0; i < 26; i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func checkInclusion(s1 string, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	if n1 > n2 {
		return false
	}

	s1Map := make([]int, 26)
	s2Map := make([]int, 26)
	for i := 0; i < n1; i++ {
		s1Map[s1[i]-'a']++
	}
	for i := 0; i < n1-1; i++ {
		s2Map[s2[i]-'a']++
	}
	for i := n1 - 1; i < n2; i++ {
		s2Map[s2[i]-'a']++
		if check(s1Map, s2Map) {
			return true
		}
		s2Map[s2[i-n1+1]-'a']--
	}
	return false
}

type Input struct {
	s1, s2 string
}

func main() {
	inputs := []Input{
		{
			s1: "ab",
			s2: "eidbaooo",
		},
		{
			s1: "ab",
			s2: "eidboaoo",
		},
		{
			s1: "hello",
			s2: "ooolleoooleh",
		},
		{
			s1: "adc",
			s2: "dcda",
		},
	}
	for _, input := range inputs {
		fmt.Println(checkInclusion(input.s1, input.s2))
	}
}
