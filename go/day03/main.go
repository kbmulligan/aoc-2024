package main

import ( 
    "fmt"
    "os"
    "strings"
    "regexp"
    "strconv"
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


    doSegments := strings.Split(inputString, "do()")
    fmt.Println(len(doSegments))

    segments := make([]string, 0)
    for i,seg := range doSegments {
        doNotSegments := strings.Split(seg, "don't()")
        fmt.Println(i, len(doNotSegments))
        segments = append(segments, doNotSegments[0])
    }
    




    sumTotal := 0
    for i,s := range segments {

        fmt.Println("========================================")
        //commands := strings.SplitAfter(newString, ")")
        re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
        commands := re.FindAllString(s, -1)

        validCommands := make([]string, 0)
        for _, command := range commands {
            if (commandIsValid(command)) {
                validCommands = append(validCommands, command)
            } 
            //else {
            //    fmt.Println("Invalid: ", command) 
            //}
        }

        fmt.Println(len(validCommands))
        
         
        trimmedCommands := make([]string, 0)
        for _, c := range validCommands {
            noMul := strings.ReplaceAll(c, "mul", "")
            noLeft := strings.ReplaceAll(noMul, "(", "")
            noRight := strings.ReplaceAll(noLeft, ")", "")
            trimmedCommands = append(trimmedCommands, noRight)
        }

        segmentTotal := 0
        for _, c := range trimmedCommands {
            arguments := strings.Split(c, ",")
            //fmt.Println(arguments)
            //fmt.Println(c)

            arg1, _ := strconv.Atoi(arguments[0])
            arg2, _ := strconv.Atoi(arguments[1])
            segmentTotal += arg1 * arg2
        }
        
        fmt.Println(i, segmentTotal)
        sumTotal += segmentTotal
    }
    
    fmt.Println("Sum Total: ", sumTotal)
}

func commandIsValid (c string) bool {
    matched, err := regexp.MatchString(`mul\(\d{1,3},\d{1,3}\)`, c)
    check(err)
    return matched
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
