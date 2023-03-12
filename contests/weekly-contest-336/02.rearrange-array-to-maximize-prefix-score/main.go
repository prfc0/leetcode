package main

import (
	"fmt"
	"sort"
)

func maxScore(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	fmt.Println(nums)
	sum, count := 0, 0
	for _, num := range nums {
		sum += num
		if sum > 0 {
			count++
		} else {
			break
		}
	}
	return count
}

func main() {
	inputs := [][]int{
		{2, -1, 0, 1, -3, 3, -3},
		{-2, -3, 0},
	}
	for _, nums := range inputs {
		fmt.Println(maxScore(nums))
	}
}
