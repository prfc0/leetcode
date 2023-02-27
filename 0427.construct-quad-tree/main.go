package main

import "fmt"

/**
 * Definition for a QuadTree node.
 */

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	n := len(grid)
	return constructQuadTree(grid, 0, 0, n-1, n-1)
}

func isLeafNode(grid [][]int, rowStart, colStart, rowEnd, colEnd int) (int, bool) {
	gridValue := grid[rowStart][colStart]
	for row := rowStart; row <= rowEnd; row++ {
		for col := colStart; col <= colEnd; col++ {
			if grid[row][col] != gridValue {
				return gridValue, false
			}
		}
	}
	return gridValue, true
}
func constructQuadTree(grid [][]int, rowStart, colStart, rowEnd, colEnd int) *Node {
	if rowStart > rowEnd || colStart > colEnd {
		return nil
	}
	leafValue, check := isLeafNode(grid, rowStart, colStart, rowEnd, colEnd)
	if check {
		if leafValue == 1 {
			return &Node{IsLeaf: true, Val: true}
		}
		return &Node{IsLeaf: true}
	}
	rowMid := (rowStart + rowEnd) / 2
	colMid := (colStart + colEnd) / 2
	return &Node{
		TopLeft:     constructQuadTree(grid, rowStart, colStart, rowMid, colMid),
		TopRight:    constructQuadTree(grid, rowStart, colMid+1, rowMid, colEnd),
		BottomLeft:  constructQuadTree(grid, rowMid+1, colStart, rowEnd, colMid),
		BottomRight: constructQuadTree(grid, rowMid+1, colMid+1, rowEnd, colEnd),
	}
}

func main() {
	fmt.Println("Check https://leetcode.com/problems/construct-quad-tree/")
}
