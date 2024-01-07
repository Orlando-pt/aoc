package day09

import (
	"strings"

	"github.com/orlando-pt/aoc/2023/utils"
)

func Part1(lines []string) int {
	measures := parseInput(lines)

	sum := 0
	for _, measure := range measures {
		untilZero := calculateUntilZero([][]int{measure})

        sum += calculateNext(untilZero)
	}

	return sum
}

func Part2(lines []string) int {
	measures := parseInput(lines)

	sum := 0
	for _, measure := range measures {
		untilZero := calculateUntilZero([][]int{measure})

        previous := calculatePrevious(untilZero)
        sum += previous
	}

	return sum
}

func calculateNext(predictions [][]int) int {
	next := 0
	for _, prediction := range predictions {
		next += prediction[len(prediction)-1]
	}

	return next
}

func calculatePrevious(predictions [][]int) int {
    previous := 0
    for i := len(predictions)-1; i >=0; i-- {
        previous = predictions[i][0] - previous
    }

    return previous
}

func calculateUntilZero(numbers [][]int) [][]int {
	last := numbers[len(numbers)-1]
	if isZero(last) {
		return numbers
	}

	var nextNumbers []int
	for i := 1; i < len(last); i++ {
		nextNumbers = append(nextNumbers, last[i]-last[i-1])
	}

	return calculateUntilZero(append(numbers, nextNumbers))
}

func isZero(numbers []int) bool {
	for _, n := range numbers {
		if n != 0 {
			return false
		}
	}

	return true
}

func parseInput(lines []string) [][]int {
	var input [][]int

	for _, line := range lines {
		var numbers []int
		numbersString := strings.Split(line, " ")
		for _, n := range numbersString {
			numbers = append(numbers, utils.StrToInt(n))
		}
		input = append(input, numbers)
	}

	return input
}
