package main

import "fmt"

func main() {
    for a := 1; a <= 999; a++ {
        for b := a; a + b <= 1000; b++ {
            c := 1000 - a - b
            if a * a + b * b == c * c {
                fmt.Println(a * b * c)
                return
            }
        }
    }
}

