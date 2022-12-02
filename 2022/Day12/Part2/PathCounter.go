package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	nodes := make(map[string]*Node)

	for _, line := range readInput("../input.txt") {
		caves := strings.Split(line, "-")
		caves0 := strings.TrimSpace(caves[0])
		caves1 := strings.TrimSpace(caves[1])
		cave1, ok := nodes[caves0]
		if !ok {
			cave1 = newNode(caves0)
			nodes[caves0] = cave1
		}

		cave2, ok := nodes[caves1]
		if !ok {
			cave2 = newNode(caves1)
			nodes[caves1] = cave2
		}

		cave1.AddEdge(cave2)
		cave2.AddEdge(cave1)
	}
	start := nodes["start"]
	end := nodes["end"]

	paths := start.CountPaths(end, false)

	fmt.Println(paths)
}

type Node struct {
	Name     string
	Edges    []*Node
	IsMarked bool
	IsLarge  bool
}

func newNode(name string) *Node {
	isLarge := unicode.IsUpper(rune(name[0]))

	return &Node{
		Name:     name,
		Edges:    []*Node{},
		IsLarge:  isLarge,
		IsMarked: false,
	}
}

func (n *Node) String() string {
	return n.Name
}

func (n *Node) AddEdge(neighbor *Node) {
	n.Edges = append(n.Edges, neighbor)
}

func (n *Node) CountPaths(end *Node, hasUsedExtraStop bool) int {
	if n == end {
		return 1
	}

	if n.IsMarked {
		if hasUsedExtraStop || n.Name == "start" {
			return 0
		} else {
			hasUsedExtraStop = true
		}
	}

	if !n.IsLarge && !n.IsMarked {
		n.IsMarked = true
		defer func() { n.IsMarked = false }()
	}

	sum := 0
	for _, adj := range n.Edges {
		sum += adj.CountPaths(end, hasUsedExtraStop)
	}

	return sum
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
