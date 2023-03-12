package main

import "fmt"

func beautifulSubarrays(nums []int) int64 {
	n := len(nums)

	count := 0

	prefix := make([]int, n)
	prefix[0] = nums[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] ^ nums[i]
	}

	mp := make(map[int]int)

	for i := 0; i < n; i++ {
		count += mp[prefix[i]]
		if prefix[i] == 0 {
			count++
		}
		mp[prefix[i]]++
	}

	return int64(count)
}

func beautifulSubarraysBF(nums []int) int64 {
	n := len(nums)
	var count int64
	for i := 0; i < n; i++ {
		xor := nums[i]
		if xor == 0 {
			count++
		}
		for j := i + 1; j < n; j++ {
			xor ^= nums[j]
			if xor == 0 {
				count++
			}
			// fmt.Print(nums[j], " ")
		}
		// fmt.Println()
	}
	return count
}

func main() {
	inputs := [][]int{
		{4, 3, 1, 2, 4},
		{1, 10, 4},
		{0},
		{0, 0},
	}
	for _, nums := range inputs {
		fmt.Println(beautifulSubarrays(nums), beautifulSubarraysBF(nums))
	}
}
