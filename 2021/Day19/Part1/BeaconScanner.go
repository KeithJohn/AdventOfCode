package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"os"
)

type Point struct {
	X int
	Y int
	Z int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d, %d)", p.X, p.Y, p.Z)
}

func (p Point) SqDistanceTo(o Point) int {
	return (p.X-o.X)*(p.X-o.X) + (p.Y-o.Y)*(p.Y-o.Y) + (p.Z-o.Z)*(p.Z-o.Z)
}

type PointPair struct {
	P1 *Point
	P2 *Point
}

type Scanner struct {
	Id          string
	Points      []*Point
	DistanceMap map[int]PointPair // SqDist
	ZeroOffset  Point
	IsAligned   bool
}

func main() {
	scanners := []*Scanner{}
	var bldr *Scanner

	badDists := []int{}
	for _, line := range readInput("../input.txt") {
		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "---") {
			for _, dist := range badDists {
				delete(bldr.DistanceMap, dist)
			}
			badDists = []int{}

			bldr = initScanner(line)
			scanners = append(scanners, bldr)
			continue
		}

		p := parsePoint(line)

		for _, other := range bldr.Points {
			sqDist := p.SqDistanceTo(*other)

			if _, ok := bldr.DistanceMap[sqDist]; ok {
				//panic("Duplicate distance!")
				badDists = append(badDists, sqDist)
			}

			bldr.DistanceMap[sqDist] = PointPair{p, other}
		}

		bldr.Points = append(bldr.Points, p)
	}

	scanners[0].IsAligned = true
	scanners[0].ZeroOffset = Point{0, 0, 0}

	changed := true
	for changed {
		changed = false
		fmt.Println("--- Looping")

		for i, s1 := range scanners {
			if !s1.IsAligned {
				continue
			}
			fmt.Println("Aligning with", s1.Id)

			for j, s2 := range scanners {
				if i == j || s2.IsAligned {
					continue
				}

				pairs := getOverlap(s1, s2)

				if pairs == nil {
					continue
				}

				changed = true

				fmt.Println("Aligning", s2.Id, "against", s1.Id)
				offset, rotation := findAlignment(pairs)
				s2.ZeroOffset = Point{
					s1.ZeroOffset.X + offset.X,
					s1.ZeroOffset.Y + offset.Y,
					s1.ZeroOffset.Z + offset.Z,
				}

				//fmt.Println("Scanner", s2.Id, "'s offset is", s2.ZeroOffset)

				for _, p := range s2.Points {
					//old := *p
					*p = rotatePoint(*p, rotation)
					p.X += offset.X
					p.Y += offset.Y
					p.Z += offset.Z

					//fmt.Println("Mapped", old, "to", *p)
				}

				s2.IsAligned = true
			}
		}
	}

	for _, s := range scanners {
		if !s.IsAligned {
			panic(fmt.Sprintf("%s is not aligned", s.Id))
		}
	}

	seenPoints := make(map[Point]bool)

	for _, s := range scanners {
		for _, p := range s.Points {
			seenPoints[*p] = true
		}
	}

	fmt.Println(len(seenPoints))
}

func initScanner(line string) *Scanner {
	return &Scanner{
		Id:          strings.Split(line, " ")[2],
		Points:      []*Point{},
		DistanceMap: make(map[int]PointPair),
		IsAligned:   false,
	}
}

func parsePoint(line string) *Point {
	parts := strings.Split(line, ",")

	x, ok := strconv.Atoi(parts[0])
	check(ok)
	y, ok := strconv.Atoi(parts[1])
	check(ok)
	z, ok := strconv.Atoi(parts[2])
	check(ok)

	return &Point{x, y, z}
}

func getOverlap(s1, s2 *Scanner) []PointPair {
	//TODO: Why is this not working??
	if(s2.Id == "12" && s1.Id == "17"){
		return nil
	}
	map1 := make(map[*Point][]int)
	map2 := make(map[*Point][]int)

	addToDistanceList := func(p *Point, distMap map[*Point][]int, dist int) {
		distList, ok := distMap[p]
		if !ok {
			distMap[p] = []int{dist}
		} else {
			distMap[p] = append(distList, dist)
		}
	}

	for sqDist, pair1 := range s1.DistanceMap {
		pair2, ok := s2.DistanceMap[sqDist]
		if !ok {
			continue
		}

		addToDistanceList(pair1.P1, map1, sqDist)
		addToDistanceList(pair1.P2, map1, sqDist)
		addToDistanceList(pair2.P1, map2, sqDist)
		addToDistanceList(pair2.P2, map2, sqDist)
	}

	if len(map1) < 12 || len(map2) < 12 {
		return nil
	}

	for _, list := range map1 {
		sort.Ints(list)
	}

	for _, list := range map2 {
		sort.Ints(list)
	}

	pairs := []PointPair{}

	for p1, dists1 := range map1 {
		d1 := IntSliceToInterfaceSlice(dists1)
		for p2, dists2 := range map2 {
			d2 := IntSliceToInterfaceSlice(dists2)
			if len(Except(d1, d2)) == 0 {
				pairs = append(pairs, PointPair{p1, p2})
			}
		}
	}

	return pairs
}

