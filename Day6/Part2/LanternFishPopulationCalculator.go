package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//TODO: No need to track each individual fish, just track the number of fishes with specific spawn timer. For example, keep array with indexes 0 - 8 if spawn timer is 0 increment array[0] 1
func main() {
	fishList := readInputs("../input.txt")
	fmt.Println("Fish List: ", fishList)
	numberOfDays := 256
	for i := 1; i <= numberOfDays; i++ {
		sum := 0
		for _, count := range fishList {
			sum += count
		}
		fmt.Println("Day ", i, " fish total: ", sum)
		new_fishes := make([]int, 9)
		for index, count := range fishList {
			if index == 0 {
				new_fishes[6] += count
				new_fishes[8] += count
			} else {
				new_fishes[index-1] += count
			}
		}
		fishList = new_fishes
	}
	sum := 0
	for _, count := range fishList {
		sum += count
	}
	fmt.Println("Fish Count: ", sum)
}

func readInputs(fileName string) []int {
	content, fileReadErr := os.ReadFile(fileName)
	check(fileReadErr)

	nums := strings.Split(string(content), ",")
	fishList := make([]int, 9)
	for _, num := range nums {
		spawnTimer, convErr := strconv.Atoi(num)
		check(convErr)
		fishList[spawnTimer] += 1
	}
	fmt.Println("Number of Fish Spawn Timers: ", fishList)
	return fishList
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
