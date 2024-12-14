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
            // fix it
			utils.Debug("Valid sequence: ", sequence)
			sum += utils.StrToInt(sequence[len(sequence)/2])
		}
	}

	return
}

func fixSequence(sequence []string, precedence PrecedenceMap) (fixedSequence []string) {
	for i := 0; i < len(sequence)-1; i++ {
        rules := precedence[sequence[i]]
        portion := sequence[i+1:]
		if !checkRules(rules, portion) {
            // TODO: continue from here
			return 
		}
	}
    return
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
