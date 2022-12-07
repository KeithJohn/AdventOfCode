package main

import (
	"adventofcode/util"
	"fmt"
	"strings"
)

func main() {
	lines := util.ReadInput("../simpleInput.txt")
	for _, line := range lines {
		fmt.Println(line)
		start, end := 0, 4
		startFound := false

		for !startFound {
			window := line[start:end]
			if !strings.Contains(window, string(line[end+1])) {
				start++
				end++
				startFound = true
				continue
			} else {
				// strings.I
			}
			start++
			end++
		}
		fmt.Println(line[start:end])
		fmt.Println(end + 1)
	}
}

func findStartOfPacket(line string, windowSize int) int {
	charMap := make(map[string]int)
	for i := 0; i < windowSize; i++ {
		charMap[string(line[i])] = i
	}

	windowSize := 0
	for true {
		break
	}
	return index
}
