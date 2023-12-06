package main

import (
	"adventofcode2023/util"
	"fmt"
	"strings"
)

func main() {
	lines := util.ReadInput("../input.txt")
	sum := 0
	for _, line := range lines {
		power := getGamePower(line)
		fmt.Println(power)
		sum += power
	}
	fmt.Println("Sum", sum)
}

func getGamePower(line string) int {
	//Determine minimum cube count and multiply together
	redMin, greenMin, blueMin := 0, 0, 0
	formattedLine := strings.Split(line, ":")[1]
	pulls := strings.Split(formattedLine, ";")
	for _, pull := range pulls {
		cubeCounts := strings.Split(strings.TrimSpace(pull), ",")
		for _, cubeCount := range cubeCounts {
			parts := strings.Split(strings.TrimSpace(cubeCount), " ")
			count, color := util.ConvertAtoi(parts[0]), parts[1]
			switch color {
			case "red":
				if redMin == 0 || count > redMin {
					redMin = count
				}
			case "green":
				if greenMin == 0 || count > greenMin {
					greenMin = count
				}
			case "blue":
				if blueMin == 0 || count > blueMin {
					blueMin = count
				}
			}
		}
	}
	fmt.Println(redMin, greenMin, blueMin)
	power := redMin * greenMin * blueMin
	return power
}
