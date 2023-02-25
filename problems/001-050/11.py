# https://projecteuler.net/problem=11


import os

PATH = os.path.dirname(os.path.abspath(__file__))


def main():
    def horizontal(grid):
        return max(
            grid[i][j] * grid[i][j + 1] * grid[i][j + 2] * grid[i][j + 3]
            for i in range(20)
            for j in range(17)
        )

    def vertical(grid):
        return max(
            grid[i][j] * grid[i + 1][j] * grid[i + 2][j] * grid[i + 3][j]
            for i in range(17)
            for j in range(20)
        )

    def diagonal_backward(grid):
        return max(
            grid[i][j] * grid[i + 1][j + 1] * grid[i + 2][j + 2] * grid[i + 3][j + 3]
            for i in range(17)
            for j in range(17)
        )

    def diagonal_forward(grid):
        return max(
            grid[i][j] * grid[i + 1][j - 1] * grid[i + 2][j - 2] * grid[i + 3][j - 3]
            for i in range(17)
            for j in range(3, 20)
        )

    with open(f"{PATH}/11.txt") as f:
        grid = [list(map(int, line.split())) for line in f]

    return max(
        horizontal(grid),
        vertical(grid),
        diagonal_backward(grid),
        diagonal_forward(grid),
    )


if __name__ == "__main__":
    print(main())
