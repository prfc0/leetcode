package main

import "fmt"

func maximizeWin(prizePositions []int, k int) int {
	n := len(prizePositions)
	max := prizePositions[n-1]
	if k >= max {
		return n
	}
	fmt.Println("Max:", max)
	prizeCount := make(map[int]int)
	for i := 0; i < n; i++ {
		prizeCount[prizePositions[i]]++
	}
	fmt.Println("prizeCount:", prizeCount)
	sum1 := 0
	for i := 1; i <= k; i++ {
		sum1 += prizeCount[i]
	}
	max1, maxIndex := 0, 0
	for i := k + 1; i <= max; i++ {
		sum1 += prizeCount[i]
		if sum1 > max1 {
			maxIndex = i
			max1 = sum1
		}
		sum1 -= prizeCount[i-k]
	}
	fmt.Println("sum1, max1, maxIndex:", sum1, max1, maxIndex)
	for i := maxIndex; i >= maxIndex-k; i-- {
		prizeCount[i] = 0
	}
	fmt.Println("prizeCount:", prizeCount)
	max2 := 0
	sum2 := 0
	for i := 1; i <= k; i++ {
		sum2 += prizeCount[i]
	}
	for i := k + 1; i <= max; i++ {
		sum2 += prizeCount[i]
		if sum2 > max2 {
			max2 = sum2
		}
		sum2 -= prizeCount[i-k]
	}
	fmt.Println("max1, sum2, max2:", max1, sum2, max2)
	return max1 + max2
}

type Input struct {
	prizePositions []int
	k              int
}

func main() {
	inputs := []Input{
		{
			prizePositions: []int{1, 1, 2, 2, 3, 3, 5},
			k:              2,
		},
		{
			prizePositions: []int{1, 2, 3, 4},
			k:              0,
		},
	}
	for _, input := range inputs {
		fmt.Println(maximizeWin(input.prizePositions, input.k))
	}
}
