package main

import "fmt"

func convert(s string, numRows int) string {
	n := len(s)
	if numRows == 1 || n <= numRows {
		return s
	}
	fmt.Println(s, n)
	var isDiagonal bool
	matrix := make([][]byte, numRows)
	for i := range matrix {
		matrix[i] = make([]byte, n)
	}
	i, j := 0, 0
	for k := 0; k < n; k++ {
		if isDiagonal {
			matrix[i][j] = s[k]
			i = i - 1
			j++
			if i == 0 {
				isDiagonal = false
			}
		} else {
			matrix[i][j] = s[k]
			i++
			if i == numRows {
				if numRows > 2 {
					isDiagonal = true
				}
				i -= 2
				j++
			}
		}
	}
	numCols := n
	for row := range matrix {
		for col := range matrix[row] {
			ch := matrix[row][col]
			if ch == 0 {
				fmt.Print(" ")
			} else {
				fmt.Printf("%s", string(ch))
			}
		}
		fmt.Println()
	}
	ans := make([]byte, 0)
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			ch := matrix[row][col]
			if ch != 0 {
				ans = append(ans, ch)
			}
		}
	}
	return string(ans)
}

type Input struct {
	s       string
	numRows int
}

func main() {
	inputs := []Input{
		{
			s:       "PAYPALISHIRING",
			numRows: 3,
		},
		{
			s:       "PAYPALISHIRING",
			numRows: 4,
		},
		{
			s:       "PAYPALISHIRING",
			numRows: 24,
		},
		{
			s:       "ABC",
			numRows: 2,
		},
		{
			s:       "ABCD",
			numRows: 2,
		},
	}
	for _, input := range inputs {
		output := convert(input.s, input.numRows)
		fmt.Println(input.s, input.numRows, output)
	}
}
