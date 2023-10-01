package main

import (
    "fmt"
    "math/big"
)

func main() {
    n := new(big.Int)
    n.MulRange(1, 100)
    sum := sumDigits(n.String())
    fmt.Println(sum)
}

func sumDigits(n string) int {
    sum := 0
    for _, digit := range n {
        sum += int(digit - '0')
    }
    return sum
}

