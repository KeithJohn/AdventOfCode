package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

type Step struct {
	X    int
	Y    int
	Cost int
}

type Interface interface {
	sort.Interface
	Push(x interface{})
	Pop() interface{}
}

type StepPriorityQueue []Step

func (h StepPriorityQueue) Len() int           { return len(h) }
func (h StepPriorityQueue) Less(i, j int) bool { return h[i].Cost < h[j].Cost }
func (h StepPriorityQueue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *StepPriorityQueue) Push(x interface{}) {
	*h = append(*h, x.(Step))
}

func (h *StepPriorityQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type Grid struct {
	Values [][]int
	Loops int
	Width int
	Height int
}

func newGrid (values [][]int, loops int) Grid {
	return Grid{
		values,
		loops,
		len(values[0]),
		len(values),
	}
}

func (g Grid) isInBounds(x, y int) bool{
	xLoop := x / g.Width
	yLoop := y / g.Height

	return x >= 0 && y >= 0 && xLoop < g.Loops && yLoop < g.Loops
}

func (g Grid) get(x, y int) int {
	xLoop := x / g.Width
	xIdx := x % g.Width

	yLoop := y / g.Height
	yIdx := y % g.Height

	baseValue := g.Values[yIdx][xIdx]

	loop := xLoop + yLoop
	value := baseValue + loop

	if value > 9 {
		value -= 9
	}

	return value
}

func main() {
	values := getGraph()
	grid := newGrid(values, 5)
	
	costs := makeCostGrid(grid.Width, grid.Height, grid.Loops)

	targetY := grid.Height*grid.Loops - 1
	targetX :=  grid.Width*grid.Loops - 1

	queue := &StepPriorityQueue{}
	queue.Push(Step{X: 0, Y: 0, Cost: 0})

	heap.Init(queue)
	for queue.Len() > 0 {
		curr := heap.Pop(queue).(Step)

		x := curr.X
		y := curr.Y


		if x == targetX && y == targetY{
			fmt.Println(curr.Cost)
			return
		}

		// Already have a cheaper path
		if costs[y][x] <= curr.Cost {
			continue
		}

		//fmt.Printf("At (%d,%d) with cost %d\n", curr.X, curr.Y, curr.Cost)

		costs[y][x] = curr.Cost

		for i := -1; i <= 1; i++ {

			if i == 0 {
				continue
			}

			if grid.isInBounds(x+i, y){
				step := Step{x + i, y, curr.Cost + grid.get(x+i, y)}
				heap.Push(queue, step)
			}

			if grid.isInBounds(x, y+i){
				step := Step{x, y + i, curr.Cost + grid.get(x, y + i)}
				heap.Push(queue, step)
			}
		}
	}
}

func makeCostGrid(x, y int, loops int) [][]int {
	grid := make([][]int, y*loops)

	for i := 0; i < y*loops; i++ {
		grid[i] = make([]int, x*loops)

		for j := 0; j < x*loops; j++ {
			grid[i][j] = math.MaxInt
		}
	}
	return grid
}

func getGraph() [][]int {
	inputLines := readInput("../input.txt")
	grid := [][]int {}

	for _, line := range inputLines{
		row := []int{}
		for _, cell := range  line{
			row = append(row, int(cell-'0'))
		}

		grid = append(grid, row)
	}

	return grid
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
