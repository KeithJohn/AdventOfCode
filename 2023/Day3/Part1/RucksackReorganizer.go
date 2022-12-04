package main

import (
	"adventofcode/util"
	"fmt"
	"unicode"
)

func main() {
	lines := util.ReadInput("../input.txt")
	total := 0
	for _, line := range lines {
		length := len(line)
		firstHalf := line[:length/2]
		secondHalf := line[length/2:]
		incorrectItem := findDuplicateItem(firstHalf, secondHalf)

		fmt.Println(incorrectItem)
		priority := getPriority(incorrectItem)
		fmt.Println(priority)
		total += priority
	}
	fmt.Println(total)
}

func findDuplicateItem(firstHalf string, secondHalf string) rune {
	firstHalfChars := make(map[rune]bool)
	for _, char := range firstHalf {
		firstHalfChars[char] = true
	}

	for _, char := range secondHalf {
		if _, ok := firstHalfChars[char]; ok {
			return char
		}
	}
	return '?'
}

func getPriority(incorrectItem rune) int {

	priority := 0
	if unicode.IsUpper(incorrectItem) {
		priority = int(incorrectItem-'0') + 10
	} else {
		priority = int(incorrectItem-'0') - 48
	}

	return priority
}
