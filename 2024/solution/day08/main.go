package day08

import "github.com/orlando-pt/aoc/2024/utils"

type Index struct {
	x, y int
}

type AntennaMap map[rune][]Index

func Part1(lines []string) (sum int) {
	validLocations := make(map[Index]bool)
	maxX, maxY := len(lines[0]), len(lines)

	antennas := parseInput(lines)
	utils.Debug(antennas)

	for _, antenna := range antennas {
		for i := 0; i < len(antenna); i++ {
			for j := i + 1; j < len(antenna); j++ {
				antinode1, antinode2 := getAntinodes(antenna[i], antenna[j])
				if antinodeIsValid(maxX, maxY, antinode1) {
					validLocations[antinode1] = true
				}
				if antinodeIsValid(maxX, maxY, antinode2) {
					validLocations[antinode2] = true
				}
			}
		}
	}
	return len(validLocations)
}

func Part2(lines []string) (sum int) {
	validLocations := make(map[Index]bool)
	maxX, maxY := len(lines[0]), len(lines)

	antennas := parseInput(lines)

	for _, antenna := range antennas {
		for i := 0; i < len(antenna); i++ {
			for j := i + 1; j < len(antenna); j++ {
				deltaX := antenna[i].x - antenna[j].x
				deltaY := antenna[i].y - antenna[j].y

				nextIdx := Index{antenna[i].x - deltaX, antenna[i].y - deltaY}
				for antinodeIsValid(maxX, maxY, nextIdx) {
					validLocations[nextIdx] = true
					nextIdx = Index{nextIdx.x - deltaX, nextIdx.y - deltaY}
				}

				nextIdx = Index{antenna[j].x + deltaX, antenna[j].y + deltaY}
				for antinodeIsValid(maxX, maxY, nextIdx) {
					validLocations[nextIdx] = true
					nextIdx = Index{nextIdx.x + deltaX, nextIdx.y + deltaY}
				}
			}
		}
	}
	return len(validLocations)
}

func parseInput(lines []string) (antennas AntennaMap) {
	antennas = make(AntennaMap)
	for y, line := range lines {
		for x, char := range line {
			if char != '.' {
				antennas[char] = append(antennas[char], Index{x, y})
			}
		}
	}
	return
}

func getAntinodes(i1, i2 Index) (first, second Index) {
	return Index{2*i1.x - i2.x, 2*i1.y - i2.y}, Index{2*i2.x - i1.x, 2*i2.y - i1.y}
}

func antinodeIsValid(maxX, maxY int, i Index) bool {
	return i.x >= 0 && i.x < maxX && i.y >= 0 && i.y < maxY
}
