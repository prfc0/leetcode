package main

import "fmt"

// 0   1   2    3    4          n-1
// 1 + 4 + 8 + 12 + 16 ... + 4*(n-1)
// 1 + 4(1+2+3+..n-1)
func coloredCells(n int) int64 {
	return int64(n*(n-1)*2 + 1)
}

func main() {
	for i := 10000; i <= 100000; i++ {
		fmt.Println(i, coloredCells(i))
	}
}
