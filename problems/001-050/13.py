# https://projecteuler.net/problem=13


import os

PATH = os.path.dirname(os.path.abspath(__file__))


def main():
    with open(f"{PATH}/13.txt") as f:
        numbers = [int(n) for n in f.readlines()]

    return str(sum(numbers))[:10]


if __name__ == "__main__":
    print(main())
