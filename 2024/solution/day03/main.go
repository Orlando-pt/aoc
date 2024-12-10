package day03

import (
	"regexp"

	"github.com/orlando-pt/aoc/2024/utils"
)

func Part1(lines []string) (sum int) {
	mul := regexp.MustCompile(`mul\(\d+,\d+\)`)
	decimals := regexp.MustCompile(`\d+`)

	for _, line := range lines {
		for _, multiplications := range mul.FindAllString(line, -1) {
			numbers := decimals.FindAllString(multiplications, -1)
			sum += multiply(numbers[0], numbers[1])
		}
	}

	return
}

func Part2(lines []string) (sum int) {
	operationsReg := regexp.MustCompile(`do\(\)|don't\(\)|mul\(\d+,\d+\)`)
	decimals := regexp.MustCompile(`\d+`)

	canMultiply := true
	for _, line := range lines {
		operations := operationsReg.FindAllString(line, -1)

		utils.Debug("operations:", operations)

		for _, operation := range operations {
			utils.Debug("operation:", operation)

			if operation == "don't()" {
				canMultiply = false
			} else if operation == "do()" {
				canMultiply = true
			} else if canMultiply {
				numbers := decimals.FindAllString(operation, -1)
				sum += multiply(numbers[0], numbers[1])
			}
		}
	}
	return
}

func multiply(a, b string) int {
	return utils.StrToInt(a) * utils.StrToInt(b)
}
