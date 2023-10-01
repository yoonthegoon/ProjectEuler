package main

import "fmt"

func main() {
    fmt.Println(sumEven(fibLimit(4_000_000)))
}

func fibLimit(limit int) []int {
    var f []int
    for i, j := 0, 1; i < limit; i, j = i + j, i {
        f = append(f, i)
    }
    return f
}

func sumEven(s []int) int {
    sum := 0
    for _, v := range s {
        if v % 2 == 0 {
            sum += v
        }
    }
    return sum
}
