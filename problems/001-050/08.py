# https://projecteuler.net/problem=8


import os

PATH = os.path.dirname(os.path.abspath(__file__))


def main():
    def product(n):
        max_product = 0
        for i in range(len(n) - 12):
            product = 1
            for j in range(13):
                product *= int(n[i + j])
            if product > max_product:
                max_product = product
        return max_product

    with open(f"{PATH}/08.txt") as f:
        number = f.read()

    return product(number)


if __name__ == "__main__":
    print(main())
