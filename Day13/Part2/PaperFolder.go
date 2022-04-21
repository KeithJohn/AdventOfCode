package main

import (
	"fmt"
	"math"
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

	for _, fold := range folds {
		foldPaper(fold, points)
	}

	fmt.Println("Number of points: ", len(points))
	maxX, maxY := 0, 0
	for _, point := range points {
		if point.x > maxX {
			maxX = point.x
		}

		if point.y > maxY {
			maxY = point.y
		}
	}
	printPoints(points, maxX, maxY)
}

func printPoints(points map[string]Point, maxX int, maxY int) {

	for y := 0; y <= maxY; y++ {
		line := ""
		for x := 0; x <= maxX; x++ {
			key := strconv.Itoa(x) + "-" + strconv.Itoa(y)
			_, ok := points[key]
			if ok {
				line = line + "#"
			} else {
				if (x % 5) == 4 {
					line = line + "  "
				} else {
					line = line + " "
				}
			}

			if y == len(points)-1 {
				line = line + ""
			}
		}
		fmt.Println(line)
	}
}

func foldPaper(fold Fold, points map[string]Point) map[string]Point {
	for key, point := range points {
		if fold.direction == "x" {
			if point.x > fold.location {
				//Point to be folded left
				newX := fold.location - int(math.Abs(float64(point.x-fold.location)))
				//Update to new x value after fold
				point.x = newX
				//Remove and points that are in the same location
				delete(points, key)
				newKey := strconv.Itoa(point.x) + "-" + strconv.Itoa(point.y)
				points[newKey] = point
			}
		} else if fold.direction == "y" {
			if point.y > fold.location {
				//Point to be folded up
				newY := fold.location - int(math.Abs(float64(point.y-fold.location)))
				//Update to new y value after fold
				point.y = newY
				//Remove and points that are in the same location
				delete(points, key)
				newKey := strconv.Itoa(point.x) + "-" + strconv.Itoa(point.y)
				points[newKey] = point
			}
		}
	}
	return points
}

func getPointsAndFolds() (map[string]Point, []Fold) {
	inputLines := readInput("../input.txt")
	//inputLines := readInput("../simpleInput.txt")
	points := make(map[string]Point)
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
			x, error := strconv.Atoi(pointInp[0])
			check(error)
			y, error := strconv.Atoi(pointInp[1])
			check(error)
			var point Point
			point.x = x
			point.y = y
			key := pointInp[0] + "-" + pointInp[1]
			points[key] = point
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
