package main

import "fmt"

func subsetXORSum(nums []int) int {
	N := len(nums)
	sum := 0
	for i := 0; i < (1 << N); i++ {
		xor := 0
		for j := 0; j < N; j++ {
			if (i>>j)&1 == 1 {
				xor ^= nums[j]
			}
		}
		sum += xor
	}
	return sum
}

func main() {
	// nums := []int{1, 3}
	nums := []int{3, 4, 5, 6, 7, 8}
	fmt.Println(subsetXORSum(nums))
}
