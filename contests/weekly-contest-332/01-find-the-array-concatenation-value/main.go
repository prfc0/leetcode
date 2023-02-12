package main

import "fmt"

func countDigits(num int) int {
	count := 0
	for ; num > 0; num /= 10 {
		count++
	}
	return count
}

func findTheArrayConcVal(nums []int) int64 {
	n := len(nums)
	var ans int64
	for i := 0; i < n/2; i++ {
		n1, n2 := nums[i], nums[n-1-i]
		numDigits := countDigits(n2)
		fmt.Println(n1, n2, numDigits)
		for j := 0; j < numDigits; j++ {
			n1 *= 10
		}
		ans += int64(n1 + n2)
	}
	if n%2 == 1 {
		ans += int64(nums[n/2])
	}
	return ans
}

func main() {
	inputs := [][]int{
		{7, 52, 2, 4},
		{5, 14, 13, 8, 12},
		{1},
	}
	for _, input := range inputs {
		fmt.Println(findTheArrayConcVal(input))
	}
}
