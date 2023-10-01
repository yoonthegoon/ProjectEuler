package utils 

func Divisors(n int) []int {
    var divisors []int
    for i := 1; i * i <= n; i++ {
        if n % i == 0 {
            divisors = append(divisors, i)
            divisors = append(divisors, n / i)
        }
    }
    return divisors
}

func IsPrime(n int) bool {
    if n < 2 {
        return false
    }
    for i := 2; i * i <= n; i++ {
        if n % i == 0 {
            return false
        }
    }
    return true
}

func PrimeFactors(n int) []int {
    var primeFactors []int
    for i := 2; i * i <= n; i++ {
        if n % i == 0 && IsPrime(i) {
            primeFactors = append(primeFactors, i)
            primeFactors = append(primeFactors, PrimeFactors(n / i)...)
            return primeFactors
        }

    }
    return []int{n}
}

func PrimeSieve(limit int) []int {
    var primes []int
    sieve := make([]bool, limit + 1)
    for i := 2; i <= limit; i++ {
        if !sieve[i] {
            primes = append(primes, i)
            for j := i * i; j <= limit; j += i {
                sieve[j] = true
            }
        }
    }
    return primes
}

func ReverseInts(nums []int) []int {
    var reversedSlice []int
    for i := len(nums) - 1; i >= 0; i-- {
        reversedSlice = append(reversedSlice, nums[i])
    }
    return reversedSlice
}

func SumInts(nums []int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    return sum
}

