package main

import "fmt"

func main() {
    nums := makeRange(1, 101, 1)
    sumOfSquares := sum(squareInts(nums))
    squareOfSum := squareInt(sum(nums))
    fmt.Println(squareOfSum - sumOfSquares)
}

func makeRange(start int, stop int, step int) []int {
    var nums []int
    for i := start; i < stop; i += step {
        nums = append(nums, i)
    }
    return nums
}

func squareInt(n int) int {
    return n * n
}

func squareInts(nums []int) []int {
    squares := make([]int, 0)
    for _, n := range nums {
        squares = append(squares, squareInt(n))
    }
    return squares
}

func sum(nums []int) int {
    var sum int
    for _, n := range nums {
        sum += n
    }
    return sum
}

