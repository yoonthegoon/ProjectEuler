package main

import (
    "fmt"

    "github.com/yoonthegoon/ProjectEuler/utils"
)

func main() {
    amicablePairs := []int{}
    cache := make(map[int]int)
    for a := 1; a < 10_000; a++ {
        b := d(a)
        if cache[b] == a {
            amicablePairs = append(amicablePairs, a, b)
            continue
        }
        if a == b {continue}
        cache[a] = b

    }
    fmt.Println(utils.SumInts(amicablePairs))
}

func d(n int) int {
    divisors := utils.Divisors(n)
    uniqueDivisors := utils.UniqueInts(divisors)
    divisorsSum := utils.SumInts(uniqueDivisors)
    return divisorsSum - n
}

