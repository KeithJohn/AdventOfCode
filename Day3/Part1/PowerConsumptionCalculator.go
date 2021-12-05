package main

import (
	"fmt"
	"os"
	"strings"
	"math"
)

//TODO: This is ugly. Fix it
func main(){
	content, readFileErr := os.ReadFile("../input.txt")
	Check(readFileErr)
	values := strings.Split(string(content), "\n")
	//Array containing number of '1's in each index
	oneCounts := make([]int, len(values[0]))
	totalLines := len(values)
	fmt.Println(totalLines)
	fmt.Println(oneCounts)
	for _, line := range values {
		for pos, char := range line{
			if(char == '1'){
				oneCounts[pos]++
			}
		}
	}

	gamma := make([]int, len(oneCounts))
	epsilon := make([]int, len(oneCounts))

	for i, value := range oneCounts {
		if value > totalLines / 2{
			gamma[i] = 1
			epsilon[i] = 0
		}else{
			gamma[i] = 0
			epsilon[i] = 1
		}
	}
	fmt.Println(oneCounts)
	fmt.Println(gamma)
	fmt.Println(epsilon)
	gammaValue := calculateBinaryValue(gamma)
	epsilonValue := calculateBinaryValue(epsilon)
	solution := gammaValue * epsilonValue
	fmt.Println(solution)
}

func Check(err error){
	if(err != nil){
		panic(err)
	}
}

func calculateBinaryValue(binary []int) int{
	sum := float64(0)
	pow := 0
	for i:= len(binary) - 1; i >= 0; i--{
		value := binary[i]
		fmt.Println(value)
		sum += math.Pow(float64(2 * value), float64(pow))
		pow++
		fmt.Println(sum)
	}
	if(binary[len(binary) - 1] == 0){
		sum -= 1
	}
	fmt.Println(sum)
	return int(sum)
}