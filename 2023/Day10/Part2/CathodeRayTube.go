package main

import (
	"adventofcode/util"
	"fmt"
	"strings"
)

type CPU struct {
	regX         int
	currCycle    int
	currStrength int
}

func main() {
	fmt.Println()
	lines := util.ReadInput("../input.txt")
	cpu := CPU{1, 1, 0}
	for _, line := range lines {
		commandParts := strings.Split(line, " ")
		if commandParts[0] == "addx" {
			cpu.addX(util.ConvertAtoi(commandParts[1]))
		} else {
			cpu.noop()
		}
	}
	fmt.Println()
}

func (c *CPU) addX(num int) {
	c.cycle()
	c.cycle()
	c.regX += num
}

func (c *CPU) noop() {
	c.cycle()
}

func (c *CPU) cycle() {
	if c.currCycle%40 == c.regX || c.currCycle%40 == c.regX+2 || c.currCycle%40 == c.regX+1 || (c.currCycle%40 == 0 && c.regX+2 == 40) {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	if c.currCycle != 0 && c.currCycle%40 == 0 {
		fmt.Println()
	}
	c.currCycle += 1
}
