package main

import "fmt"

func distinctIntegers(n int) int {
	mp := make(map[int]bool)
	mp[n] = true
	for i := 2; i <= n; i++ {
		fmt.Println("Day:", i)
		for num := range mp {
			fmt.Println(" - Num:", num)
			for j := 1; j < num; j++ {
				if num%j == 1 {
					mp[j] = true
					fmt.Println("  - New:", j)
				}
			}
		}
	}
	return len(mp)
}

func main() {
	inputs := []int{5, 3, 1000}
	for _, input := range inputs {
		fmt.Println(distinctIntegers(input))
	}
}
