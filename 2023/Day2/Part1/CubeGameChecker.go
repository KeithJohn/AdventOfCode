package main

import (
	"adventofcode2023/util"
	"fmt"
	"strings"
)

func main() {
	lines := util.ReadInput("../input.txt")
	fmt.Println(lines[0])
	sum := 0
	for id, line := range lines {
		if checkGame(line) {
			sum += id + 1
			fmt.Println(id)
		}
	}
	fmt.Println("Sum", sum)
}

func checkGame(line string) bool {
	//12 red, 13 green, 14 blue
	// Must determine if the game is possible based on the given cubes
	formattedLine := strings.Split(line, ":")[1]
	pulls := strings.Split(formattedLine, ";")
	for _, pull := range pulls {
		cubeCounts := strings.Split(strings.TrimSpace(pull), ",")
		for _, cubeCount := range cubeCounts {
			parts := strings.Split(strings.TrimSpace(cubeCount), " ")
			count, color := util.ConvertAtoi(parts[0]), parts[1]
			switch color {
			case "red":
				if count > 12 {
					return false
				}
			case "green":
				if count > 13 {
					return false
				}
			case "blue":
				if count > 14 {
					return false
				}
			}
		}
	}
	return true
}
