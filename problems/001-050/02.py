# https://projecteuler.net/problem=2


def main():
    def fibonacci(n):
        if n == 0:
            return 0
        if n == 1:
            return 1
        return fibonacci(n - 1) + fibonacci(n - 2)

    total_sum = 0
    i = 0
    while fibonacci(i) < 4000000:
        if fibonacci(i) % 2 == 0:
            total_sum += fibonacci(i)
        i += 1
    return total_sum


if __name__ == "__main__":
    print(main())
