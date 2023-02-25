# https://projecteuler.net/problem=7

import sys

sys.path.append("/home/debian/repos/ProjectEuler")

from utilities import is_prime


def main():
    num = 1
    count = 0
    while True:
        if is_prime(num):
            count += 1
        if count == 10001:
            return num
        num += 1


if __name__ == "__main__":
    print(main())
