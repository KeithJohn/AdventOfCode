package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	calls := []int{}

	var boardBuild *Board

	cellRefs := make(map[int][]BoardCellRef)

	for i, line := range readInputLines("../input.txt") {
		if i == 0 {
			for _, call := range strings.Split(line, ",") {
				call, convErr := strconv.Atoi(call)
				check(convErr)
				calls = append(calls, call)
			}
		}else  if (len(line) == 0){
			if(boardBuild != nil){
				fmt.Println(boardBuild.String())
				fmt.Println()
			}

			boardBuild = &Board{RemainingScore: 0, Cells: [][]int{}}
		}else{
			cells := []int{}

			y:= len(boardBuilds.Cells)

			for x, value := range strings.FieldsFunc(line, func(c rune) bool {return c == ' '}){
				value, err = strconv.Atoi(strings.TrimSpace(value))
				check(err)

				cells = append(cells, value)
				boardBuild.RemainingScore += value
				cellRef := BoardCellRef{X:x, Y:y, Board:boardBuild}
				refList, ok := cellRefs[value]
				if !ok {
					refList = []BoardCellRef{}
				}

				cellRefs[value] = append(refList, cellRef)
			}
			boardBuild.Cells = append(boardBuild.Cells, cells)
		}
	}
	if boardBuild != nil{
		fmt.Println(boardBuild.String())
		fmt.Println()
	}

	for _, call := range calls{
		refs, ok := cellRefs[call]
		if !ok{
			continue //No issue, number not on a board
		}

		for _, 
	}
}

type BoardCellRef struct {
	X     int
	Y     int
	Board *Board
}

type Board struct {
	RemainingScore int
	Cells          [][]int
}

//Returns true if board has a complete row or column
func (b *Board) call(x, y int) bool {
	b.RemainingScore -= b.Cells[y][x]
	b.Cells[y][x] = -1
	xWin := true
	yWin := true

	for i := 0; i < 5; i++ {
		xWin = xWin && (b.Cells[y][i] == -1)
		yWin = yWin && (b.Cells[i][x] == -1)
	}

	return xWin || yWin
}

func (b *Board) String() string {
	lines := make([]string, len(b.Cells))

	for i, line := range b.Cells {
		cellStrs := make([]string, len(lines))

		for j, cell := range line {
			cellStrs[j] = fmt.Sprintf("%2d", cell)
		}

		lines[i] = strings.Join(cellStrs, " ")
	}

	return strings.Join(lines, "\n")
}

func readInputLines(fileName string) []string {
	return nil
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
