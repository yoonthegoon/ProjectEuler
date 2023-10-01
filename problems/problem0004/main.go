package main

import "fmt"

func main() {
    var largest int
    for i := 999; i > 99; i-- {
        for j := 999; j > 99; j-- {
            if isPalindrome(i * j) && i * j > largest {
                largest = i * j
            }
        }
    }
    fmt.Println(largest)
}

func isPalindrome(n int) bool {
    var reversed int
    for i := n; i > 0; i /= 10 {
        reversed = reversed * 10 + i % 10
    }
    return reversed == n
}
