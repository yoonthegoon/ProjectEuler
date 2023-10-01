package main

import (
    "fmt"
    "math"

    "github.com/yoonthegoon/ProjectEuler/utils"
)

func main() {
    n := 10
    primeFactors := make(map[int]int)
    for i := 2; i <= n; i++ {
        for _, primeFactor := range utils.PrimeFactors(i) {
            if primeFactors[primeFactor] == 0 {
                primeFactors[primeFactor]++
            }

    }
    fmt.Println(primeFactors)
    product := 1
    for primeFactor, exponent := range primeFactors {
        product *= int(math.Pow(float64(primeFactor), float64(exponent)))
    }
    fmt.Println(product)
}
