package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
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

func main() {
	graph := getGraph()
	costs := makeCostGrid(len(graph[0]), len(graph))

	targetY := len(graph) - 1
	targetX := len(graph[targetY]) - 1

	queue := &StepPriorityQueue{}
	queue.Push(Step{X: 0, Y: 0, Cost: 0})

	heap.Init(queue)
	for queue.Len() > 0 {
		curr := heap.Pop(queue).(Step)

		// Already have a cheaper path
		if costs[curr.Y][curr.X] <= curr.Cost {
			continue
		}

		//fmt.Printf("At (%d,%d) with cost %d\n", curr.X, curr.Y, curr.Cost)

		costs[curr.Y][curr.X] = curr.Cost

		if curr.X == targetX && curr.Y == targetY {
			fmt.Println(curr.Cost)
			return
		}

		for i := -1; i <= 1; i++ {
			if curr.X+i >= 0 && curr.X+i < len(graph[curr.Y]) && i != 0 {
				step := Step{curr.X + i, curr.Y, curr.Cost + graph[curr.Y][curr.X+i]}
				heap.Push(queue, step)
			}

			if curr.Y+i >= 0 && curr.Y+i < len(graph) && i != 0 {
				step := Step{curr.X, curr.Y + i, curr.Cost + graph[curr.Y+i][curr.X]}
				heap.Push(queue, step)
			}
		}
	}
}

func makeCostGrid(x, y int) [][]int {
	grid := make([][]int, y)

	for i := 0; i < y; i++ {
		grid[i] = make([]int, x)

		for j := 0; j < x; j++ {
			grid[i][j] = math.MaxInt
		}
	}
	return grid
}

func getGraph() [][]int {
	inputLines := readInput("../input.txt")
	rows := len(inputLines)
	columns := len(inputLines[0]) - 1
	//Populate Grid
	graph := make([][]int, rows)
	for i, inputLine := range inputLines {
		graph[i] = make([]int, columns)
		vals := strings.Split(strings.TrimSuffix(inputLine, "\r"), "")
		for j, inputNum := range vals {
			num, convErr := strconv.Atoi(inputNum)
			check(convErr)
			graph[i][j] = num
		}
	}

	return graph
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
