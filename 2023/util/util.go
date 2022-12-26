package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadInput(fileName string) []string {
	content, fileReadErr := os.ReadFile(fileName)
	Check(fileReadErr)
	inputLines := strings.Split(string(content), "\n")
	return inputLines
}

func ReadGridInput(fileName string) [][]int {
	content, fileReadErr := os.ReadFile(fileName)
	Check(fileReadErr)
	inputLines := strings.Split(string(content), "\n")
	grid := make([][]int, 0)
	for _, line := range inputLines {
		row := make([]int, 0)
		for _, char := range line {
			row = append(row, ConvertAtoi(string(char)))
		}
		grid = append(grid, row)
	}
	return grid
}

func PrintGrid(grid [][]int) {
	for _, row := range grid {
		for _, element := range row {
			fmt.Print(element)
		}
		fmt.Println()
	}
}

func ConvertAtoi(str string) int {
	intVal, err := strconv.Atoi(str)
	Check(err)
	return intVal
}

type Stack []string

func (s Stack) Push(v string) Stack {
	return append(s, v)
}

func (s Stack) Pop() (Stack, string) {
	l := len(s)
	if l == 0 {
		return s, ""
	}
	return s[:l-1], s[l-1]
}
