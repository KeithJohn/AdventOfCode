package main

import (
    "strconv"
    "fmt"
    "bufio"
    "os"
    "math"
)

func main(){
    file, err := os.Open("inputP1.txt")
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    currNum := math.MaxInt64
    numOfIncreases := 0
    for scanner.Scan(){
        nextNum, convErr := strconv.Atoi(scanner.Text())
        check(convErr)
        fmt.Println(nextNum," ", currNum)
        if(nextNum > currNum){
            numOfIncreases += 1
            fmt.Println(numOfIncreases)
        }
        currNum = nextNum
    }

    fmt.Println(numOfIncreases)
    scanErr := scanner.Err()
    check(scanErr)
}

func check(e error){
    if e != nil{
        panic(e)
    }
}
