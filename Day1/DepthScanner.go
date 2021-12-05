package main

import (
    "strconv"
    "fmt"
	"strings"
    "os"
)

func main(){
	windowSize := 1
	window := make([]int, windowSize)

	depthIncreases := 0

	content, fileReadErr := os.ReadFile("input.txt")
	Check(fileReadErr)
	values := strings.Split(string(content), "\n")
	for i, line := range values{
		value, convErr := strconv.Atoi(strings.TrimSpace(line))
		Check(convErr)

		oldValue := window[i%windowSize]
		window[i%windowSize] = value
		if(i < windowSize){
			continue
		}

		if(oldValue < value){
			depthIncreases++
		}
	}
	
	fmt.Println(depthIncreases)
}

func Check(err error){
	if (err != nil){
		panic(err)
	}
}