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
