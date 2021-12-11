package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	numbersMap  map[int]int
	markedBoard [5][5]int
}

type Input struct {
	numbers []int
	boards  []Board
}

func main() {
	fmt.Println("hello")
	input := readInputs("../input.txt")
	solution := 0
	for _, number := range input.numbers {
		fmt.Println(number)
		for _, board := range input.boards {
			index := board.numbersMap[number]
			if index != 0 {
				fmt.Println(index)
				fmt.Println("Board contains number ", board.numbersMap)
				//Board contains number
				column := index % 5
				row := 0
				if index < 5 {
					row = 0
				} else if index < 10 {
					row = 1
				} else if index < 15 {
					row = 2
				} else if index < 20 {
					row = 3
				} else {
					row = 4
				}
				fmt.Println("Row Column:", row, column)
				board.markedBoard[row][column] = 1
				board.numbersMap[number] = 0
				fmt.Println("Before Check", board.markedBoard)
				if checkBoard(board, row, column) {
					//Board is completed. Calculate sum
					solution = calculateSum(board, number)
					break
				}

			}
		}
		if solution != 0 {
			break
		}
	}
	fmt.Println(solution)
}

func readInputs(fileName string) Input {
	var input Input
	content, fileReadErr := os.ReadFile(fileName)
	check(fileReadErr)
	lines := strings.Split(string(content), "\n")

	//First line will be number inputs
	numLine := lines[0]
	var numbers []int
	for _, val := range strings.Split(numLine, ",") {
		number, convErr := strconv.Atoi(strings.TrimSpace(val))
		check(convErr)
		numbers = append(numbers, number)
	}
	input.numbers = numbers
	//new line then five lines for board. Repeat until no more lines
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		if line != "" {
			var boardLines [][]string
			for j := i; j < i+5; j++ {
				boardLine := strings.Split(lines[j], " ")
				r := make([]string, 0, 5)
				for _, str := range boardLine {
					if str != "" {
						r = append(r, strings.TrimSpace(str))
					}
				}
				boardLines = append(boardLines, r)
			}
			i += 5
			board := newBoard(boardLines)
			input.boards = append(input.boards, board)
		}
	}
	return input
}

func newBoard(rows [][]string) Board {
	numMap := make(map[int]int)
	var markedBoard [5][5]int
	for i, row := range rows {
		for j, value := range row {
			//5* i + j is index
			index := (5 * i) + j
			intValue, err := strconv.Atoi(value)
			check(err)
			numMap[intValue] = index
		}
	}
	var newBoard Board
	newBoard.markedBoard = markedBoard
	newBoard.numbersMap = numMap
	return newBoard
}

func checkBoard(board Board, row int, column int) bool {
	fmt.Println("Checking Board")
	markedBoard := board.markedBoard
	fmt.Println(markedBoard)
	//Check row
	rowComplete := true
	for i := 0; i < 5; i++ {
		if markedBoard[row][i] != 1 {
			rowComplete = false
		}
	}
	//Check column
	columnComplete := true
	for i := 0; i < 5; i++ {
		if markedBoard[i][column] != 1 {
			columnComplete = false
		}
	}

	return rowComplete || columnComplete
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func calculateSum(board Board, number int) int {
	sum := 0
	for _, value := range board.numbersMap {
		sum += value
	}
	sum *= number
	return sum
}

//Each board consists of hashmap and 2d array
// Hashmap contains the numbers as key and position as value
// column is value % 5
// row 1 is if value < 5
// row 2 is if value < 10
// row 3 is if value < 15
// row 4 is if value < 20
// row 5 is else

//
// 0  1  2  3  4
// 5  6  7  8  9
// 10 11 12 13 14
// 15 16 17 18 19
// 20 21 22 23 24

//When checking the board, only check the rows affected by the new number. No need to look at rows/columns that did not change

//get list of inputs
//get and populate boards
//loop through list of inputs
//for each board check if number is in hashmap
// if yes mark the number off in the 2d array and remove from hashmap
// check the column and row for the number that was marked off
// if any are bingo calculate score
// to calculate score sum the values in hashmap
// multiply by last number
