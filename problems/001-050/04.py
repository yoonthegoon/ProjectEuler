# https://projecteuler.net/problem=4


def main():
    largest_palindrome = 0
    for i in range(100, 1000):
        for j in range(100, 1000):
            product = i * j
            if str(product) == str(product)[::-1]:
                if product > largest_palindrome:
                    largest_palindrome = product
    return largest_palindrome


if __name__ == "__main__":
    print(main())
