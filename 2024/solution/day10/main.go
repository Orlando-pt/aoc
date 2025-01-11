package day10

import (
	"github.com/orlando-pt/aoc/2024/utils"
)

type Index struct {
	x, y int
}

type PositionsVisited map[Index]bool

var DIRECTIONS = [...]Index{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}

const (
	TRAILEND = 9
)

func Part1(lines []string) (sum int) {

	for y, line := range lines {
		zeros := getZeroPositions(line)
		for _, zero := range zeros {
			positionsVisited := make(PositionsVisited)
			sum += findHikePaths(lines, Index{zero, y}, &positionsVisited)
		}
	}

	return
}

func Part2(lines []string) (sum int) {
	for y, line := range lines {
		zeros := getZeroPositions(line)
		for _, zero := range zeros {
			sum += calculateRating(lines, Index{zero, y})
		}
	}
	return
}

func findHikePaths(lines []string, position Index, visited *PositionsVisited) (sum int) {
	utils.Debug("X: ", position.x, " Y: ", position.y)

	currentHeight := bToInt(lines[position.y][position.x])
	if currentHeight == TRAILEND {
		if !(*visited)[position] {
			(*visited)[position] = true
			return 1
		}
		return 0
	}

	utils.Debug("Current height: ", currentHeight)

	width, height := len(lines[0]), len(lines)
	for _, direction := range DIRECTIONS {
		newPosition := Index{position.x + direction.x, position.y + direction.y}

		if checkBonds(newPosition, width, height) && bToInt(lines[newPosition.y][newPosition.x]) == currentHeight+1 {
			sum += findHikePaths(lines, newPosition, visited)
		}
	}

	return
}

func calculateRating(lines []string, position Index) (sum int) {
	currentHeight := bToInt(lines[position.y][position.x])
	if currentHeight == TRAILEND {
		return 1
	}

	width, height := len(lines[0]), len(lines)
	for _, direction := range DIRECTIONS {
		newPosition := Index{position.x + direction.x, position.y + direction.y}
		if checkBonds(newPosition, width, height) && bToInt(lines[newPosition.y][newPosition.x]) == currentHeight+1 {
			sum += calculateRating(lines, newPosition)
		}

	}
	return
}

func bToInt(r byte) int {
	return int(r - '0')
}

func getZeroPositions(line string) (positions []int) {
	for i, char := range line {
		if char == '0' {
			positions = append(positions, i)
		}
	}

	return
}

func checkBonds(pos Index, width, height int) bool {
	return pos.x >= 0 && pos.x < width && pos.y >= 0 && pos.y < height
}
