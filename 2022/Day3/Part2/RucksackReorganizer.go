package main

import (
	"adventofcode/util"
	"fmt"
	"unicode"
)

func main() {
	lines := util.ReadInput("../input.txt")
	total := 0
	for i := 0; i < len(lines); i += 3 {
		badgeItem := findBadgeItem(lines[i], lines[i+1], lines[i+2])
		fmt.Println(string(badgeItem))
		pri := getPriority(badgeItem)
		fmt.Println(pri)
		total += pri
	}
	fmt.Println("total: ", total)
}

func findBadgeItem(firstSack string, secondSack string, thirdSack string) rune {
	firstSackChars := make(map[rune]bool)
	for _, char := range firstSack {
		firstSackChars[char] = true
	}

	firstSecondChars := make(map[rune]bool)
	for _, char := range secondSack {
		if _, ok := firstSackChars[char]; ok {
			firstSecondChars[char] = true
		}
	}

	for _, char := range thirdSack {
		if _, ok := firstSecondChars[char]; ok {
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
