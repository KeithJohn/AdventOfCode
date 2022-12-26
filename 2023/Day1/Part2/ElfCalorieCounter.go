package main

import (
	"adventofcode/util"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	lines := util.ReadInput("../input.txt")
	topElfs := [3]int{-1, -1, -1}
	currElfCalTotal := 0

	for _, line := range lines {
		if line == "" {
			if currElfCalTotal > topElfs[0] {
				topElfs[0] = currElfCalTotal
				sort.Ints(topElfs[:])
			}
			currElfCalTotal = 0
		} else {
			lineNum, err := strconv.Atoi(line)
			util.Check(err)
			currElfCalTotal += lineNum
		}
	}

	fmt.Println(topElfs)
	fmt.Println(topElfs[0] + topElfs[1] + topElfs[2])
}
