package day04

import (
	"math"
	"strings"

	"golang.org/x/exp/slices"
)

func Part1(lines []string) int {
    ret := 0

    for _, line := range lines {
        numbers := strings.TrimSpace(strings.Split(line, ":")[1])

        cardNumbers := strings.Split(numbers, "|")

        winners := strings.Split(strings.TrimSpace(cardNumbers[0]), " ")
        pocketNumbers := strings.Split(strings.TrimSpace(cardNumbers[1]), " ")

        cardPoints := -1
        for _, winner := range winners {
            if winner != "" && slices.Contains(pocketNumbers, winner) {
                cardPoints++
            }
        }

        if cardPoints > -1 {
            ret += int(math.Pow(2, float64(cardPoints)))
        }
    }

    return ret
}

func Part2(lines []string) int {
    return 0
}
