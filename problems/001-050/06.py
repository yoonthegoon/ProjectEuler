# https://projecteuler.net/problem=6


def main():
    return sum(range(1, 101)) ** 2 - sum(i**2 for i in range(1, 101))


if __name__ == "__main__":
    print(main())
