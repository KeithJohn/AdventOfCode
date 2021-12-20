package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Stack []rune

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

//Push a new value onto the stack
func (s *Stack) Push(r rune) {
	*s = append(*s, r) // Simply append the new value to the end of the stack
}

//Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return 'X', false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func main() {
	scoreMap := make(map[rune]int)
	scoreMap[')'] = 1
	scoreMap[']'] = 2
	scoreMap['}'] = 3
	scoreMap['>'] = 4
	inputLines := readInput("../input.txt")
	var syntaxAdditions []string
line:
	for _, inputLine := range inputLines {
		var stack Stack
		for _, char := range inputLine {
			if char == '(' || char == '[' || char == '{' || char == '<' {
				//Open syntax. Add to stack
				stack.Push(char)
			} else if char == ')' || char == ']' || char == '}' || char == '>' {
				//Close syntax. Pop from stack if possible and compare
				top, hasNext := stack.Pop()
				corresponding := getCorresponding(top)
				if !hasNext || corresponding != char {
					//Syntax error. Skip
					continue line
				}
			} else {
				//Error. Invalid Character
				panic("Invalid character!!")
			}
		}

		if !stack.IsEmpty() {
			addition := getSyntaxAddition(stack)
			fmt.Println("Adding addition ", addition, " for stack:")
			syntaxAdditions = append(syntaxAdditions, addition)
		}
	}

	additionScores := make([]int, len(syntaxAdditions))
	for i, addition := range syntaxAdditions {
		additionScore := 0
		for _, char := range addition {
			additionScore *= 5
			additionScore += scoreMap[char]
		}
		additionScores[i] = additionScore
	}

	sort.Ints(additionScores)
	middleScore := additionScores[len(additionScores)/2]
	fmt.Println("Auto complete score: ", middleScore)
}

func readInput(fileName string) []string {
	content, fileReadErr := os.ReadFile(fileName)
	check(fileReadErr)
	inputLines := strings.Split(string(content), "\n")
	return inputLines
}

func printStack(s Stack) {
	for _, char := range s {
		fmt.Print(string(char), " ")
	}
}

func getSyntaxAddition(s Stack) string {
	printStack(s)
	var b strings.Builder
	b.Grow(len(s))
	for !s.IsEmpty() {
		char, _ := s.Pop()
		b.WriteRune(getCorresponding(char))
	}
	return b.String()
}

func getCorresponding(char rune) rune {
	switch char {
	case '(':
		return ')'
	case '[':
		return ']'
	case '{':
		return '}'
	case '<':
		return '>'
	default:
		return 'X'
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
