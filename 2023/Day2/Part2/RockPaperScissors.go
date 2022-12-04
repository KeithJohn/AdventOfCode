package main

import (
	"adventofcode/util"
	"fmt"
	"strings"
)

func main() {

	//A - Rock - 1 point
	//B - Paper - 2 point
	//C - Scissors - 3 point

	//X - Lose
	//Y - Draw
	//Z - Win

	//Lost - 0 point
	//Draw - 3 point
	//Win - 6 point
	lines := util.ReadInput("../input.txt")
	score := 0
	for _, line := range lines {
		choices := strings.Split(line, " ")
		oppChoice := choices[0]
		result := choices[1]
		score += getScore(oppChoice, result)
	}
	fmt.Print(score)
}

func getScore(oppChoice string, result string) int {
	totalScore := 0
	switch result {
	case "X":
		totalScore += 0
		if oppChoice == "A" {
			totalScore += 3
		} else if oppChoice == "B" {
			totalScore += 1
		} else {
			totalScore += 2
		}
	case "Y":
		totalScore += 3
		if oppChoice == "A" {
			totalScore += 1
		} else if oppChoice == "B" {
			totalScore += 2
		} else {
			totalScore += 3
		}
	case "Z":
		totalScore += 6
		if oppChoice == "A" {
			totalScore += 2
		} else if oppChoice == "B" {
			totalScore += 3
		} else {
			totalScore += 1
		}
	}
	return totalScore
}
