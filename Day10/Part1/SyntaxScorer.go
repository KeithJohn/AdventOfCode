package main

import (
	"fmt"
	"os"
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
	scoreMap[')'] = 3
	scoreMap[']'] = 57
	scoreMap['}'] = 1197
	scoreMap['>'] = 25137
	inputLines := readInput("../input.txt")
	invalidSyntaxList := make([]rune, len(inputLines))
line:
	for _, inputLine := range inputLines {
		var stack Stack
		for _, char := range inputLine {
			if char == '(' || char == '[' || char == '{' || char == '<' {
				//Open syntax. Add to stack
				fmt.Println("Adding to stack")
				printStack(stack)
				stack.Push(char)
			} else if char == ')' || char == ']' || char == '}' || char == '>' {
				//Close syntax. Pop from stack if possible and compare
				top, hasNext := stack.Pop()
				corresponding := getCorresponding(top)
				if !hasNext || corresponding != char {
					//Syntax error. Track char and jump to next line
					fmt.Println(hasNext)
					fmt.Println("Invalid syntax: Expecting :", string(corresponding), " but got ", string(char))
					invalidSyntaxList = append(invalidSyntaxList, char)
					continue line
				}
			} else {
				//Error. Invalid Character
				panic("Invalid character on the ")
			}
		}
	}

	score := 0
	for _, char := range invalidSyntaxList {
		score += scoreMap[char]
	}
	fmt.Println("Syntax Score: ", score)
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
