package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	positions := readInputs("../input.txt")
	max := getMax(positions)
	lowestUsage := int(^uint(0) >> 1)
	index := -1
	for i := 0; i <= max; i++ {
		usage := calcFuelUsage(positions, i)
		if usage < lowestUsage {
			lowestUsage = usage
			index = i
		}
	}

	fmt.Println("Inefficient lowest usage: ", lowestUsage, ", index: ", index)
	// sum := 0
	// for _, pos := range positions {
	// 	sum += pos
	// }
	// fmt.Println("sum: ", sum, ", number of subs: ", len(positions))
	// bestIndex := int(sum / len(positions))
	// fmt.Println("Most efficient position: ", bestIndex)

	// fuelUsage := 0
	// for _, pos := range positions {
	// 	fuelUsage += int(math.Abs(float64(pos - bestIndex)))
	// }
	// fmt.Println("Fuel Usage for index ", bestIndex, ": ", fuelUsage)
}

func getMax(arr []int) int {
	max := -1
	for _, e := range arr {
		if e > max {
			max = e
		}
	}
	return max
}

func calcFuelUsage(positions []int, index int) int {
	fuelUsage := 0
	for _, pos := range positions {
		if pos > index {
			fuelUsage += pos - index
		} else if pos < index {
			fuelUsage += index - pos
		}
	}
	return fuelUsage
}

func readInputs(fileName string) []int {
	content, fileReadErr := os.ReadFile(fileName)
	check(fileReadErr)
	inputs := strings.Split(string(content), ",")
	var horizPos []int
	for _, input := range inputs {
		num, convErr := strconv.Atoi(input)
		check(convErr)
		horizPos = append(horizPos, num)
	}
	return horizPos
}

func check(err error) {
	if err != nil {
		panic(err)
	}

}
