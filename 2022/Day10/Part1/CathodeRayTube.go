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
	fmt.Printf("CPU Strength: %d\n", cpu.currStrength)
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
	if c.currCycle%40 == 20 {
		strength := c.currCycle * c.regX
		fmt.Printf("CPU %dth cycle: %d * %d regex = %d\n", c.currCycle, c.currCycle, c.regX, strength)
		c.currStrength += strength
	}
	c.currCycle += 1
}
