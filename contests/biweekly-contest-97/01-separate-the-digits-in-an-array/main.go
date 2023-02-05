package main

import "fmt"

func separateDigits(nums []int) []int {
	ans := make([]int, 0)
	for _, num := range nums {
		tmp := make([]int, 0)
		for ; num > 0; num /= 10 {
			rem := num % 10
			tmp = append(tmp, rem)
		}
		for i := len(tmp) - 1; i >= 0; i-- {
			ans = append(ans, tmp[i])
		}
	}
	return ans
}

func main() {
	inputs := [][]int{
		{13, 25, 83, 77},
		{7, 1, 3, 9},
		{10921},
	}
	for _, input := range inputs {
		fmt.Println(separateDigits(input))
	}
}
