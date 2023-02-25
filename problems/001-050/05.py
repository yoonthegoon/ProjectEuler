# https://projecteuler.net/problem=5


def main():
    num = 2520
    while True:
        if all(num % i == 0 for i in range(1, 21)):
            return num
        num += 2520


if __name__ == "__main__":
    print(main())
