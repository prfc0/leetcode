package main

import "fmt"

func makeStringsEqual(s string, target string) bool {
	xor := 0
	for i := 0; i < len(s); i++ {
		sChar := s[i]
		sInt := int(sChar) - '0'
		tChar := target[i]
		tInt := int(tChar) - '0'
		fmt.Println(sInt, tInt)
		xor += sInt ^ tInt
	}
	if xor > 1 {
		return false
	}
	return true
}

type Input struct {
	str, target string
}

func main() {
	inputs := []Input{
		{
			str:    "1010",
			target: "0110",
		},
		{
			str:    "11",
			target: "00",
		},
	}
	for _, input := range inputs {
		fmt.Println(makeStringsEqual(input.str, input.target))
	}
}
