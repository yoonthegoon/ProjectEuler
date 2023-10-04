package main

import (
    "fmt"

    "github.com/yoonthegoon/ProjectEuler/utils"
)

var factorialsCache map[int]int = make(map[int]int)

func main() {
    curiousNums := []int{}
    for i := 10; i < 10_000_000; i++ {
        if isCurious(i) {
            curiousNums = append(curiousNums, i)
        }
    }
    fmt.Println(utils.SumInts(curiousNums))
}

func isCurious(n int) bool {
    var sum int
    for i := n; i > 0; i /= 10 {
        sum += getFactorialCache(i % 10)
    }
    return sum == n
}

func getFactorialCache(n int) int {
    factorial, ok := factorialsCache[n]
    if !ok {
        factorialsCache[n] = utils.Factorial(n)
        factorial = factorialsCache[n]
    }
    return factorial
}
