package main

import (
	"adventofcode2023/util"
	"fmt"
	"unicode"
)

func main() {
	lines := util.ReadInput("../input.txt")
	sum := 0
	for x, line := range lines {
		num := parseCalibrationNumber(line)
		if num == -1 {
			fmt.Println(line, x)
			break
		} else {
			fmt.Println(num)
		}

		sum += num
	}
	fmt.Println(sum)
}

func parseCalibrationNumber(line string) int {

	firstNum, lastNum := -1, -1
	for i := 0; i < len(line); i++ {
		currChar := rune(line[i])
		if unicode.IsDigit(currChar) {
			//fmt.Println(string(currChar))
			if firstNum == -1 {
				firstNum = util.ConvertAtoi(string(currChar))
			} else {
				lastNum = util.ConvertAtoi(string(currChar))
			}
		}
		switch currChar {
		case 'o':
			if i+2 < len(line) && line[i:i+3] == "one" {
				//Found number!
				if firstNum == -1 {
					firstNum = 1
				} else {
					lastNum = 1
				}
			}
		case 't':
			if i+2 < len(line) && line[i:i+3] == "two" {
				//Found number!
				if firstNum == -1 {
					firstNum = 2
				} else {
					lastNum = 2
				}
			} else if i+4 < len(line) && line[i:i+5] == "three" {
				//Found number!
				if firstNum == -1 {
					firstNum = 3
				} else {
					lastNum = 3
				}
			}
		case 'f':
			if i+3 < len(line) && line[i:i+4] == "four" {
				//Found number!
				if firstNum == -1 {
					firstNum = 4
				} else {
					lastNum = 4
				}
			} else if i+3 < len(line) && line[i:i+4] == "five" {
				//Found number!
				if firstNum == -1 {
					firstNum = 5
				} else {
					lastNum = 5
				}
			}
		case 's':
			if i+2 < len(line) && line[i:i+3] == "six" {
				//Found number!
				if firstNum == -1 {
					firstNum = 6
				} else {
					lastNum = 6
				}
			} else if i+4 < len(line) && line[i:i+5] == "seven" {
				if firstNum == -1 {
					firstNum = 7
				} else {
					lastNum = 7
				}
			}
		case 'e':
			if i+4 < len(line) {
				fmt.Println(line[i : i+5])
			}

			if i+4 < len(line) && line[i:i+5] == "eight" {
				//Found number!

				if firstNum == -1 {
					firstNum = 8
				} else {
					lastNum = 8
				}

			}
		case 'n':
			if i+3 < len(line) && line[i:i+4] == "nine" {
				//Found number!
				if firstNum == -1 {
					firstNum = 9
				} else {
					lastNum = 9
				}
			}
		default:
			//no number found
		}
	}

	num := -1
	if firstNum == -1 {
		fmt.Println("FAIL---- ")
		fmt.Println(firstNum, lastNum)
	} else {
		if lastNum == -1 {
			lastNum = firstNum
		}
		num = (firstNum * 10) + lastNum
	}

	return num
}

//First thought is how do we determine whether we begin the number check or not
//Could check each letter, if the letter is the start of a number check the next n characters and see if it spells a number.
//	In this case, we would need to iterate through the entire string and track the last number as well
// 	Need to allow for digits as well as text numbers.

//First letters: o, t, f, s, e, n
//	o: one
//	t: two, three
//	f: four, five
//	s: six, seven
//	e: eight
//	n: nine

//Last letters: e, o, r, x, n, t
//	e: one, three, five, nine
//	o: two
//	r: four
// 	x: six
//	n: seven
//	t: eight

//Can we do this by iterating through the list? Possibly Recursively?
//Recursively build the letters backwards
// Check if the built string is contained with one of the number strings
// If not contained return current char, if contained, return the built string.
// Can see an issue where the current string is part of a number, but the first letter is the end of a number
// For example,   'asdfthreee'. In this case we would start with the last 'e', return, then add the next e to build the string 'ee'.
// This is the end of three, so we would return this. Next letter would be 'e' so we would add to get 'eee' and then fail. We would return the newest 'e' and miss the scenario

//Linear check
//Iterate through the string. If the letter is a start of a number, check the next n chars and compare against the potential numbers
//If yes, add to tracked list. Would track the first and constantly update the last found number
