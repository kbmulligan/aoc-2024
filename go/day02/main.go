package main

import ( 
    "fmt"
    "os"
    "strings"
    //"sort"
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


func convertReport (r []string) []int {
    rInt := make([]int, 0) 
    for _,v := range r {
        converted, _ := strconv.Atoi(v)
        rInt = append(rInt, converted)
    }
    return rInt
}

func main() {
    
    fmt.Println("Hello AOC! Day 2")

    input, err := os.ReadFile("../../input/day02.txt")
    check(err)

    reportStrings := strings.Split(string(input), "\n")
    reports := make([][]int, 0) 
    safeReports := make([]int, 0) 

    for i := range(len(reportStrings)) {
        levelStrings := strings.Split(reportStrings[i], " ")
        reports = append(reports, convertReport(levelStrings))
    }
    reports = reports[:len(reports) - 1] 
    
    for i,v := range reports {
        reportIsSafe := isSafe(v) || anyVariationSafe(v)
        if (reportIsSafe) {
            safeReports = append(safeReports, i)
        } else {
            fmt.Println("Report", i, ":", v)
            deltas := getDeltas(v)
            fmt.Println("Deltas: ", deltas)
            fmt.Println("Samesign: ", sameSign(deltas))
            fmt.Println("Slopes  : ", slopesInWindow(deltas))
            fmt.Println("Safe    : ", reportIsSafe)
        }
    }

    fmt.Println("Total Report Count:", len(reports))
    fmt.Println("Safe Report Count:", len(safeReports))
}

func getDeltas (report []int) []int {
    deltas := make([]int,0)
    for i := range(len(report) - 1) {
        deltas = append(deltas, report[i+1] - report[i])
    }
    return deltas
}

func sameSign (report []int) bool {
    return allPositive(report) || allNegative(report)
}

func allPositive (list []int) bool {
    result := true 
    for _,v := range list {
        if (v <= 0) {
            result = false
        }
    } 
    return result
}

func allNegative (list []int) bool {
    result := true 
    for _,v := range list {
        if (v >= 0) {
            result = false
        }
    } 
    return result
}

func slopesInWindow (list []int) bool {
    result := true 
    for _,v := range list {
        slope := absInt(v)
        if (slope < 1 || slope > 3) {
            result = false
        }
    } 
    return result
}

func isSafe (report []int) bool {
    deltas := getDeltas(report)
    return sameSign(deltas) && slopesInWindow(deltas)
}

func removeItem (report []int, index int) []int {
    newSlice := make([]int,0)
    for i := 0; i < len(report); i++ { 
        if (i != index) {
            newSlice = append(newSlice, report[i])
        }
    }
    return newSlice
}

func anyVariationSafe (reportOriginal []int) bool {
    anyVariationSafe := false

    report := make([]int, len(reportOriginal))
    copy(report, reportOriginal)

    fmt.Println("Report original:")
    fmt.Println(reportOriginal)
    
    fmt.Println("Report:")
    fmt.Println(report)
    fmt.Println("Variants:")


    // loop through report indexes
    for i:=0; i < len(report); i++ {
        fmt.Println(i, report[i])

        // remove this index
 
        variantReport := removeItem(report, i)
        fmt.Println(variantReport)

        // test that one
        // if safe, then variation is safe
        if (isSafe(variantReport)) {
            anyVariationSafe = true;
        } 
    }

    return anyVariationSafe
}
