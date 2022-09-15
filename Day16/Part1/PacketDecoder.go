package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main(){
	line := readInput("../input.txt")[0]
	i, err := strconv.ParseUint(line, 16, 32)
	check(err)
	fmt.Println(asBits(i))
}

func asBits(val uint64) []uint64 {
	bits := []uint64{}
	for i:= 0; i<24; i++{
		bits = append([]uint64{val & 0x1}, bits...)
		val = val >> 1
	}
	return bits
}

func readInput(fileName string) []string {
	content, fileReadErr := os.ReadFile(fileName)
	check(fileReadErr)
	inputLines := strings.Split(string(content), "\n")
	return inputLines
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}