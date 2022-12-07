package main

import (
	"adventofcode/util"
	"fmt"
	"strconv"
	"strings"
)

type Instruction struct {
	amount       int
	fromStackNum int
	toStackNum   int
}

func main() {
	lines := util.ReadInput("../input.txt")
	stacks := getStartingStacks(lines)
	instructions := getInstructions(lines)
	for _, inst := range instructions {
		runInstructions(stacks, inst)
	}

	fmt.Print("Solution: ")
	for _, stack := range stacks {
		_, val := stack.Pop()
		fmt.Print(val)
	}
}

func getStartingStacks(lines []string) []util.Stack {
	stacks := make([]util.Stack, 0)
	numOfStacks := len(lines[0]) / 4
	for j := 0; j < numOfStacks; j++ {
		stacks = append(stacks, make(util.Stack, 0))
	}

	for _, line := range lines {
		if len(line) == 1 {
			break
		}

		for j := 0; j < numOfStacks; j++ {
			box := line[4*j : 4*(j+1)]
			if strings.HasPrefix(box, "[") {
				stacks[j] = append([]string{string(box[1])}, stacks[j]...)
			}
		}
	}
	return stacks
}

func getInstructions(lines []string) []Instruction {
	instructions := make([]Instruction, 0)
	for _, line := range lines {
		if strings.HasPrefix(line, "move") {
			instructionParts := strings.Split(line, " ")
			instruction := new(Instruction)
			amount, err := strconv.Atoi(instructionParts[1])
			util.Check(err)
			from, err := strconv.Atoi(instructionParts[3])
			util.Check(err)
			to, err := strconv.Atoi(strings.ReplaceAll(instructionParts[5], "\r", ""))
			util.Check(err)
			instruction.amount = amount
			instruction.fromStackNum = from
			instruction.toStackNum = to
			instructions = append(instructions, *instruction)
		}
	}
	return instructions
}

func runInstructions(stacks []util.Stack, instruction Instruction) []util.Stack {
	val := ""
	for i := 1; i <= instruction.amount; i++ {
		stacks[instruction.fromStackNum-1], val = stacks[instruction.fromStackNum-1].Pop()
		stacks[instruction.toStackNum-1] = stacks[instruction.toStackNum-1].Push(val)
	}
	return stacks
}
