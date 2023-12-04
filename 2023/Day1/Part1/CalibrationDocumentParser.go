package main

import (
	"adventofcode2023/util"
	"fmt"
	"unicode"
)

func main() {
	lines := util.ReadInput("../input.txt")
	sum := 0
	for _, line := range lines {
		num := parseCalibrationNumber(line)
		fmt.Println(num)
		sum += num
	}
	fmt.Println(sum)
}

func parseCalibrationNumber(line string) int {
	nums := make([]int, 2)
	num := 0
	//Find front number character
	for _, char := range line {
		if isNum(char) {
			nums[0] = util.ConvertAtoi(string(char))
			break
		}
	}

	//Find last number character
	for i := len(line) - 1; i >= 0; i-- {
		if isNum(rune(line[i])) {
			nums[1] = util.ConvertAtoi(string(line[i]))
			break
		}
	}

	num += nums[0] * 10
	num += nums[1]
	fmt.Println(num)
	return num
}

func isNum(char rune) bool {
	return unicode.IsDigit(char)
}
