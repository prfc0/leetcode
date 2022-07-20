package main

/*
https://leetcode.com/problems/longest-common-prefix/
Write a function to find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string "".
*/

import "fmt"

func longestCommonPrefix(strs []string) string {
	prefix := []byte{}
	for i := 0; i < len(strs[0]); i++ {
		ch := strs[0][i]
		mismatch := false
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) < i+1 {
				mismatch = true
				break
			}
			if strs[j][i] != ch {
				mismatch = true
			}
		}
		if mismatch {
			break
		} else {
			prefix = append(prefix, ch)
		}
	}
	return string(prefix)
}

func main() {
	// strs := []string{"flower", "flow", "flight"}
	strs := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs))
}
