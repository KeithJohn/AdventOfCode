package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputLines := readInputs("../input.txt")
	rows := len(inputLines)
	columns := len(inputLines[0])
	//Populate Grid
	grid := make([][]int, rows)
	for i, inputLine := range inputLines {
		grid[i] = make([]int, columns)
		for j, inputNum := range inputLine {
			num, convErr := strconv.Atoi(string(inputNum))
			check(convErr)
			grid[i][j] = num
		}
		fmt.Println(grid[i])
	}

	//fmt.Println(grid)

	lowPoints := make(map[string]bool)

	totalRisk := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			risk := isLowest(grid, i, j)
			if risk != -1 {
				totalRisk += risk
				fmt.Println("Point at row ", i, " and column ", j, " is the lowest point with risk ", risk, " total risk is ", totalRisk)
				key := strings.Join([]string{strconv.Itoa(i), strconv.Itoa(j)}, "_")
				lowPoints[key] = true
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			key := strings.Join([]string{strconv.Itoa(i), strconv.Itoa(j)}, "_")
			if lowPoints[key] == true {
				fmt.Print(string("\033[31m"), grid[i][j])
			} else {
				fmt.Print(string("\033[37m"), grid[i][j])
			}
		}
		fmt.Println()
	}
	fmt.Println("Total Risk: ", totalRisk)
}

func isLowest(grid [][]int, row int, column int) int {
	num := grid[row][column]
	if (row < 1 || num < grid[row-1][column]) && (row == len(grid)-1 || num < grid[row+1][column]) && (column < 1 || num < grid[row][column-1]) && (column == len(grid[0])-1 || num < grid[row][column+1]) {
		return num + 1
	} else {
		//Not lowest point
		return -1
	}
}

func readInputs(fileName string) []string {
	content, fileReadErr := os.ReadFile(fileName)
	check(fileReadErr)
	inputLines := strings.Split(string(content), "\n")
	return inputLines
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
