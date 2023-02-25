# https://projecteuler.net/problem=19


from datetime import date


def main():
    sundays = 0
    for year in range(1901, 2001):
        for month in range(1, 13):
            if date(year, month, 1).weekday() == 6:
                sundays += 1
    return sundays


if __name__ == "__main__":
    print(main())
