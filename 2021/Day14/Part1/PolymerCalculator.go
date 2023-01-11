package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	template, pairInsertionRules := getTemplateAndPairInsertionRules()
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
		}
		newTemplate.WriteString(string(template[len(template)-1]))
		template = newTemplate.String()
	}
	solution := calcSolution(template)
	return solution
}

func calcSolution(template string) int {
	templateChars := strings.Split(template, "")
	sort.Strings(templateChars)
	currCount := 1
	maxCount := 0
	maxChar := ""
	minCount := math.MaxInt64
	minChar := ""
	for i := 1; i < len(templateChars); i++ {
		if templateChars[i-1] == templateChars[i] {
			currCount++
		} else {
			if currCount > maxCount {
				maxCount = currCount
				maxChar = templateChars[i-1]
			}
			if currCount < minCount {
				minCount = currCount
				minChar = templateChars[i-1]
			}
			currCount = 0
		}
	}
	fmt.Println("Max: ", maxChar, "-", maxCount)
	fmt.Println("Min: ", minChar, "-", minCount)

	return maxCount - minCount
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
