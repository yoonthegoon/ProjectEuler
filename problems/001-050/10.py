import sys

sys.path.append("/home/debian/repos/ProjectEuler")

from utilities import is_prime


def main():
    prime_sum = 0
    for i in range(2, 2000000):
        if not is_prime(i):
            continue
        prime_sum += i
    return prime_sum


if __name__ == "__main__":
    print(main())
