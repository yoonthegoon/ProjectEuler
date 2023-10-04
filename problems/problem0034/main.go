package main

import (
    "fmt"

    "github.com/yoonthegoon/ProjectEuler/utils"
)

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
        sum += utils.Factorial(i % 10)
    }
    return sum == n
}

func isPalindrome(n int) bool {
    var reversed int
    for i := n; i > 0; i /= 10 {
        reversed = reversed * 10 + i % 10
    }
    return reversed == n
}
