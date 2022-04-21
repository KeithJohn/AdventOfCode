package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Fold struct {
	direction string
	location  int
}

func main() {
	points, folds := getPointsAndFolds()

	for _, readPoint := range points {
		fmt.Println(readPoint.x, ",", readPoint.y)
	}

	for _, readFold := range folds {
		fmt.Println(readFold.direction, ": ", readFold.location)
	}
}

func getPointsAndFolds() ([]Point, []Fold) {
	inputLines := readInput("../input.txt")
	var points []Point
	var folds []Fold
	for _, line := range inputLines {
		if strings.HasPrefix(line, "fold along") {
			//Fold
			foldLine := strings.TrimPrefix(line, "fold along ")
			foldInp := strings.Split(strings.TrimRight(foldLine, "\r"), "=")
			var fold Fold
			fold.direction = foldInp[0]
			loc, error := strconv.Atoi(foldInp[1])
			check(error)
			fold.location = loc
			folds = append(folds, fold)

		} else if line != "\r" {
			pointInp := strings.Split(strings.TrimRight(line, "\r"), ",")
			fmt.Println(pointInp)
			x, error := strconv.Atoi(pointInp[0])
			check(error)
			y, error := strconv.Atoi(pointInp[1])
			check(error)
			var point Point
			point.x = x
			point.y = y
			points = append(points, point)
		}
	}
	return points, folds
}

func readInput(fileName string) []string {
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
