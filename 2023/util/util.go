package util

import (
	"os"
	"strings"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadInput(fileName string) []string {
	content, fileReadErr := os.ReadFile(fileName)
	Check(fileReadErr)
	inputLines := strings.Split(string(content), "\n")
	return inputLines
}

type Stack []string

func (s Stack) Push(v string) Stack {
	return append(s, v)
}

func (s Stack) Pop() (Stack, string) {
	l := len(s)
	if l == 0 {
		return s, ""
	}
	return s[:l-1], s[l-1]
}
