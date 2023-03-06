package main

import "fmt"

func minMaxDifference(num int) int {
	nums := make([]int, 0)
	if num == 0 {
		nums = append(nums, 0)
	} else {
		for i := num; i > 0; i = i / 10 {
			nums = append(nums, i%10)
		}
	}
	n := len(nums)

	numToReplace := -1
	for i := n - 1; i >= 0; i-- {
		if nums[i] == 9 {
			continue
		}
		numToReplace = nums[i]
		break
	}
	maxValue := 0
	tens := 1
	for i := 0; i < n; i++ {
		x := nums[i]
		if x == numToReplace {
			x = 9
		}
		maxValue += x * tens
		tens *= 10
	}

	numToReplace = nums[n-1]
	minValue := 0
	tens = 1
	for i := 0; i < n; i++ {
		x := nums[i]
		if x == numToReplace {
			x = 0
		}
		minValue += x * tens
		tens *= 10
	}

	return maxValue - minValue
}

func main() {
	inputs := []int{11891, 90, 5}
	for _, num := range inputs {
		fmt.Println(num, minMaxDifference(num))
	}
}
