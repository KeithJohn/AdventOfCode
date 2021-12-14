package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type fish struct {
	spawnTimer int
	isNewToday bool
}

func main() {
	fishList := readInputs("../input.txt")
	fmt.Println("Fish List: ", fishList)
	numberOfDays := 80
	for i := 1; i <= numberOfDays; i++ {
		numOfFishBorn := 0
		for fishIndex, fish := range fishList {
			fish.spawnTimer -= 1
			if fish.spawnTimer == -1 {
				//Fish is born!
				numOfFishBorn += 1
				fish.spawnTimer = 6
			}
			fishList[fishIndex] = fish
		}
		fishList = birthFish(fishList, numOfFishBorn)
		fmt.Println("Fish Count after ", i, " days:", len(fishList))
	}
	fmt.Println("Fish Count: ", len(fishList))
}

func birthFish(fishList []fish, numberOfFishToBeBorn int) []fish {
	for i := 0; i < numberOfFishToBeBorn; i++ {
		var fish fish
		fish.spawnTimer = 8
		fishList = append(fishList, fish)
	}
	return fishList
}

func readInputs(fileName string) []fish {
	content, fileReadErr := os.ReadFile(fileName)
	check(fileReadErr)

	nums := strings.Split(string(content), ",")
	var fishList []fish
	for _, num := range nums {
		var fish fish
		spawnTimer, convErr := strconv.Atoi(num)
		check(convErr)
		fish.spawnTimer = spawnTimer
		fishList = append(fishList, fish)
	}
	fmt.Println("Number of Fish: ", len(fishList))
	return fishList
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
