package day04

import (
	"math"
	"strings"

	"golang.org/x/exp/slices"
)

func Part1(lines []string) int {
	ret := 0

	for _, line := range lines {
		cardPoints := sameNumbers(line)-1

		if cardPoints > -1 {
			ret += int(math.Pow(2, float64(cardPoints)))
		}
	}

	return ret
}

func Part2(lines []string) int {
	cardInstances := make([]int, len(lines))

    for idx, line := range lines {
        cardInstances[idx] += 1
        sameNumbers := sameNumbers(line)

        for i := 1; i <= sameNumbers; i++ {
            cardInstances[idx+i] += cardInstances[idx]
        }
    }

	ret := 0
    for _, cardInstance := range cardInstances {
        ret += cardInstance
    }

	return ret
}

func sameNumbers(line string) int {
	numbers := strings.TrimSpace(strings.Split(line, ":")[1])

	cardNumbers := strings.Split(numbers, "|")

	winners := strings.Split(strings.TrimSpace(cardNumbers[0]), " ")
	pocketNumbers := strings.Split(strings.TrimSpace(cardNumbers[1]), " ")

	cardPoints := 0
	for _, winner := range winners {
		if winner != "" && slices.Contains(pocketNumbers, winner) {
			cardPoints++
		}
	}

	return cardPoints
}
