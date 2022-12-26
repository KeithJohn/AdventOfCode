package main

import (
	"adventofcode/util"
	"fmt"
	"strconv"
)

func main() {
	lines := util.ReadInput("../input.txt")
	maxElfCalTotal := -1
	currElfCalTotal := 0
	for _, line := range lines {
		if line == "" {
			if currElfCalTotal > maxElfCalTotal {
				maxElfCalTotal = currElfCalTotal
			}
			currElfCalTotal = 0
		} else {
			lineNum, err := strconv.Atoi(line)
			util.Check(err)
			currElfCalTotal += lineNum
		}
	}

	fmt.Println(maxElfCalTotal)
}
