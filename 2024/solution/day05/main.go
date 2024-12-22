package day05

import (
	"strings"

	"github.com/orlando-pt/aoc/2024/utils"
)

type PrecedenceMap map[string][]string

func Part1(lines []string) (sum int) {
	precedence, updateSequences := parseInput(lines)

	for _, sequence := range updateSequences {
		if checkSequence(sequence, precedence) {
			utils.Debug("Valid sequence: ", sequence)
			sum += utils.StrToInt(sequence[len(sequence)/2])
		}
	}
	return
}

func Part2(lines []string) (sum int) {
	precedence, updateSequences := parseInput(lines)

	for _, sequence := range updateSequences {
		if !checkSequence(sequence, precedence) {
			sequence = fixSequence(sequence, precedence)
			sum += utils.StrToInt(sequence[len(sequence)/2])
		}
	}

	return
}

func fixSequence(sequence []string, precedence PrecedenceMap) []string {
	l, r := 0, 1
	for r < len(sequence) {
		if utils.ArrStrContains(precedence[sequence[r]], sequence[l]) {
			sequence[l], sequence[r] = sequence[r], sequence[l]
		}
		r++

		if r == len(sequence) {
			l++
			r = l + 1
		}
	}
	return sequence
}

func checkSequence(sequence []string, precedence PrecedenceMap) bool {
	for i := 0; i < len(sequence)-1; i++ {
		if !checkRules(precedence[sequence[i]], sequence[i+1:]) {
			return false
		}
	}

	return true
}

func checkRules(pagePrecedence []string, sequence []string) bool {
	utils.Debug("Check rules: ", pagePrecedence, sequence)
	for _, page := range sequence {
		if !utils.ArrStrContains(pagePrecedence, page) {
			return false
		}
	}

	return true
}

func parseInput(lines []string) (precedence PrecedenceMap, updateSequence [][]string) {
	precedence = make(PrecedenceMap)

	i := 0
	for ; i < len(lines); i++ {
		if lines[i] == "" {
			break
		}

		numbers := strings.Split(lines[i], "|")
		precedence[numbers[0]] = append(precedence[numbers[0]], numbers[1])
	}
	utils.Debug("Precedence: ", precedence)

	for i++; i < len(lines); i++ {
		sequence := strings.Split(lines[i], ",")
		updateSequence = append(updateSequence, sequence)
	}
	utils.Debug("Sequences: ", updateSequence)

	return
}
