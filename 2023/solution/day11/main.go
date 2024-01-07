package day11

import (
	"fmt"
	"math"
	"strings"
)

func Part1(lines []string) int {
	universe := parseInput(lines)

	galaxies := findGalaxyCoordinates(universe)

	sum := 0
	for i, galaxy := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			distance := galaxy.manhattanDistance(galaxies[j])
			fmt.Println(galaxy, galaxies[j], distance)
			sum += distance
		}
	}
	fmt.Println(galaxies)
	return sum
}

func Part2(lines []string) int {
	return 0
}

func (g galaxy) manhattanDistance(other galaxy) int {
	x := math.Abs(float64(g.coordinates.x - other.coordinates.x))
	y := math.Abs(float64(g.coordinates.y - other.coordinates.y))
	return int(x + y)
}

func findGalaxyCoordinates(universe [][]int) []galaxy {
	var galaxies []galaxy

	for x, line := range universe {
		for y, obj := range line {
			if obj != 0 {
				galaxies = append(galaxies, galaxy{
					id: obj,
					coordinates: coordinate{
						x: x,
						y: y,
					},
				})
			}
		}
	}

	return galaxies
}

type galaxy struct {
	id          int
	coordinates coordinate
}

type coordinate struct {
	x int
	y int
}

func parseInput(lines []string) [][]int {
	universe := [][]int{}

	// check column idx expansions
	var verticalExpansionIdx []int
	for i := 0; i < len(lines); i++ {
		empty := true
		for _, line := range lines {
			if line[i] != '.' {
				empty = false
				break
			}
		}

		if empty {
			verticalExpansionIdx = append(verticalExpansionIdx, i)
		}
	}

	var horizontalExpansionIdx []int
	for i, line := range lines {
		if !strings.Contains(line, "#") {
			horizontalExpansionIdx = append(horizontalExpansionIdx, i)
		}
	}

	galaxyNr := 1
	width := len(lines[0]) + len(verticalExpansionIdx)
	for _, line := range lines {
		if !strings.Contains(line, "#") {
			universe = append(universe, make([]int, width))
			universe = append(universe, make([]int, width))
			continue
		}

		measure := make([]int, width)

		for j, char := range line {
			switch char {
			case '.':
				// do nothing
			case '#':
				measure[j+calculateOffset(verticalExpansionIdx, j)] = galaxyNr
				galaxyNr++
			default:
				panic("invalid character: " + string(char))
			}
		}
		universe = append(universe, measure)
	}

	return universe
}

func calculateOffset(vIdx []int, idx int) int {
	offset := 0

	for _, i := range vIdx {
		if idx > i {
			offset++
		}
	}

	return offset
}
