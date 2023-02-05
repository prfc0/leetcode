package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	origMp := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		origMp[s1[i]-'a']++
	}
	for i := 0; i < len(s2)-len(s1); i++ {
		if origMp[s2[i]-'a'] == 0 {
			continue
		}

		mp := make([]int, 26)
		copy(mp, origMp)

		matched := true
		for j := i; j < i+len(s1) && j < len(s2); j++ {
			ch := s2[j] - 'a'
			if mp[ch] == 0 {
				matched = false
				break
			}
			mp[ch]--
		}
		if matched {
			return true
		}
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
	}
	for _, input := range inputs {
		fmt.Println(checkInclusion(input.s1, input.s2))
	}
}
