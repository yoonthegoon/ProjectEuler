package main

import (
    "fmt"

    "github.com/yoonthegoon/ProjectEuler/utils"
)

func main() {
    var maxDivisors int
    var triangleNumber int
    for i := 1; maxDivisors <= 500; i++ {
        triangleNumber += i
        divisors := utils.Divisors(triangleNumber)
        maxDivisors = len(divisors)
    }
    fmt.Println(triangleNumber)
}

