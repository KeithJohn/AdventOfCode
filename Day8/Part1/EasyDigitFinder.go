package main

import (
	"fmt"
	"os"
	"strings"
)

type Line struct {
	inputCode  []string
	outputCode []string
}

func main() {
	lines := readInputs("../input.txt")

	numOfEasyDigits := 0
	for _, line := range lines {
		for _, input := range line.outputCode {
			if len(input) == 2 || len(input) == 4 || len(input) == 3 || len(input) == 7 {
				numOfEasyDigits += 1
			}
		}
	}
	fmt.Println(numOfEasyDigits)
}

func readInputs(fileName string) []Line {
	content, fileReadErr := os.ReadFile(fileName)
	check(fileReadErr)
	inputLines := strings.Split(string(content), "\n")
	var lines []Line
	for _, inputLine := range inputLines {
		var line Line
		lineData := strings.Split(string(inputLine), " | ")
		line.inputCode = strings.Split(string(lineData[0]), " ")
		line.outputCode = strings.Split(string(lineData[1]), " ")
		lines = append(lines, line)
	}
	//fmt.Println("Input lines: ", lines)
	return lines
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

//0 - 6 segments, contains all but 1 characters from 8
//1 - 2 segments, unique
//2 - 5 segments, TODO
//3 - 5 segments, TODO
//4 - 4 segments, unique
//5 - 5 segments, TODO
//6 - 6 segments, contains all but 1 characters from 8, does not share any with 1
//7 - 3 segments, unique
//8 - 7 segments, unique
//9 - 6 segments, contains all of 4 characters

//Code contains 2, 4, 3, or 7 segments
//This is a unique number of segments. Either 1, 4. 7, or 8

// Code contains 6 segments
// Code is 9 if code contains 6 segments and shares all segments with 4
// Code is 0 if code contains 6 segments and does not share all segments with 4 and shares all segments with 1
// Code is 6 if code contains 6 segments and does not share all segments with 4 and does not share all segments with 1

//Code contains 5 segments
// code is 3 if code contains 5 segments and shares all of segments with 1
// code is 2 if code contains 5 segments and
// code is 5 if code contains 5 segments and
