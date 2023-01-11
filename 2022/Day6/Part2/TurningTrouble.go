package main

import (
	"adventofcode/util"
	"fmt"
)

func main() {
	lines := util.ReadInput("../input.txt")
	for _, line := range lines {
		fmt.Println(line)
		fmt.Println(findStartOfPacket(line, 14))
	}
}

func findStartOfPacket(line string, windowSize int) int {
	start := -1
	end := windowSize - 1
	hasDuplicate := true
	for hasDuplicate {
		start++
		end++
		hasDuplicate = checkDuplicate(line[start:end])
	}
	fmt.Println("No duplicate at ", line[start:end])

	return end
}

func checkDuplicate(s string) bool {
	charMap := make(map[string]int)
	for _, char := range s {
		if _, ok := charMap[string(char)]; ok {
			return true
		} else {
			charMap[string(char)] = 1
		}
	}
	return false
}
