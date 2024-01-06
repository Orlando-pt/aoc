package day08

import (
	"fmt"
	"strings"
)

const start = "AAA"
const end = "ZZZ"

func Part1(lines []string) int {
	instructions := parseLines(lines)

	currentLocation := start
	currentDirectionIdx := 0
	path := ""

	for currentLocation != end {
		if instructions.directions[currentDirectionIdx] == 'L' {
			currentLocation = instructions.navigationMap[currentLocation][0]
			path += "L"
		} else {
			currentLocation = instructions.navigationMap[currentLocation][1]
			path += "R"
		}

		currentDirectionIdx = (currentDirectionIdx + 1) % len(instructions.directions)
	}

	fmt.Println(path)
	return len(path)
}

func Part2(lines []string) int {
	instructions := parseLines(lines)

	startingPoints := instructions.getStartingPoints()
	currentDirectionIdx := 0
	steps := 0

	for !solved(startingPoints) {
		startingPoints = instructions.navigationMap.getNextLocations(
			startingPoints, rune(instructions.directions[currentDirectionIdx]))

		currentDirectionIdx = (currentDirectionIdx + 1) % len(instructions.directions)
		steps++

	}

	fmt.Println(startingPoints)
	return steps
}

func solved(locations []string) bool {
	for _, location := range locations {
		if location[2] != 'Z' {
			return false
		}
	}
	return true
}

func (nav navigation) getNextLocations(locations []string, direction rune) []string {
	var nextLocations []string
	for _, location := range locations {
		if direction == 'L' {
			nextLocations = append(nextLocations, nav[location][0])
		} else {
			nextLocations = append(nextLocations, nav[location][1])
		}
	}
	return nextLocations
}

func (instructions instructions) getStartingPoints() []string {
	var startingPoints []string
	for key := range instructions.navigationMap {
		if key[2] == 'A' {
			startingPoints = append(startingPoints, key)
		}
	}

	return startingPoints
}

type navigation map[string][]string

type instructions struct {
	directions    string
	navigationMap navigation
}

func parseLines(lines []string) instructions {
	instructions := instructions{
		navigationMap: make(navigation),
	}

	instructions.directions = strings.TrimSpace(lines[0])

	for _, line := range lines[1:] {
		if line == "" {
			continue
		}

		start := line[0:3]
		left := line[7:10]
		right := line[12:15]
		instructions.navigationMap[start] = []string{left, right}
	}

	return instructions
}
