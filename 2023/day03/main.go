package day03

import (
	"math"
	"regexp"

	"github.com/orlando-pt/aoc/2023/utils"
)

type Coord struct {
	x int
	y int
}

type Number struct {
    start Coord
    end Coord
    val string
}

func Part1(lines []string) int {
	ret := 0
	numbers, specialCharsCoords := parseLines(lines)

    for _, specialCoord := range specialCharsCoords {
        for _, number := range numbers {
            if specialCoord.isAdjacentNumber(number) {
                ret += utils.StrToInt(number.val)
            }
        }
    }
	return ret
}

func Part2(lines []string) int {
	ret := 0
	numbers, specialCharsCoords := parseLines(lines)

    var adjacentNumbers []string
    for _, specialCoord := range specialCharsCoords {
        if lines[specialCoord.x][specialCoord.y] != '*' {
            continue
        }

        adjacentNumbers = make([]string, 0)
        for _, number := range numbers {
            if specialCoord.isAdjacentNumber(number) {
                adjacentNumbers = append(adjacentNumbers, number.val)
            }
        }

        if len(adjacentNumbers) == 2 {
            ret += utils.StrToInt(adjacentNumbers[0]) * utils.StrToInt(adjacentNumbers[1])
        }
    }
	return ret
}

func (c Coord) isAdjacent(other Coord) bool {
    return math.Abs(float64(c.x - other.x)) <= 1 && math.Abs(float64(c.y - other.y)) <= 1
}

func (c Coord) isAdjacentNumber(number Number) bool {
    return c.isAdjacent(number.start) || c.isAdjacent(number.end)
}

func parseLines(lines []string) ([]Number, []Coord) {
	numbers := make([]Number, 0)
	specialChars := make([]Coord, 0)

    numbersRegex := regexp.MustCompile(`\d+`)
    specialRegex := regexp.MustCompile(`[^\d\.]`)
    for x, line := range lines {

        numbersCoords := numbersRegex.FindAllStringIndex(line, -1)
        numbersMatches := numbersRegex.FindAllString(line, -1)
        for i, number := range numbersMatches {
            numbers = append(numbers, Number{
                start: Coord{x, numbersCoords[i][0]},
                end: Coord{x, numbersCoords[i][1]-1},
                val: number,
            })
        }

        specialCharsCoords := specialRegex.FindAllStringIndex(line, -1)
        for _, match := range specialCharsCoords {
            specialChars = append(specialChars, Coord{x, match[0]})
        }
    }

	return numbers, specialChars
}

