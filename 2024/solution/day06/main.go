package day06

import "github.com/orlando-pt/aoc/2024/utils"

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

func Part1(lines []string) (sum int) {
	row, col := findGuardPosition(lines)
	utils.Debug("Guard position: ", row, col)

	sum = traverseMap(lines, row, col)
	return
}

func Part2(lines []string) (sum int) {
	return
}

func findGuardPosition(lines []string) (int, int) {
	for i, line := range lines {
		for j, char := range line {
			if char == '^' {
				return i, j
			}
		}
	}

	return -1, -1
}

func traverseMap(lines []string, row, col int) (distinctPositions int) {
	direction := UP
	positionsCovered := make([]bool, len(lines)*len(lines[0]))
	positionsCovered[row*len(lines)+col] = true

	for {
		nextRow, nextCol := nextCoords(row, col, direction)

		if nextRow < 0 || nextRow >= len(lines) || nextCol < 0 || nextCol >= len(lines[0]) {
			utils.Debug("Out of bounds")
			break
		}

		if lines[nextRow][nextCol] == '#' {
			utils.Debug("Found wall: ", nextRow, nextCol)
			direction = changeDirection(direction)
		}

		row, col = nextCoords(row, col, direction)
		utils.Debug("Moved to: ", row, col)

		positionsCovered[row*len(lines)+col] = true
	}

	for _, wasCovered := range positionsCovered {
		if wasCovered {
			distinctPositions++
		}
	}
	return
}

func nextCoords(row, col, direction int) (int, int) {
	switch direction {
	case UP:
		return row - 1, col
	case RIGHT:
		return row, col + 1
	case DOWN:
		return row + 1, col
	case LEFT:
		return row, col - 1
	}

	return row, col
}

func changeDirection(direction int) int {
	return (direction + 1) % 4
}
