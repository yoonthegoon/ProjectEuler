# https://projecteuler.net/problem=3


import sys

sys.path.append("/home/debian/repos/ProjectEuler")

from utilities import is_prime


def main():
    def prime_factors(n):
        factors = []
        for i in range(2, int(n**0.5 + 1)):
            if n % i == 0:
                if is_prime(i):
                    factors.append(i)
                if is_prime(n // i):
                    factors.append(n // i)
        return factors

    return max(prime_factors(600851475143))


if __name__ == "__main__":
    print(main())
