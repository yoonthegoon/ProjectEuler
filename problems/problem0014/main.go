package main

import "fmt"

func main() {
    cache := make(map[int]int)
    var maxCollatz int
    var maxCollatzStart int
    for i := 1; i < 1_000_000; i++ {
        collatz := collatz(i, cache)
        if collatz > maxCollatz {
            maxCollatz = collatz
            maxCollatzStart = i
        }
    }
    fmt.Println(maxCollatzStart)
}

func collatz(n int, cache map[int]int) int {
    if n == 1 {
        return 1
    }
    if cache[n] != 0 {
        return cache[n]
    }
    if n % 2 == 0 {
        cache[n] = 1 + collatz(n / 2, cache)
    } else {
        cache[n] = 1 + collatz(3 * n + 1, cache)
    }
    return cache[n]
}

