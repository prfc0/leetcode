package main

import (
	"fmt"
	"sort"
)

func floorBS(nums []int, lo, hi, k int) int {
	ans := -1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if k >= nums[mid] {
			ans = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return ans
}

func ceilBS(nums []int, lo, hi, k int) int {
	ans := len(nums)
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if k <= nums[mid] {
			ans = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return ans
}

func countFairPairs(nums []int, lower int, upper int) int64 {
	n := len(nums)
	var ans int64
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		leftIndex := ceilBS(nums, i+1, n-1, lower-nums[i])
		rightIndex := floorBS(nums, i+1, n-1, upper-nums[i])
		if rightIndex >= leftIndex {
			ans += int64(rightIndex - leftIndex + 1)
		}
	}
	return ans
}

type Input struct {
	nums         []int
	lower, upper int
}

func main() {
	inputs := []Input{
		{
			nums:  []int{0, 1, 7, 4, 4, 5},
			lower: 3,
			upper: 6,
		},
		{
			nums:  []int{1, 7, 9, 2, 5},
			lower: 11,
			upper: 11,
		},
		{
			nums:  []int{0, 0, 0, 0, 0, 0},
			lower: 0,
			upper: 0,
		},
	}
	for _, input := range inputs {
		fmt.Println("Nums:", input.nums, "| Lower:", input.lower, "| Upper:", input.upper, "| Answer:", countFairPairs(input.nums, input.lower, input.upper))
	}
}
