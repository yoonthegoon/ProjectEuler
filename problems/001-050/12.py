# https://projecteuler.net/problem=12


def main():
    def divisors(n):
        _divisors = {1, n}
        for i in range(2, int(n**0.5) + 1):
            if n % i == 0:
                _divisors.add(i)
                _divisors.add(n // i)
        return _divisors

    def triangle_number(n):
        return n * (n + 1) // 2

    for i in range(1, 100000):
        if len(divisors(triangle_number(i))) > 500:
            return triangle_number(i)


if __name__ == "__main__":
    print(main())
