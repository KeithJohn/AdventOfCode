package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Line struct {
	inputCode  []string
	outputCode []string
}

//TODO: Fix this up. This is very sloppy and innefficent
func main() {
	lines := readInputs("../input.txt")
	sum := 0
	fmt.Println(len(lines))
	for i, line := range lines {
		numberCode := getEasyDigits(line.inputCode)
		for _, input := range line.inputCode {
			if len(input) == 6 {
				number := findSixSegNumberCode(input, numberCode)
				numberCode[number] = input
			}
		}

		for _, input := range line.inputCode {
			if len(input) == 5 {
				number := findFiveSegNumberCode(input, numberCode)
				numberCode[number] = input
			}
		}

		//Calculate output
		outputNum := ""
		for _, output := range line.outputCode {
			if equals(output, numberCode[0]) {
				outputNum = outputNum + "0"
			} else if equals(output, numberCode[1]) {
				outputNum = outputNum + "1"
			} else if equals(output, numberCode[2]) {
				outputNum = outputNum + "2"
			} else if equals(output, numberCode[3]) {
				outputNum = outputNum + "3"
			} else if equals(output, numberCode[4]) {
				outputNum = outputNum + "4"
			} else if equals(output, numberCode[5]) {
				outputNum = outputNum + "5"
			} else if equals(output, numberCode[6]) {
				outputNum = outputNum + "6"
			} else if equals(output, numberCode[7]) {
				outputNum = outputNum + "7"
			} else if equals(output, numberCode[8]) {
				outputNum = outputNum + "8"
			} else if equals(output, numberCode[9]) {
				outputNum = outputNum + "9"
			}
		}
		fmt.Println("Output Number ", i, ": ", outputNum)
		outputSum, convErr := strconv.Atoi(outputNum)
		check(convErr)
		sum += outputSum
	}
	fmt.Println(sum)
}

func equals(inputCode string, numberCode string) bool {
	if len(inputCode) != len(numberCode) {
		return false
	}

	for _, char := range numberCode {
		if !strings.Contains(inputCode, string(char)) {
			return false
		}
	}
	return true
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

func getEasyDigits(inputCodes []string) map[int]string {
	numberCodes := make(map[int]string)
	for _, input := range inputCodes {
		if len(input) == 2 {
			//Input is 1
			numberCodes[1] = input
		} else if len(input) == 4 {
			//Input is 4
			numberCodes[4] = input
		} else if len(input) == 3 {
			//Input is 7
			numberCodes[7] = input
		} else if len(input) == 7 {
			//Input is 8
			numberCodes[8] = input
		}
	}
	return numberCodes
}

func findSixSegNumberCode(inputCode string, numberCodes map[int]string) int {
	oneCode := numberCodes[1]
	fourCode := numberCodes[4]

	if strings.Contains(inputCode, string(fourCode[0])) && strings.Contains(inputCode, string(fourCode[1])) && strings.Contains(inputCode, string(fourCode[2])) && strings.Contains(inputCode, string(fourCode[3])) {
		//Input is 9
		return 9
	} else {
		if strings.Contains(inputCode, string(oneCode[0])) && strings.Contains(inputCode, string(oneCode[1])) {
			//Input is 0
			return 0
		} else {
			//Input is 6
			return 6
		}
	}
}

func findFiveSegNumberCode(inputCode string, numberCodes map[int]string) int {
	oneCode := numberCodes[1]
	sixCode := numberCodes[6]

	if strings.Contains(inputCode, string(oneCode[0])) && strings.Contains(inputCode, string(oneCode[1])) {
		//Input is 3
		return 3
	} else {
		for _, char := range sixCode {
			oneCode = strings.ReplaceAll(oneCode, string(char), "")
		}
		if strings.Contains(inputCode, string(oneCode[0])) {
			//Input is 2
			return 2
		} else {
			return 5
		}
	}
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
