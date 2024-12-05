package main

import ( 
    "fmt"
    "os"
    "strings"
    //"sort"
    //"strconv"
)

func absInt (x int) int {
    if (x < 0) {
        return -x
    } 
    return x 
}

func check (e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    
    fmt.Println("Hello AOC! Day 3")

    input, err := os.ReadFile("../../input/day03.txt")
    check(err)

    inputString := string(input)
    fmt.Println(inputString)

    newString := ""

    for _, v := range inputString {
        if (runeAllowed(v)) {
            newString += string(v)
        }
    }    

    fmt.Println("========================================")
    commands := strings.SplitAfter(newString, ")")

    validCommands := make([]string, 0)
    for _, command := range commands {
        if (strings.Contains(command, "mul")) {
            validCommands = append(validCommands, command)
        }
    }

    for _, command := range validCommands {
        fmt.Println(command)
    }

    for _, command := range validCommands {
        fmt.Println(command)
    }

    fmt.Println(len(validCommands))
    
}

func runeAllowed (v rune) bool {
    allowedSet := "mul(),012345679"
    allowed := false
    for _,r := range allowedSet {
        if (v == r) {
            allowed = true
        }
    }
    return allowed   
}
