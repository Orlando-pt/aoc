package day02

import (
	"strings"

	"github.com/orlando-pt/aoc/2024/utils"
)

func Part1(lines []string) (sum int) {
	matrix := loadMatrix(lines)

	for _, row := range matrix {
		if checkConditions(row) {
			sum++
		}
	}

	return
}

func Part2(lines []string) (sum int) {
	matrix := loadMatrix(lines)

	for _, row := range matrix {

		if checkConditions(row) {
			sum++
		} else if checkWithDampener(row) {
			sum++
		}
	}

	return
}

func checkWithDampener(row []int) bool {
	for i := 0; i < len(row); i++ {
		if checkConditions(removeElement(row, i)) {
			return true
		}
	}
	return false
}

func checkConditions(row []int) bool {
	condition := getOrderedCondition(row[0] < row[1])

	return checkOrdered(row, condition) && checkAdjacent(row)
}

func getOrderedCondition(increasing bool) func(n1, n2 int) bool {
	if increasing {
		return func(n1, n2 int) bool { return n1 > n2 }
	}

	return func(n1, n2 int) bool { return n1 < n2 }
}

func removeElement(slice []int, s int) (arr []int) {
	arr = make([]int, len(slice)-1)
	index := 0

	for i, v := range slice {
		if i == s {
			continue
		}

		arr[index] = v
		index++
	}
	return
}

func checkAdjacent(numbers []int) bool {
	for i := 1; i < len(numbers); i++ {
		if abs := utils.IntAbs(numbers[i-1] - numbers[i]); abs < 1 || abs > 3 {
			return false
		}
	}
	return true
}

func checkOrdered(numbers []int, condition func(n1, n2 int) bool) bool {
	for i := 1; i < len(numbers); i++ {
		if condition(numbers[i-1], numbers[i]) {
			return false
		}
	}
	return true
}

func loadMatrix(lines []string) (matrix [][]int) {
	matrix = make([][]int, len(lines))
	for i, line := range lines {
		numbers := strings.Split(line, " ")

		matrix[i] = make([]int, len(numbers))

		for j, number := range numbers {
			matrix[i][j] = utils.StrToInt(number)
		}
	}

	return
}
