package main

import (
	"fmt"
	"math"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minCapability(nums []int, k int) int {
	ans := math.MaxInt
	for i := 0; i < len(nums)-k; i++ {
		// tmpMax := math.MinInt
		for j := i; j < i+k; j++ {
			ans = min(ans, max(nums[i], nums[j]))
			// tmpMax = max(tmpMax, max(nums[i], nums[j]))
		}
		// ans = min(ans, tmpMax)
	}
	return ans
}

type Input struct {
	nums []int
	k    int
}

func main() {
	inputs := []Input{
		{
			nums: []int{2, 3, 5, 9},
			k:    2,
		},
		{
			nums: []int{2, 7, 9, 3, 1},
			k:    2,
		},
		{
			nums: []int{1, 4, 5},
			k:    1,
		},
	}
	for _, input := range inputs {
		fmt.Println(minCapability(input.nums, input.k))
	}
}
