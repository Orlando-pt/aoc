package day01

import (
	"sort"
	"strings"

	"github.com/orlando-pt/aoc/2024/utils"
)

func Part1(lines []string) int {
	col1, col2 := getSortedColumns(lines)
	return findColDistances(col1, col2)
}

func Part2(lines []string) int {
	sum := 0
	col1, col2 := getSortedColumns(lines)

	col2Occurrencies := mapNumberOccurrencies(col2)

	for _, n := range col1 {
		if occurrencies, ok := col2Occurrencies[n]; ok {
			sum += n * occurrencies
		}
	}

	return sum
}

func getSortedColumns(lines []string) (col1, col2 []int) {
	col1 = make([]int, len(lines))
	col2 = make([]int, len(lines))

	for _, line := range lines {
		cols := strings.Split(line, " ")

		col1 = append(col1, utils.StrToInt(cols[0]))
		col2 = append(col2, utils.StrToInt(cols[len(cols)-1]))
	}

	sort.Slice(col1, func(i, j int) bool {
		return col1[i] < col1[j]
	})

	sort.Slice(col2, func(i, j int) bool {
		return col2[i] < col2[j]
	})

	return
}

func findColDistances(col1, col2 []int) (sum int) {
	for i := 0; i < len(col1); i++ {
		sum += utils.IntAbs(col1[i] - col2[i])
	}

	return
}
func mapNumberOccurrencies(numbers []int) map[int]int {
	occurrences := make(map[int]int)

	for _, n := range numbers {
		occurrences[n]++
	}

	return occurrences
}
