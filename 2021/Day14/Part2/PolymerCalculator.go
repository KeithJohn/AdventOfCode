package main

import (
	"fmt"
	"os"
	"strings"
	"math"
)

func main() {
	template, pairInsertionRules := getTemplateAndPairInsertionRules()
	calculatePolymer(template, pairInsertionRules,40)
}

func calculatePolymer(template string, pairInsertionRules map[string]string, numOfSteps int) int {
	//Create map tracking pairs and counts
	pairCount := make(map[string]int)
	for key := range pairInsertionRules {
		pairCount[key] = 0
	}

	for i := 0; i < len(template)-1; i++ {
		pair := string(template[i]) + string(template[i+1])
		pairCount[pair] = pairCount[pair] + 1
	}

	for pair, count := range pairCount{
	if count > 0{
		fmt.Println(pair, "-", count)
	}
} 

	for step := 0; step < numOfSteps; step++{
		stepPairCount := make(map[string]int)
		//fmt.Println(step)
		//Loop through keys and update numbers of pairs
		for pair, count := range pairCount{
			if pairCount[pair] > 0 {
				newPairFront := string(pair[0]) + pairInsertionRules[pair]
				newPairBack := pairInsertionRules[pair] + string(pair[1])
				stepPairCount[newPairFront] = stepPairCount[newPairFront] + count
				stepPairCount[newPairBack] = stepPairCount[newPairBack] + count
			}
		}
		pairCount = stepPairCount
	}

	for pair, count := range pairCount{
		if count > 0{
			fmt.Println(pair, "-", count)
		}
	} 

	charCount := make(map[string]int)
	fmt.Println("-------------")
	for pair, count := range pairCount{
		if count > 0{
			charCount[string(pair[0])] = charCount[string(pair[0])] + count
			charCount[string(pair[1])] = charCount[string(pair[1])] + count
		}
	} 
	
	maxChar := ""
	maxCount := 0
	minChar := ""
	minCount := math.MaxInt64
	firstChar := string(template[0])
	lastChar := string(template[len(template) - 1])
	for char, count := range charCount {
		charCount[char] = count / 2
		if(char == firstChar || char == lastChar){
			charCount[char] += 1
		}
		fmt.Println("-----------")
		fmt.Println(char, "-", count)
		fmt.Println(char, "-", charCount[char])

		if charCount[char] > maxCount {
			maxChar = char
			maxCount = charCount[char]
		}
		if charCount[char] < minCount {
			minChar = char
			minCount = charCount[char]
		}
	}
	fmt.Println("Max ", maxChar, "-", maxCount)
	fmt.Println("Min ", minChar, "-", minCount)
	fmt.Println(maxChar, " - ", minChar, " = ", maxCount - minCount)
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
