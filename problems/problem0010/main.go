package main

import (
    "fmt"

    "github.com/yoonthegoon/ProjectEuler/utils"
)

func main() {
    primes := utils.PrimeSieve(2_000_000)
    sum := utils.SumInts(primes)
    fmt.Println(sum)
}

