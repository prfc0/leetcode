package main

import (
	"fmt"
	"sort"
)

type sortByScore struct {
	score [][]int
	k     int
}

func (s sortByScore) Len() int           { return len(s.score) }
func (s sortByScore) Less(i, j int) bool { return s.score[i][s.k] > s.score[j][s.k] }
func (s sortByScore) Swap(i, j int)      { s.score[i], s.score[j] = s.score[j], s.score[i] }

func sortTheStudents(score [][]int, k int) [][]int {
	s := sortByScore{score: score, k: k}
	sort.Sort(sortByScore(s))
	return s.score
}

func main() {
	// matrix := [][]int{{10, 6, 9, 1}, {7, 5, 11, 2}, {4, 8, 3, 15}}
	// k := 2
	matrix := [][]int{{3, 4}, {5, 6}}
	k := 0
	fmt.Println(sortTheStudents(matrix, k))
}
