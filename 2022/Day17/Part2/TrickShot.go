package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"math"
)

func main(){
	lines := readInput("../input.txt")
	targetArea := getTargetArea(lines[0])
	fmt.Println("X: ",targetArea.X1, ", ", targetArea.X2, " Y: ", targetArea.Y1, ", ", targetArea.Y2)
	
	// maxX:= getMaxX(targetArea)
	maxY:= int(math.Abs(float64(targetArea.Y1)))
	counter:=0
	for y := targetArea.Y1; y <= maxY; y++{
		for x := targetArea.Y1; x <= maxY; x++{
			probe := &Probe{0, 0, x, y, -1, targetArea}
			shoot(probe)
			if probe.IsInTargetArea(){
				counter++
			}
		}
	}
	fmt.Println(counter)
}

type TargetArea struct{
	X1 int
	X2 int
	Y1 int
	Y2 int
}

type Probe struct{
	Xpos int
	Ypos int
	Xvel int
	Yvel int
	maxY int
	targetArea *TargetArea
}

func shoot(p *Probe) *Probe{
	step := 1
	startXVel := p.Xvel
	startYVel := p.Yvel
	for !p.IsInTargetArea() && p.IsMoving(){
		p.moveProbe()
		step++
	}
	if p.IsInTargetArea(){
		fmt.Println("DONE: ", step, " Max height: ", p.maxY, "Start X Vel: ", startXVel, "Start Y Vel: ", startYVel, " Made Target Area: ", p.IsInTargetArea())
	}
	return p
}

func (p *Probe) IsMoving() bool{
	return !(p.Xvel == 0 && p.Ypos < p.targetArea.Y1)
}

func (p *Probe) moveProbe() *Probe{
	p.Xpos = p.Xpos + p.Xvel
	p.Ypos = p.Ypos + p.Yvel
	if p.Ypos > p.maxY {
		p.maxY = p.Ypos
	}
	
	//X velocity
	if p.Xvel > 0 {
		p.Xvel -= 1
	}else if p.Xvel < 0{
		p.Xvel += 1
	}

	//Y velocity
	p.Yvel -= 1
	
	return p
}

func getMaxX(targetArea *TargetArea) int{
	root1 := 1 + math.Sqrt(1 - (4*(float64(-targetArea.X2*2))))/2
	return int(math.Round(root1))
}

func (p *Probe) IsInTargetArea() bool{
	if p.Xpos >= p.targetArea.X1 && p.Xpos <= p.targetArea.X2 && p.Ypos >= p.targetArea.Y1 && p.Ypos <= p.targetArea.Y2 {
		return true
	}
	return false
}

func getTargetArea(inputString string) *TargetArea{
	fields := strings.Fields(inputString)
	targetX := strings.Split(fields[2], "=")[1]
	targetY := strings.Split(fields[3], "=")[1]
	X1,err := strconv.Atoi(strings.Split(targetX, "..")[0])
	check(err)
	X2, err:= strconv.Atoi(strings.Trim(strings.Split(targetX, "..")[1], ","))
	check(err)
	Y1, err:= strconv.Atoi(strings.Split(targetY, "..")[0])
	check(err)
	Y2, err:= strconv.Atoi(strings.Split(targetY, "..")[1])
	check(err)
	targetArea := &TargetArea{X1, X2, Y1, Y2}
	return targetArea
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