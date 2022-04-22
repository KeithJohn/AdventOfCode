package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	template, pairInsertionRules := getTemplateAndPairInsertionRules()
	fmt.Println(template)
	//fmt.Println(pairInsertionRules)
	fmt.Println(calculatePolymer(template, pairInsertionRules, 10))
}

func calculatePolymer(template string, pairInsertionRules map[string]string, numOfSteps int) int {
	var newTemplate strings.Builder
	for step := 0; step < numOfSteps; step++ {
		newTemplate.Reset()
		for i := 0; i < len(template)-1; i++ {
			pair := string(template[i]) + string(template[i+1])
			char, ok := pairInsertionRules[pair]
			if !ok {
				fmt.Println("ERROR: missing pair")
			}
			newTemplate.WriteString(string(pair[0]))
			newTemplate.WriteString(char)
			// fmt.Println(newTemplate.String())
			// fmt.Println(pair + " -> " + char)
			// fmt.Println("-----")
			//newTemplate = newTemplate[:i] + char + newTemplate[i:]
		}
		newTemplate.WriteString(string(template[len(template)-1]))
		template = newTemplate.String()
	}
	//fmt.Println(newTemplate.String())
	max := mostCommon(template)
	min := leastCommon(template)
	return max - min
}

//TODO: Fix this
func leastCommon(template string) int {
	runes := []rune(template)
	sort.Slice(runes, func(i int, j int) bool { return runes[i] < runes[j] })
	currCount := 0
	minCount := len(template)
	var result rune
	for i := 1; i < len(runes); i++ {
		if runes[i] == runes[i-1] {
			currCount++
		} else {
			if currCount < minCount {
				minCount = currCount
				result = runes[i-1]
			}
		}
	}
	fmt.Println(string(result) + " " + strconv.Itoa(minCount))
	return minCount
}

//TODO: Fix this
func mostCommon(template string) int {
	runes := []rune(template)
	sort.Slice(runes, func(i int, j int) bool { return runes[i] < runes[j] })
	fmt.Println(string(runes))
	currCount := 0
	maxCount := 0
	var result rune
	for i := 1; i < len(runes); i++ {
		if runes[i] == runes[i-1] {
			currCount++
		} else {
			if currCount > maxCount {
				maxCount = currCount
				result = runes[i-1]
			}
		}
	}
	fmt.Println(string(result) + " " + strconv.Itoa(maxCount))
	return maxCount
}
func getTemplateAndPairInsertionRules() (string, map[string]string) {
	inputLines := readInput("../input.txt")
	var template string
	pairInsertionRules := make(map[string]string)
	for lineNum, line := range inputLines {
		if lineNum == 0 {
			template = strings.TrimSpace(line)
		} else if lineNum == 1 {
			continue
		} else {
			pairRule := strings.Split(line, " -> ")
			pairInsertionRules[strings.TrimSpace(pairRule[0])] = strings.Trim(pairRule[1], "\r")
		}
	}
	return template, pairInsertionRules
}

func readInput(fileName string) []string {
	content, fileReadErr := os.ReadFile(fileName)
	check(fileReadErr)
	inputLines := strings.Split(string(content), "\n")
	return inputLines
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
