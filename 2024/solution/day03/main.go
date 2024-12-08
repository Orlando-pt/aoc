package day03

import (
	"regexp"

	"github.com/orlando-pt/aoc/2024/utils"
)

func Part1(lines []string) (sum int) {
	mul := regexp.MustCompile("mul\\(\\d+,\\d+\\)")
	decimals := regexp.MustCompile("\\d+")

	for _, line := range lines {
		utils.Debug(line)
		for _, multiplications := range mul.FindAllString(line, -1) {
			numbers := decimals.FindAllString(multiplications, -1)
			sum += multiply(numbers[0], numbers[1])
		}
	}

	return
}

func Part2(lines []string) (sum int) {
	do := regexp.MustCompile("do\\(\\)")
	dont := regexp.MustCompile("don\\'t\\(\\)")
	mul := regexp.MustCompile("mul\\(\\d+,\\d+\\)")
	decimals := regexp.MustCompile("\\d+")

	for _, line := range lines {
		dos := do.FindAllStringIndex(line, -1)
		donts := dont.FindAllStringIndex(line, -1)
		muls := mul.FindAllStringIndex(line, -1)

        utils.Debug("do index:", dos)
        utils.Debug("don't index:", donts)

		invalid := invalidIntervals(dos, donts, len(line))
		utils.Debug("invalid Intervals:", invalid)

        utils.Debug("mul index:", muls)

		for _, mulIndex := range muls {
			mulString := line[mulIndex[0]:mulIndex[1]]
			numbers := decimals.FindAllString(mulString, -1)

			if canMultiply(mulIndex[0], invalid) {
				sum += multiply(numbers[0], numbers[1])
			}
		}
	}
	return
}

func canMultiply(index int, invalidIntervals [][]int) bool {
	for _, interval := range invalidIntervals {
		if index > interval[0] && index < interval[1] {
			return false
		}
	}
	return true
}

func invalidIntervals(dos, donts [][]int, lineSize int) (intervals [][]int) {
    rightBond := -1
	for _, dont := range donts {
		start := dont[0]
		if start <= rightBond {
			continue
		}

		for i, do := range dos {
			if do[0] > start {
                rightBond = do[0]
				intervals = append(intervals, []int{start, rightBond})
                break
			} else if i == len(dos)-1 {
                rightBond = lineSize
				intervals = append(intervals, []int{start, lineSize})
			}
		}
	}

	return intervals
}

func multiply(a, b string) int {
	return utils.StrToInt(a) * utils.StrToInt(b)
}
