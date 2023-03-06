package main

import (
	"fmt"
	"sort"
)

const (
	MOD = int(1e9) + 7
)

func calcPower(base, power int) int {
	ans := 1
	for power > 0 {
		if power%2 != 0 {
			ans = (ans * base) % MOD
			power--
		} else {
			power = power / 2
			base = (base * base) % MOD
		}
	}
	return ans
}

func countWays(ranges [][]int) int {
	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i][0] == ranges[j][0] {
			return ranges[i][1] < ranges[j][1]
		}
		return ranges[i][0] < ranges[j][0]
	})
	newRanges := make([][]int, 0)
	for i := 0; i < len(ranges); i++ {
		pair := ranges[i]
		x1, y1 := pair[0], pair[1]
		j := i + 1
		for j < len(ranges) {
			x2, y2 := ranges[j][0], ranges[j][1]
			if x2 <= y1 && x2 >= x1 {
				if y2 > y1 {
					y1 = y2
				}
				j++
			} else {
				break
			}
		}
		i = j - 1
		newRanges = append(newRanges, []int{x1, y1})
	}
	fmt.Println(newRanges)
	return calcPower(2, len(newRanges))
}

func main() {
	inputs := [][][]int{
		{{6, 10}, {5, 15}},
		{{1, 3}, {10, 20}, {2, 5}, {4, 8}},
		{{34, 56}, {28, 29}, {12, 16}, {11, 48}, {28, 54}, {22, 55}, {28, 41}, {41, 44}},
	}
	for _, nums := range inputs {
		fmt.Println(nums, countWays(nums))
	}
}
