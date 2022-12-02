package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func main() {
	lines := readInputs("../input.txt")
	var ventGrid [1000][1000]int
	for _, line := range lines {
		fmt.Println("Line: ", line)
		if line.x1 == line.x2 {
			//vertical line
			y1, y2 := 0, 0
			if line.y1 < line.y2 {
				y1 = line.y1
				y2 = line.y2
			} else {
				y1 = line.y2
				y2 = line.y1
			}
			for y := y1; y <= y2; y++ {
				//Mark Grid
				ventGrid[line.x1][y] += 1
			}

		} else if line.y1 == line.y2 {
			//horizontal line
			x1, x2 := 0, 0
			if line.x1 < line.x2 {
				x1 = line.x1
				x2 = line.x2
			} else {
				x1 = line.x2
				x2 = line.x1
			}
			for x := x1; x <= x2; x++ {
				//Mark Grid
				ventGrid[x][line.y1] += 1
			}
		}
	}
	fmt.Println("vent Grid: ", ventGrid)

	twoVentsOrMore := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if ventGrid[i][j] >= 2 {
				twoVentsOrMore += 1
			}
		}
	}

	fmt.Println("Points with more than two vents: ", twoVentsOrMore)
}

func readInputs(fileName string) []Line {
	content, fileReadErr := os.ReadFile(fileName)
	check(fileReadErr)
	inputLines := strings.Split(string(content), "\n")
	var lines []Line
	for _, inputLine := range inputLines {
		inputs := strings.Split(inputLine, " -> ")
		startPoint := strings.Split(string(inputs[0]), ",")
		endPoint := strings.Split(string(inputs[1]), ",")
		var line Line
		x1, x1Err := strconv.Atoi(startPoint[0])
		check(x1Err)
		y1, y1Err := strconv.Atoi(startPoint[1])
		check(y1Err)
		x2, x2Err := strconv.Atoi(endPoint[0])
		check(x2Err)
		y2, y2Err := strconv.Atoi(endPoint[1])
		check(y2Err)

		line.x1 = x1
		line.y1 = y1
		line.x2 = x2
		line.y2 = y2

		lines = append(lines, line)
	}
	fmt.Println("Input lines: ", lines)
	return lines
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
