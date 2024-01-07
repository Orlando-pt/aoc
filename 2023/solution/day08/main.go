package day08

import "strings"

func Part1(lines []string) int {
	instructions := parseLines(lines)

	return instructions.find("AAA", func(location string) bool {
		return location == "ZZZ"
	})
}

func Part2(lines []string) int {
	instructions := parseLines(lines)

	startingPoints := instructions.getStartingPoints()
	var solutions []int

	for _, startingPoint := range startingPoints {
		solutions = append(solutions, instructions.find(startingPoint, func(location string) bool {
			return location[2] == 'Z'
		}))
	}
	return lcmAll(solutions)
}

func (ins instructions) find(start string, foundObjective func(string) bool) int {
	currentLocation := start
	currentDirectionIdx := 0
	steps := 0

	for !foundObjective(currentLocation) {
		if ins.directions[currentDirectionIdx] == 'L' {
			currentLocation = ins.navigationMap[currentLocation][0]
		} else {
			currentLocation = ins.navigationMap[currentLocation][1]
		}

		steps += 1
		currentDirectionIdx = (currentDirectionIdx + 1) % len(ins.directions)
	}
	return steps
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// thanks to github.com/Guergeiro
func lcmAll(numbers []int) int {
	multiple := numbers[0]
	for _, n := range numbers[1:] {
		multiple = lcm(multiple, n)
	}

	return multiple
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