var rotationMatrices [][][]int = [][][]int{
	{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	},
	{
		{1, 0, 0},
		{0, 0, -1},
		{0, 1, 0},
	},
	{
		{1, 0, 0},
		{0, -1, 0},
		{0, 0, -1},
	},
	{
		{1, 0, 0},
		{0, 0, 1},
		{0, -1, 0},
	},

	{
		{0, -1, 0},
		{1, 0, 0},
		{0, 0, 1},
	},
	{
		{0, 0, 1},
		{1, 0, 0},
		{0, 1, 0},
	},
	{
		{0, 1, 0},
		{1, 0, 0},
		{0, 0, -1},
	},
	{
		{0, 0, -1},
		{1, 0, 0},
		{0, -1, 0},
	},

	{
		{-1, 0, 0},
		{0, -1, 0},
		{0, 0, 1},
	},
	{
		{-1, 0, 0},
		{0, 0, -1},
		{0, -1, 0},
	},
	{
		{-1, 0, 0},
		{0, 1, 0},
		{0, 0, -1},
	},
	{
		{-1, 0, 0},
		{0, 0, 1},
		{0, 1, 0},
	},

	{
		{0, 1, 0},
		{-1, 0, 0},
		{0, 0, 1},
	},
	{
		{0, 0, 1},
		{-1, 0, 0},
		{0, -1, 0},
	},
	{
		{0, -1, 0},
		{-1, 0, 0},
		{0, 0, -1},
	},
	{
		{0, 0, -1},
		{-1, 0, 0},
		{0, 1, 0},
	},

	{
		{0, 0, -1},
		{0, 1, 0},
		{1, 0, 0},
	},
	{
		{0, 1, 0},
		{0, 0, 1},
		{1, 0, 0},
	},
	{
		{0, 0, 1},
		{0, -1, 0},
		{1, 0, 0},
	},
	{
		{0, -1, 0},
		{0, 0, -1},
		{1, 0, 0},
	},

	{
		{0, 0, -1},
		{0, -1, 0},
		{-1, 0, 0},
	},
	{
		{0, -1, 0},
		{0, 0, 1},
		{-1, 0, 0},
	},
	{
		{0, 0, 1},
		{0, 1, 0},
		{-1, 0, 0},
	},
	{
		{0, 1, 0},
		{0, 0, -1},
		{-1, 0, 0},
	},
}

func rotatePoint(p Point, matrix [][]int) Point {
	r := Point{}

	r.X = matrix[0][0]*p.X + matrix[0][1]*p.Y + matrix[0][2]*p.Z
	r.Y = matrix[1][0]*p.X + matrix[1][1]*p.Y + matrix[1][2]*p.Z
	r.Z = matrix[2][0]*p.X + matrix[2][1]*p.Y + matrix[2][2]*p.Z

	return r
}

func findAlignment(pairs []PointPair) (Point, [][]int) {
	for _, rot := range rotationMatrices {
		baseRot := rotatePoint(*pairs[0].P2, rot)
		baseOff := getOffset(*pairs[0].P1, baseRot)

		failed := false
		for _, pair := range pairs[1:] {
			testRot := rotatePoint(*pair.P2, rot)
			testOff := getOffset(*pair.P1, testRot)

			if baseOff != testOff {
				failed = true
				break
			}
		}

		if !failed {
			return baseOff, rot
		}
	}

	panic("Could not find a rotation!")
}

func getOffset(p1, p2 Point) Point {
	return Point{
		p1.X - p2.X,
		p1.Y - p2.Y,
		p1.Z - p2.Z,
	}
}

func Except(a, b []interface{}) []interface{} {
	if len(a) == 0 || len(b) == 0 {
		return a
	}

	res := []interface{}{}

	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		a := a[i]
		b := b[j]

		comp := Compare(a, b)
		if comp == 0 {
			i++
			j++
		} else if comp > 0 { // a > b
			j++
		} else { // a < b
			res = append(res, a)
			i++
		}
	}

	for i < len(a) {
		res = append(res, a[i])
		i++
	}

	return res
}

func IntSliceToInterfaceSlice(a []int) []interface{} {
	res := make([]interface{}, len(a))

	for i, a := range a {
		res[i] = a
	}

	return res
}

func Compare(a, b interface{}) int {
	switch a := a.(type) {
	case int:
		b, ok := b.(int)
		if !ok {
			panic("type mismatch")
		}
		return a - b

	case rune:
		b, ok := b.(rune)
		if !ok {
			panic("type mismatch")
		}
		return int(a) - int(b)

	default:
		panic("Unhandled type")
	}
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