package main

import (
    "fmt"
    "os"
    "strconv"
    )

func firstCell(rule, width, steps int) {
    //build the array with the initial state of each cell
    var cellState []int
    var cellNextState []int
    cellState = make([]int, width)
    cellNextState = make([]int, width)
    var mid int = width / 2
    for i := 0; i < mid; i ++ {
        cellState[i] = 0
    }
    cellState[mid] = 1
    for i := mid + 1; i < width; i++ {
        cellState[i] = 0
    }
    //print the first line
    for j := 0; j < width; j++ {
        if cellState[j] == 1 {
            fmt.Printf("#")
        }
        if cellState[j] == 0 {
            fmt.Printf(" ")
        }
    }
    fmt.Println("")
    //convert the rule into an array,
    //no matter the value of rule is 0/1 string or between 0 and 255
    var ruleArray [8]int
    if rule >= 0 && rule <= 255 {
        for num := 0; num < 8; num++ {
            ruleArray[7 - num] = rule % 2
            rule = rule / 2
        }
    }
    if rule > 255 {
        for a:=0; a < 8; a++ {
            ruleArray[7 - a] = rule % 10
            rule = rule / 10
        }
    }
    //calculate the state of each cell in the next step
    for n := 0; n < steps; n++{
        for k := 0; k < width; k++ {
            if (k == 0) {
                var cellValue int = 2 * cellState[k] + cellState[k + 1]
                cellNextState[k] = ruleArray[7 - cellValue]
            }
            if (k == width - 1) {
                var cellValue int = 4 * cellState[k - 1] + 2 * cellState[k]
                cellNextState[k] = ruleArray[7 - cellValue]
            }
            if (k >=1 && k <= width - 2) {
                var cellValue int = 4 * cellState[k - 1] + 2 * cellState[k] + cellState[k + 1]
                cellNextState[k] = ruleArray[7 - cellValue]
            }
        }
        for l := 0; l < width; l++ {
            if cellNextState[l] == 1 {
                fmt.Printf("#")
            }
            if cellNextState[l] == 0 {
                fmt.Printf(" ")
            }
        }
        fmt.Println("")
        for m := 0; m < width; m++ {
            cellState[m] = cellNextState[m]
        }
    }
}

func main() {
    //check if the input number is enough
    if len(os.Args) != 4 {
        fmt.Println("Error:  Not enough input number")
        return
    }
    //get the input number
    r, err1 := strconv.Atoi(os.Args[1])
    w, err2 := strconv.Atoi(os.Args[2])
    s, err3 := strconv.Atoi(os.Args[3])
    //check if the input value and format is qualified
    var checkRule string = os.Args[1]
    if len(checkRule) > 3 && len(checkRule) < 8 || len(checkRule) > 8 {
        fmt.Println("Error:  Wrong format of rule")
        return
    }

    if err1 != nil || r < 0 {
        fmt.Println("Error:  Rule should not be negative integer")
        return
    }
    if err2 != nil || w <= 0 {
        fmt.Println("Error:  Height should be a positive integer")
        return
    }
    if err3 != nil || s <= 0 {
        fmt.Println("Error:  Stepsize should be a positive integer")
        return
    }
    //use the function cellState
    firstCell(r, w, s)
}