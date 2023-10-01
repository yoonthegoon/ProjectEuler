package main

import (
    "fmt"
    "math"

    "github.com/yoonthegoon/ProjectEuler/utils"
)

func main() {
    var n int = 600_851_475_143
    for i := int(math.Sqrt(float64(n))); i >= 2; i-- {
        if n % i == 0 && utils.IsPrime(i) {
            fmt.Println(i)
            break
        }
    }
}
