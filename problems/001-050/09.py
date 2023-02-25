# https://projecteuler.net/problem=9


def main():
    for a in range(1, 1000):
        for b in range(a + 1, 500):
            c = 1000 - a - b
            if a * a + b * b == c * c:
                return a * b * c


if __name__ == "__main__":
    print(main())
