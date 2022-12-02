package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"strconv"
)

func main(){
	file, err := os.Open("input.txt")
	check(err)
    defer file.Close()

	horizontalLoc := 0
	verticlLoc := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		command := strings.Fields(scanner.Text())
		direction := command[0]
		distance, convErr := strconv.Atoi(command[1])
		check(convErr)
		switch direction{
		case "up":
			verticlLoc -= distance
			fmt.Println("Up", distance)
		case "down":
			verticlLoc += distance
			fmt.Println("Down", distance)
		case "forward":
			horizontalLoc += distance
			fmt.Println("Forward")
		}
	}
	fmt.Println("Horizontal", horizontalLoc)
	fmt.Println("Vertical", verticlLoc)

	fmt.Println("Solution", horizontalLoc * verticlLoc)
}

func check(e error){
    if e != nil{
        panic(e)
    }
}