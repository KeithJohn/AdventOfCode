package main

import (
    "strconv"
    "fmt"
    "bufio"
    "os"
)

func main(){
    depthIncreases := calculateSumForWindow("input.txt", 1)
	fmt.Println(depthIncreases)
}

func calculateSumForWindow(fileName string, windowSize int) int{
	file, err := os.Open(fileName)
	check(err)
    defer file.Close()
	scanner := bufio.NewScanner(file)

	counter := 0
	previousDepthSum := 0
	depthIncreases := 0
	var window = make([]int, windowSize)
	for scanner.Scan(){
		currIndex := counter % windowSize
		num, _ := strconv.Atoi(scanner.Text())
		window[currIndex] = num
		if(counter >= windowSize){
			currSum := sumArray(window)
			if(currSum > previousDepthSum){
				depthIncreases += 1
			}
			previousDepthSum = currSum
		}
		counter++
	}
	return depthIncreases
}

func sumArray(arr []int) int {
	result := 0
	for _, v := range arr {
		result += v
	}
	return result
}

func check(e error){
    if e != nil{
        panic(e)
    }
}
