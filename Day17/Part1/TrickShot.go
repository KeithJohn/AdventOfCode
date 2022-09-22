package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main(){
	lines := readInput("../simpleInput.txt")
	targetArea := getTargetArea(lines[0])
	fmt.Println("X: ",targetArea.X1, ", ", targetArea.X2, " Y: ", targetArea.Y1, ", ", targetArea.Y2)
	
	probe := &Probe{0, 0, 6, 9, -1, targetArea}
	shoot(probe)

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
	for !p.IsInTargetArea() && p.IsMoving(){
		p.moveProbe()
		step++
	}
	fmt.Println("DONE: ", step, " Max height: ", p.maxY, " Made Target Area: ", p.IsInTargetArea())
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
	
	fmt.Println("X: ", p.Xpos, " Y: ", p.Ypos, " XVel: ", p.Xvel, " YVel: ", p.Yvel, "maxY: ", p.maxY)
	return p
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