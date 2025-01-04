package day07

import (
	"fmt"
	"regexp"

	"github.com/orlando-pt/aoc/2024/utils"
)

type Equation struct {
	value   int
	numbers []int
}

func Part1(lines []string) (sum int) {
	for _, line := range lines {
		equation := parseEquation(line)
		if checkCalibration(0, 0, equation) {
			sum += equation.value
		}
	}
	return
}

func Part2(lines []string) (sum int) {
	for _, line := range lines {
		equation := parseEquation(line)
		utils.Debug("Equation: ", equation)

		if checkCalibrationPart2(0, 0, equation) {
			utils.Debug("Equation is valid: ", equation)
			sum += equation.value
		}
	}
	return
}

func checkCalibration(index, currentVal int, equation Equation) bool {
	if index == len(equation.numbers) || currentVal > equation.value {
		return currentVal == equation.value
	}
	return checkCalibration(index+1, currentVal+equation.numbers[index], equation) ||
		checkCalibration(index+1, currentVal*equation.numbers[index], equation)
}

func checkCalibrationPart2(index, currentVal int, equation Equation) bool {
	if index == len(equation.numbers) || currentVal > equation.value {
		return currentVal == equation.value
	}
	return checkCalibrationPart2(index+1, currentVal+equation.numbers[index], equation) ||
		checkCalibrationPart2(index+1, currentVal*equation.numbers[index], equation) ||
		checkCalibrationPart2(index+1, concatNumbers(currentVal, equation.numbers[index]), equation)
}

func concatNumbers(n1, n2 int) int {
	n1String := fmt.Sprintf("%d", n1)
	n2String := fmt.Sprintf("%d", n2)
	return utils.StrToInt(n1String + n2String)
}

func parseEquation(line string) (equation Equation) {
	content := regexp.MustCompile(`(\d+)`).FindAllString(line, -1)
	equation.value = utils.StrToInt(content[0])

	for _, number := range content[1:] {
		equation.numbers = append(equation.numbers, utils.StrToInt(number))
	}
	return
}
