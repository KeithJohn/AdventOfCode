package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main(){

	lines := readInput("../input.txt")
	
	algorithm := lines[0]
	fmt.Println(algorithm)
	fmt.Println()

	grid := parseGrid(lines[2:])
	printGrid(grid)

	enhancedImage, litCount := enhanceImage(grid, algorithm)
	fmt.Println(litCount)
	printGrid(enhancedImage)
	fmt.Println(len(enhancedImage))

	// enhancedImage2, litCount2 := enhanceImage(enhancedImage, algorithm)
	// fmt.Println(litCount2)
	// printGrid(enhancedImage2)


	// finalLitCount := 0
	// for _, row := range enhancedImage2 {
	// 	for _, entry := range row{
	// 		if entry == "#" {
	// 			finalLitCount++
	// 		}
	// 	}
	// }
	// fmt.Println(finalLitCount)
	// fmt.Println(len(enhancedImage2))
	//printGrid(enhancedImage)

	// enhancedPixel := enhancePixel(2, 2, grid, algorithm)
	// fmt.Println(enhancedPixel)

	// newImage := make([][]string, 3)
	
	// row1 := make([]string, 3)
	// row1[0] = "#"
	// row1[1] = "."
	// row1[2] = "."
	// row2 := make([]string, 3)
	// row2[0] = "."
	// row2[1] = "#"
	// row2[2] = "."
	// row3 := make([]string, 3)
	// row3[0] = "."
	// row3[1] = "."
	// row3[2] = "#"

	// newImage[0] = row1
	// newImage[1] = row2
	// newImage[2] = row3

	// printGrid(newImage)

	// newImage2 := expandImage(newImage)
	// fmt.Println()
	// printGrid(newImage2)

	// newImage3 := expandImage(newImage2)
	// fmt.Println()
	// printGrid(newImage3)
	// fmt.Println(newImage2)
}

func enhanceImage(grid [][]string, algorithm string) ([][]string, int){
	newGrid := expandImage(grid)
	grid = expandImage(grid)

	litCount := 0
	for i := 1; i < len(grid) - 1; i++{
		for j := 1; j < len(grid[0]) - 1; j ++{
			enhancedPixel := enhancePixel(i, j, grid, algorithm)
			if enhancedPixel == "#"{
				litCount++
			}
			newGrid[i][j] = enhancedPixel
		}
	}

	return newGrid, litCount
}

func enhancePixel(x, y int, grid [][]string, algorithm string) string{
	bits := ""
	bits += grid[x-1][y-1]
	bits += grid[x-1][y]
	bits += grid[x-1][y+1]
	bits += grid[x][y-1]
	bits += grid[x][y]
	bits += grid[x][y+1]
	bits += grid[x+1][y-1]
	bits += grid[x+1][y]
	bits += grid[x+1][y+1]

	binary := ""
	for _, char := range bits{
		if char == '.'{
			binary += "0"
		}else if char == '#'{
			binary += "1"
		}
	}

	i, err := strconv.ParseInt(binary, 2, 64)
	check(err)

	// fmt.Println("algorithm[",i,"]: ", algorithm[i])
	return string(algorithm[i])
}

// # . .     . . . . . . .                                                                                                . . . . . . .
// . # .  -> . . . . . . .                                                                                                . X X X X X .
// . . #     . . # . . . .                                                                                                . X X X X X .
//           . . . # . . .      ->      We will be checking coords that could possibly have a # character next to it ->   . X X X X X .
//           . . . . # . .                                                                                                . X X X X X .
//           . . . . . . .                                                                                                . X X X X X .
//           . . . . . . .                                                                                                . . . . . . .

func expandImage(image [][]string) [][]string{
	newImage := make([][]string, len(image) + 4) //[len(image) + 4][len(image[0]) + 4]string

	for i := 0; i < len(image) + 4; i++{
		row := make([]string, len(image[0]) + 4)
		
		if i < 2 || i >= len(image) + 2{
			for j:=0; j < len(row); j++{
				row[j] = "."
			}
		}else{
			for j:=0; j < len(image[0]) + 4; j++ {
				if j < 2  || j >= len(image[0]) + 2{
					row[j] = "."
				}else{
					row[j] = image[i-2][j-2]
				}
			}
		}
		newImage[i] = row
	}

	return newImage
}

func parseGrid(lines []string) [][]string {
	grid := make([][]string, len(lines))
	for i, line := range lines{
		row := make([]string, len(line))
		for j, char := range line{
			row[j] = string(char)
		}
		grid[i] = row
	}
	return grid
}

func printGrid(grid [][]string){
	for _, row := range grid{
		for _, entry := range row{
			fmt.Print(" ",entry, " ")
		}
		fmt.Println()
		fmt.Println()
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