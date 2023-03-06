package main

import (
	"fmt"
	"math"
)

func minimizeSum(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	// Initialize variables to hold minimum and maximum values.
	min1, min2 := math.MaxInt64, math.MaxInt64
	max1, max2 := math.MinInt64, math.MinInt64

	// Find the two smallest and two largest numbers in the array.
	for _, num := range nums {
		if num <= min1 {
			min1, min2 = num, min1
		} else if num < min2 {
			min2 = num
		}
		if num >= max1 {
			max1, max2 = num, max1
		} else if num > max2 {
			max2 = num
		}
	}

	// Calculate the score with the original array.
	lowScore := abs(min1 - min2)
	highScore := abs(max1 - max2)
	originalScore := lowScore + highScore

	// Try changing each number to the minimum or maximum value.
	newLowScore := math.MaxInt64
	newHighScore := math.MinInt64

	for _, num := range []int{min1, min2, max1, max2} {
		// Calculate the new low and high scores after changing the number.
		if num != min1 {
			newLowScore = min(newLowScore, abs(min2-num))
		}
		if num != min2 {
			newLowScore = min(newLowScore, abs(min1-num))
		}
		if num != max1 {
			newHighScore = max(newHighScore, abs(max2-num))
		}
		if num != max2 {
			newHighScore = max(newHighScore, abs(max1-num))
		}
	}

	// Check if changing two numbers can result in a lower score.
	if n >= 3 {
		for i := 0; i < n-2; i++ {
			for j := i + 1; j < n-1; j++ {
				for k := j + 1; k < n; k++ {
					// Calculate the new low and high scores after changing two numbers.
					newMin1, newMin2, newMax1, newMax2 := nums[i], nums[j], nums[i], nums[j]
					if nums[k] < newMin1 {
						newMin1, newMin2 = nums[k], newMin1
					} else if nums[k] < newMin2 {
						newMin2 = nums[k]
					}
					if nums[k] > newMax1 {
						newMax1, newMax2 = nums[k], newMax1
					} else if nums[k] > newMax2 {
						newMax2 = nums[k]
					}
					newLowScore = min(newLowScore, abs(newMin1-newMin2))
					newHighScore = max(newHighScore, abs(newMax1-newMax2))
				}
			}
		}
	}

	// Calculate the new score after changing one or two numbers.
	newScore := newLowScore + newHighScore

	// Return the minimum score between the original and new scores.
	return min(originalScore, newScore)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
25 31 65 72 74 79 : 54 48 49 43
    6 34  7  2  5

1 4 5 7 8 : 3, 6, 4, 7
  3 1 2 1

*/
func main() {
	inputs := [][]int{
		{1, 4, 3},
		{1, 4, 7, 8, 5},
		{31, 25, 72, 79, 74, 65},
	}
	for _, nums := range inputs {
		fmt.Println(nums, minimizeSum(nums))
	}
}
