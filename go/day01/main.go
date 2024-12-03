package main

import ( 
    "fmt"
    "os"
    "strings"
    "sort"
    "strconv"
)

func absInt (x int) int {
    if (x < 0) {
        return -x
    } 
    return x 
}

func main() {
    
    fmt.Println("Hello AOC!")

    input, err := os.ReadFile("../../input/day01.txt")
    if err != nil {
        panic("IO Error, quitting due to panic!")
    }
    inputs := strings.Split(string(input), "\n")
    inputs = inputs[:len(inputs)-1]

    
    firsts := make([]string, 0)
    seconds := make([]string, 0)
    differences := make([]int, 0)
    crossCounts := make([]int, 0)

    for _, line := range inputs {
        //fmt.Println(i, line)
        parsed := strings.Fields(line)
        //fmt.Println(parsed)
        firsts = append(firsts, parsed[0])
        seconds = append(seconds, parsed[1])
    }
    
    // print unsorted
    //fmt.Println(firsts)
    //fmt.Println(seconds)
    
    sort.Strings(firsts)
    sort.Strings(seconds)

    // print sorted
    //fmt.Println(firsts)
    //fmt.Println(seconds)

    itemCount := len(inputs)

    for i := range(itemCount) {
        a, _ := strconv.Atoi(firsts[i])
        b, _ := strconv.Atoi(seconds[i])
        diff := absInt(a - b)
        differences = append(differences, diff)
        //fmt.Println(i, diff)
    }

    sum := sumSlice(differences) 
    fmt.Println("Difference Sum:", sum)

    for _, v := range firsts {
        firstInt, _ := strconv.Atoi(v)
        count := countInt(firstInt, seconds)
        //fmt.Println(i, v, count)
        crossCounts = append(crossCounts, count * firstInt)
    }
    
    crossSum := sumSlice(crossCounts)
    fmt.Println("Similarity Sum:", crossSum)
}

func sumSlice (items []int) int {
    total := 0
    for i, _ := range(items) {
        total += items[i]
    }
    return total
}

func countInt (subj int, field []string) int {
    count := 0

    for _, v := range(field) {
        vInt, _ := strconv.Atoi(v)
        if (vInt == subj) {
            count += 1
        }
    }

    return count
}


//func parsePair (pair string) {
//    return strings.Fields(pair)
//}
