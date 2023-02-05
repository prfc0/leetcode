package main

import "fmt"

func convert(s string, numRows int) string {
	n := len(s)
	if n == 1 || n <= numRows {
		return s
	}
	ans := make([]byte, n)
	interval := 2*numRows - 2
	i := 0
	for row := 0; row < numRows; row++ {
		ans[i] = s[row]
		i++
		if row == 0 || row == numRows-1 {
			for k := row + interval; k < n; k += interval {
				ans[i] = s[k]
				i++
			}
		} else {
			for k := row + interval; k < n+interval; k += interval {
				middle := k - 2*row
				if middle < n {
					ans[i] = s[middle]
					i++
				}
				if k < n {
					ans[i] = s[k]
					i++
				}
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
		/*
			{
				s:       "PAYPALISHIRING",
				numRows: 24,
			},
		*/
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
