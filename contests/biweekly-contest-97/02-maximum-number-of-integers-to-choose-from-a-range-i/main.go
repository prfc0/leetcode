package main

import "fmt"

func maxCount(banned []int, n int, maxSum int) int {
	bannedMap := make(map[int]bool)
	for _, num := range banned {
		bannedMap[num] = true
	}

	sum, count := 0, 0
	for i := 1; i <= n; i++ {
		if bannedMap[i] {
			continue
		}
		sum += i
		if sum > maxSum {
			return count
		}
		count++
	}
	return count
}

type Input struct {
	banned    []int
	n, maxSum int
}

func main() {
	inputs := []Input{
		{
			banned: []int{1, 6, 5},
			n:      5,
			maxSum: 6,
		},
		{
			banned: []int{1, 2, 3, 4, 5, 6, 7},
			n:      8,
			maxSum: 1,
		},
		{
			banned: []int{11},
			n:      7,
			maxSum: 50,
		},
	}
	for _, input := range inputs {
		fmt.Println(maxCount(input.banned, input.n, input.maxSum))
	}
}
