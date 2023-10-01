package main

import (
    "fmt"
    "math/big"
)

func main() {
    base := big.NewInt(2)
    exponent := big.NewInt(1000)
    n := base.Exp(base, exponent, nil)
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
