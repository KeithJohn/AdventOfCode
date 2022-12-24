package main

import (
	"adventofcode/util"
	"fmt"
)

func main() {
	grid := util.ReadGridInput("../input.txt")
	scenicScore := findBestScenicScore(grid)
	fmt.Println(scenicScore)
}

func findBestScenicScore(grid [][]int) int {
	topScenicScore := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			currHeight := grid[i][j]

			//Check Up
			checkIndex := i - 1
			upCount := 0
			for checkIndex >= 0 {
				upCount++
				if grid[checkIndex][j] >= currHeight {
					break
				}
				checkIndex--
			}

			//Check Right
			checkIndex = j + 1
			rightCount := 0
			for checkIndex < len(grid[i]) {
				rightCount++
				if grid[i][checkIndex] >= currHeight {
					break
				}
				checkIndex++
			}

			//Check Down
			checkIndex = i + 1
			downCount := 0
			for checkIndex < len(grid) {
				downCount++
				if grid[checkIndex][j] >= currHeight {
					break
				}
				checkIndex++
			}

			//Check Left
			checkIndex = j - 1
			leftCount := 0
			for checkIndex >= 0 {
				leftCount++
				if grid[i][checkIndex] >= currHeight {
					break
				}
				checkIndex--
			}

			scenicScore := upCount * leftCount * downCount * rightCount
			if scenicScore > topScenicScore {
				topScenicScore = scenicScore
				//fmt.Println(i, j)
			}
		}
	}
	return topScenicScore
}
