package main

import (
    "fmt"

    "github.com/yoonthegoon/ProjectEuler/utils"
)

func main() {
    var prime int
    var nthPrime int
    for i := 1; prime < 10_001; i++ {
        if utils.IsPrime(i) {
            prime++
            nthPrime = i
        }
    }
    fmt.Println(nthPrime)
}

