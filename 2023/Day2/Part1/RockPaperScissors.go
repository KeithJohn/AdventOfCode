package main

import (
	"adventofcode/util"
	"fmt"
	"strings"
)

func main() {

	//A/X - Rock - 1 point
	//B/Y - Paper - 2 point
	//C/Z - Scissors - 3 point

	//Lost - 0 point
	//Draw - 3 point
	//Win - 6 point
	lines := util.ReadInput("../input.txt")
	score := 0
	for _, line := range lines {
		choices := strings.Split(line, " ")
		oppChoice := choices[0]
		myChoice := choices[1]
		score += getScore(oppChoice, myChoice)
	}
	fmt.Print(score)
}

func getScore(oppChoice string, myChoice string) int {
	totalScore := 0
	switch myChoice {
	case "X":
		totalScore += 1
		if oppChoice == "A" {
			totalScore += 3
		} else if oppChoice == "C" {
			totalScore += 6
		}
	case "Y":
		totalScore += 2
		if oppChoice == "B" {
			totalScore += 3
		} else if oppChoice == "A" {
			totalScore += 6
		}
	case "Z":
		totalScore += 3
		if oppChoice == "C" {
			totalScore += 3
		} else if oppChoice == "B" {
			totalScore += 6
		}
	}
	return totalScore
}
