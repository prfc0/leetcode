package main

import (
	"fmt"
	"sort"
)

func splitNum(num int) int {
	nums := make([]int, 0)
	for ; num > 0; num /= 10 {
		nums = append(nums, num%10)
	}
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
	num1, num2 := 0, 0
	for i, tens := 0, 1; i < len(nums); i, tens = i+2, tens*10 {
		num1 += nums[i] * tens
	}
	for i, tens := 1, 1; i < len(nums); i, tens = i+2, tens*10 {
		num2 += nums[i] * tens
	}
	return num1 + num2
}

func main() {
	inputs := []int{4325, 687}
	for _, num := range inputs {
		fmt.Println(num, splitNum(num))
	}
}
