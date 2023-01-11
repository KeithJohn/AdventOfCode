package main

import (
	"fmt"
	"os"
	"strings"
)

type SnailfishNumber struct{
	left *SnailfishNumber
	right *SnailfishNumber
	parent *SnailfishNumber

	value int
}

func (s SnailfishNumber) String() string{
	if s.IsLeaf() {
		return fmt.Sprintf("%d", s.value)
	}

	return fmt.Sprintf("[%s, %s]", s.left.String(), s.right.String())
}

func (s SnailfishNumber) IsLeaf() bool{
	if s.left == nil{
		return true;
	}
	return false
}

func (s *SnailfishNumber) IsLeftChild() bool{
	if s.parent != nil && s.parent.left == s{
		return true
	}
	return false
}

func (s *SnailfishNumber) IsRightChild() bool{
	if s.parent != nil && s.parent.right == s{
		return true
	}
	return false
}

func (s *SnailfishNumber) IsRoot() bool{
	if s.parent == nil {
		return true
	}
	return false
}

func (s *SnailfishNumber) Magnitude() int{
	if s.IsLeaf(){
		return s.value
	}

	return (3 * s.left.Magnitude()) + (2 * s.right.Magnitude())
}

func (s *SnailfishNumber) FindExplodeCandidate(depth int) *SnailfishNumber {
	if s.IsLeaf() {
		return nil
	}

	if s.left.IsLeaf() && s.right.IsLeaf() && depth >= 4{
		return s
	}

	left := s.left.FindExplodeCandidate(depth+1)
	if left != nil {
		return left
	}
	right := s.right.FindExplodeCandidate(depth+1)
	if right != nil {
		return right
	}

	return nil
}

func (s *SnailfishNumber) Explode() {
	//TODO:
	s.value = 0
	curr := s.left
	for curr.IsLeftChild() {
		curr = curr.parent
	}

	if !curr.IsRoot() {
		//TODO:
		curr = curr.parent.left
		for !curr.IsLeaf() {
			curr = curr.right
		}
		curr.value += s.left.value
	}

	s.left = nil

	curr = s.right
	for curr.IsRightChild(){
		curr = curr.parent
	}

	if !curr.IsRoot() {
		curr = curr.parent.right
		for !curr.IsLeaf() {
			curr = curr.left
		}
		curr.value += s.right.value
	}
	s.right = nil
}

func (s *SnailfishNumber) FindSplitCandidate() *SnailfishNumber{
	//TODO:
	if s.IsLeaf() {
		if s.value >= 10{
			return s
		}
		return nil
	}
	

	left := s.left.FindSplitCandidate()
	if left != nil {
		return left
	}

	right := s.right.FindSplitCandidate()
	if right != nil{
		return right
	}
	return nil
}

func (s *SnailfishNumber) Split(){
	leftValue := s.value/2
	rightValue := s.value - leftValue
	left := &SnailfishNumber{nil, nil, s, leftValue}
	s.left = left

	right := &SnailfishNumber{nil, nil, s, rightValue}
	s.right = right

	s.value = 0
}

func (s *SnailfishNumber) AddAndReduce(addend *SnailfishNumber) *SnailfishNumber{
	temp := &SnailfishNumber{s, addend, nil, 0}
	s.parent = temp
	addend.parent = temp
	for {
		explodeCandidate := temp.FindExplodeCandidate(0)
		if explodeCandidate != nil{
			explodeCandidate.Explode()
			continue
		}

		splitCandidate := temp.FindSplitCandidate()
		if splitCandidate != nil {
			splitCandidate.Split()
			continue
		}

		break
	}
	return temp
}  

type Scanner struct{
	text string
}

func (s Scanner) Peek() rune{
	return rune(s.text[0])
}

func (s *Scanner) Next() rune {
	next := rune(s.text[0])
	s.text = s.text[1:]
	return next
}

func (s *Scanner) AssertNext(r rune) {
	if s.Peek() != r{
		//Error
		panic(fmt.Sprintf("Error: Next character should be %s but was %s instead.", r, s.Peek()))
	}
	s.Next()
}

func (s *Scanner) NextInt() int{
	val := 0

	for s.Peek() >= '0' && s.Peek() <= '9' {
		r := s.Next()
		val = val*10 + int(r-'0')
	}

	return val
}

func main(){
	lines := readInput("../input.txt")
	numbers := []*SnailfishNumber{}
	for _, line := range lines{
		scanner := &Scanner{line}
		numbers = append(numbers, parseNumber(scanner))
	}
	sum := numbers[0]
	for i:=1; i < len(numbers); i++{
		sum = sum.AddAndReduce(numbers[i])
	}
	fmt.Println("SUM: ", sum.String(), " - Magnitude: ", sum.Magnitude())

}

func parseNumber(s *Scanner) *SnailfishNumber{
	snailfishNumber := &SnailfishNumber{}
	if s.Peek() == '[' {
		s.Next()

		snailfishNumber.left = parseNumber(s)
		snailfishNumber.left.parent = snailfishNumber
		
		s.AssertNext(',')

		snailfishNumber.right = parseNumber(s)
		snailfishNumber.right.parent = snailfishNumber

		s.AssertNext(']')
	}else{
		snailfishNumber.value = s.NextInt()
	}

	return snailfishNumber
}

func readInput(fileName string) []string {
	content, fileReadErr := os.ReadFile(fileName)
	check(fileReadErr)
	inputLines := strings.Split(string(content), "\n")
	return inputLines
} 

func check(err error){
	if err != nil {
		panic(err)
	}
}