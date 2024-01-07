package day10

import (
	"fmt"
)

const (
	ns     = '|'
	ew     = '-'
	ne     = 'L'
	nw     = 'J'
	sw     = '7'
	se     = 'F'
	ground = '.'
	start  = 'S'
)

func Part1(lines []string) int {
	start := findStart(lines)

	surroundings := start.findPossibleSurroundings(lines)

	var possibleSteps []int
	for _, surrounding := range surroundings {
		possibleSteps = append(possibleSteps, surrounding.findLoop(lines, start))

	}

	// find duplicates
	steps := 0
	for i := 0; i < len(possibleSteps); i++ {
		for j := i + 1; j < len(possibleSteps); j++ {
			if possibleSteps[i] == possibleSteps[j] {
				steps = possibleSteps[i]
				break
			}
		}
		if steps != 0 {
			break
		}
	}
	fmt.Println(steps)

	return steps / 2
}

func Part2(lines []string) int {
	return 0
}

func (c coordinate) findPossibleSurroundings(lines []string) []coordinate {
	surroundings := []coordinate{
		{c.x, c.y - 1},
		{c.x, c.y + 1},
		{c.x - 1, c.y},
		{c.x + 1, c.y},
	}

	var possibleSurroundings []coordinate
	for _, surrounding := range surroundings {
		if surrounding.isValid(len(lines)) && lines[surrounding.x][surrounding.y] != ground {
			possibleSurroundings = append(possibleSurroundings, surrounding)
		}
	}

	return possibleSurroundings
}

func (c coordinate) findLoop(lines []string, last coordinate) int {
	// fmt.Println(last, c, string(lines[c.y][c.x]))
	if !c.isValid(len(lines)) || lines[c.y][c.x] == ground {
		return 0
	}

	if lines[c.y][c.x] == start {
		return 1
	}

	return 1 + c.nextMove(last, rune(lines[c.y][c.x])).findLoop(lines, c)
}

func findStart(lines []string) coordinate {
	for y, line := range lines {
		for x, char := range line {
			if char == start {
				return coordinate{x, y}
			}
		}
	}

	return coordinate{-1, -1}
}

func (c coordinate) isValid(max int) bool {
	return c.x >= 0 && c.y >= 0 && c.x < max && c.y < max
}

func (c coordinate) nextMove(last coordinate, direction rune) coordinate {
	switch direction {
	case ns:
		if last.y < c.y {
			c.y++
		} else {
			c.y--
		}
	case ew:
		if last.x < c.x {
			c.x++
		} else {
			c.x--
		}
	case ne:
		if last.y < c.y {
			c.x++
		} else {
			c.y--
		}
	case nw:
		if last.y < c.y {
			c.x--
		} else {
			c.y--
		}
	case sw:
		if last.y > c.y {
			c.x--
		} else {
			c.y++
		}
	case se:
		if last.y > c.y {
			c.x++
		} else {
			c.y++
		}
	}

	return c
}

type coordinate struct {
	x int
	y int
}
