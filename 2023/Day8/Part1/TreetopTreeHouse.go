package main

import (
	"adventofcode/util"
	"fmt"
)

func main() {
	grid := util.ReadGridInput("../input.txt")
	visibleCount := findVisibleTreeCount(grid)
	fmt.Println(visibleCount)
}

func findVisibleTreeCount(grid [][]int) int {
	visibleGrid := make([][]int, 0)
	for i := 0; i < len(grid); i++ {
		visibleGrid = append(visibleGrid, make([]int, len(grid[i])))
	}

	//left to right
	for i := 0; i < len(grid); i++ {
		currTallestTree := grid[i][0]
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > currTallestTree {
				currTallestTree = grid[i][j]
				visibleGrid[i][j] = 1
			}
		}
	}

	//0 012
	//1 012
	//2 012

	//right to left
	for i := 0; i < len(grid); i++ {
		currTallestTree := grid[i][len(grid[i])-1]
		for j := len(grid[i]) - 1; j > 0; j-- {
			if grid[i][j] > currTallestTree {
				currTallestTree = grid[i][j]
				visibleGrid[i][j] = 1
			}
		}
	}

	//top to bottom
	for i := 0; i < len(grid); i++ {
		currTallestTree := grid[0][i]
		for j := 0; j < len(grid[i]); j++ {
			if grid[j][i] > currTallestTree {
				currTallestTree = grid[j][i]
				visibleGrid[j][i] = 1
			}
		}
	}

	//bottom to top
	for i := 0; i < len(grid); i++ {
		currTallestTree := grid[len(grid)-1][i]
		for j := len(grid) - 1; j > 0; j-- {
			if grid[j][i] > currTallestTree {
				currTallestTree = grid[j][i]
				visibleGrid[j][i] = 1
			}
		}
	}

	visibleCount := 0
	for i, row := range visibleGrid {
		for j := range row {
			if i == 0 || j == 0 || i == len(visibleGrid)-1 || j == len(row)-1 {
				visibleGrid[i][j] = 1
			}
			visibleCount += visibleGrid[i][j]
		}
	}
	util.PrintGrid(visibleGrid)
	return visibleCount
}
