package main

import (
	"adventofcode/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := util.ReadInput("../input.txt")
	total := 0
	for _, line := range lines {
		ranges := strings.Split(line, ",")
		firstRange := parseRange(strings.Split(ranges[0], "-"))
		secondRange := parseRange(strings.Split(ranges[1], "-"))

		if isRangeFullyContained(firstRange, secondRange) {
			total++
		}
	}
	fmt.Print(total)
}

func isRangeFullyContained(firstRange []int, secondRange []int) bool {
	return (firstRange[0] <= secondRange[0] && firstRange[1] >= secondRange[1]) || (secondRange[0] <= firstRange[0] && secondRange[1] >= firstRange[1])
}

func parseRange(rangeStr []string) []int {
	rangeInt := make([]int, 2)

	num, err := strconv.Atoi(rangeStr[0])
	util.Check(err)
	rangeInt[0] = num

	num, err = strconv.Atoi(rangeStr[1])
	util.Check(err)
	rangeInt[1] = num

	return rangeInt
}
