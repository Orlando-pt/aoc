package day11

import (
	"regexp"
	"strconv"

	"github.com/orlando-pt/aoc/2024/utils"
)

func Part1(lines []string) (sum int) {
	return solve(lines, 25)
}

func Part2(lines []string) (sum int) {
	return solve(lines, 75)
}

func solve(lines []string, blinks int) (sum int) {
	numbers := parseInput(lines)
	utils.Debug(numbers)

	for blink := 0; blink < blinks; blink++ {
		newNumbers := map[int]int{}
		for number, occurrences := range numbers {
			if occurrences > 0 {
				evolution := evolve(number)
				for _, e := range evolution {
					newNumbers[e] += occurrences
				}

				utils.Debug("Evolution of: ", number, " is ", evolution)
				numbers[number] -= occurrences
			}
		}

		for number, occurrences := range newNumbers {
			numbers[number] += occurrences
		}
	}

	for _, occurrences := range numbers {
		if occurrences > 0 {
			sum += occurrences
		}
	}

	return
}

func evolve(number int) (evolution []int) {
	if number == 0 {
		return []int{1}
	}

	stringNumber := strconv.Itoa(number)
	if len(stringNumber)%2 == 0 {
		firstHalf := utils.StrToInt(stringNumber[:len(stringNumber)/2])
		secondHalf := utils.StrToInt(stringNumber[len(stringNumber)/2:])
		return []int{firstHalf, secondHalf}
	}

	return []int{number * 2024}
}

func parseInput(lines []string) (numbers map[int]int) {
	numbers = make(map[int]int)
	numbersInLine := regexp.MustCompile(`\d+`).FindAllString(lines[0], -1)
	for _, n := range numbersInLine {
		numbers[utils.StrToInt(n)] = 1
	}
	return
}
