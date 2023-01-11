package main

import (
	"adventofcode/util"
	"fmt"
	"strconv"
	"strings"
)

type File struct {
	parentDir   *File
	name        string
	size        int
	children    map[string]*File
	isDirectory bool
}

func main() {
	lines := util.ReadInput("../input.txt")
	//First line is cd /
	currCommand := "cd"
	currFile := createFile(nil, "/", 0, true)
	rootFile := currFile
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		//fmt.Println(line)
		if strings.HasPrefix(line, "$") {
			//Command
			commandParts := strings.Split(line, " ")
			currCommand = commandParts[1]
			if len(commandParts) > 2 {
				//cd command
				//Find what dir to change to
				if commandParts[2] == ".." {
					currFile = currFile.parentDir
				} else {
					fmt.Println("CDing to ", strings.Trim(commandParts[2], " "))
					currFile = currFile.children[strings.Trim(commandParts[2], " ")]
				}
			}
		} else {
			//Parse based on command
			if currCommand == "ls" {
				outputParts := strings.Split(line, " ")
				if outputParts[0] == "dir" {
					//fmt.Println("Creating Directory: ", outputParts[1])
					newFile := createFile(currFile, outputParts[1], 0, true)
					newFile.printFile(0)
				} else {
					//fmt.Println("Creating File: ", outputParts[1])
					size, err := strconv.Atoi(outputParts[0])
					util.Check(err)
					createFile(currFile, outputParts[1], size, false)
				}
			}
		}
	}
	printDirTree(rootFile, 0)
	fmt.Println(getTotalSum(rootFile))
}

func createFile(parentDir *File, name string, size int, isDir bool) *File {
	newFile := new(File)
	newFile.parentDir = parentDir
	newFile.name = name
	newFile.size = size
	newFile.isDirectory = isDir
	newFile.children = make(map[string]*File)
	if parentDir != nil {
		parentDir.children[newFile.name] = newFile
	}
	return newFile
}

func printDirTree(file *File, depth int) {

	file.printFile(depth)

	for _, value := range file.children {
		printDirTree(value, depth+1)
	}
}

func (f *File) printFile(depth int) {
	padding := strings.Repeat("   ", depth)
	if f.isDirectory {
		fmt.Println(padding, "dir ", f.name, " ", getSize(f))
		fmt.Println(padding, "|__")
	} else {
		fmt.Println(padding, f.size, f.name)
	}
}

func getTotalSum(f *File) int {
	currTotal := 0
	for _, element := range f.children {
		currTotal += getTotalSum(element)
	}

	if f.isDirectory && f.size <= 100000 {
		currTotal += f.size
	}
	return currTotal
}

func getSize(f *File) int {
	currSize := 0
	if f.isDirectory {
		for _, element := range f.children {
			currSize += getSize(element)
		}
	} else {
		currSize += f.size
	}
	f.size = currSize
	return currSize
}
