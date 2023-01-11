package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
)

func main(){
	//Read values from input
	inputValues := readInput("../input.txt")

	oxygenGeneratorRatingBin := getOxygenGeneratorRating(inputValues)
	oxygenGeneratorRating, err := strconv.ParseInt(oxygenGeneratorRatingBin, 2, 64)
	check(err)
	fmt.Println(oxygenGeneratorRatingBin)
	fmt.Println(oxygenGeneratorRating)

	//fmt.Println(binaryCalculator(getOxygenGeneratorRating(inputValues)))
	fmt.Println(getC02ScrubberRaing(inputValues))
	co2ScrubberRatingBin := getC02ScrubberRaing(inputValues)
	co2ScrubberRating, err := strconv.ParseInt(co2ScrubberRatingBin, 2, 64)
	check(err)
	fmt.Println(co2ScrubberRatingBin)
	fmt.Println(co2ScrubberRating)

	solution := oxygenGeneratorRating * co2ScrubberRating
	fmt.Println("SOLUTION: ", solution)
}

func getOxygenGeneratorRating(inputValues []string) string{
	charCount := len(inputValues[0])
	gammaFilteredInputs := inputValues
	for i:= 0; i < charCount; i++{
		gammaVal := gammaCalculator(gammaFilteredInputs, i)

		gammaFilteredInputs = filterByCriteria(gammaFilteredInputs, gammaVal, i)

		if(len(gammaFilteredInputs) == 1){
			return gammaFilteredInputs[0]
		}
	}
	return "FAIL"
}

func getC02ScrubberRaing(inputValues []string) string{
	charCount := len(inputValues[0])
	epsilonFilteredInputs := inputValues
	for i:=0; i< charCount; i++{
		epsilonVal := epsilonCalculator(epsilonFilteredInputs, i)

		epsilonFilteredInputs = filterByCriteria(epsilonFilteredInputs, epsilonVal, i)
		if(len(epsilonFilteredInputs) == 1){
			return epsilonFilteredInputs[0]
		}
	}
	return "FAIL"
}

func readInput(fileName string) []string {
	content, readFileErr := os.ReadFile(fileName)
	check(readFileErr)
	return strings.Split(string(content), "\n")
}

func gammaCalculator(values []string, index int) string{
	oneCount := 0
	//Check index of each line in values
	for _, line := range values{
		if line[index] == '1'{
			oneCount++
		}
	}
	if(float64(oneCount) >= float64(len(values)) / 2){
		return "1"
	}else{
		return "0"
	}
}

func epsilonCalculator(values []string, index int) string{
	oneCount := 0
	//Check index of each line in values
	for _, line := range values{
		if line[index] == '1'{
			oneCount++
		}
	}
	if(float64(oneCount) >= float64(len(values)) / 2){
		return "0"
	}else{
		return "1"
	}
}

func filterByCriteria(values []string, critValue string, index int) []string{
	filteredValues := make([]string, 0, len(values))
	for _, value := range values{
		if string(value[index]) == critValue{
			filteredValues = append(filteredValues, value)
		}
	}
	fmt.Println("filtered values: ", len(filteredValues))
	return filteredValues
}

func binaryCalculator(binary string) int{
	sum := float64(0)
	pow := 0
	for i:= len(binary) - 1; i >= 0; i--{
		value := int(binary[i])
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

func check(err error){
	if(err != nil){
		panic(err)
	}
}