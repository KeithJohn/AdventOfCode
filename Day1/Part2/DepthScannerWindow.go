package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputs, err := ReadInputs()
	check(err)
	numOfIncreases := 0
	for i := 0; i < len(inputs)-3; i++ {
		win1Sum := inputs[i] + inputs[i+1] + inputs[i+2]
		fmt.Println("window 1", inputs[i], inputs[i+1], inputs[i+2])
		win2Sum := inputs[i+1] + inputs[i+2] + inputs[i+3]
		fmt.Println("window 2", inputs[i+1], inputs[i+2], inputs[i+3])
		fmt.Println(win1Sum, " ", win2Sum)
		if win2Sum > win1Sum {
			numOfIncreases += 1
		}
		fmt.Println(numOfIncreases)
	}
	fmt.Println(numOfIncreases)
}

func ReadInputs() ([]int, error) {
	file, fileErr := os.Open("inputP2.txt")
	check(fileErr)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		check(err)
		result = append(result, x)
	}
	return result, scanner.Err()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
