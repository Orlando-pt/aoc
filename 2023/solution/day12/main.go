package day12

import (
	"fmt"
	"strings"

	"github.com/orlando-pt/aoc/2023/utils"
)


func Part1(lines []string) (sum int) {
    records := parseInput(lines)
    fmt.Println(records)
	return 1
}

func Part2(lines []string) int {
	return 0
}

type record struct {
    conditions string
    arrangements []int
}

func parseInput(lines []string) (info []record) {

    for _, line := range lines {
        lineSplit := strings.Split(line, " ")
        var r record

        r.conditions = lineSplit[0]

        arrangementsStr := strings.Split(lineSplit[1], ",")
        for _, a := range arrangementsStr {
            r.arrangements = append(r.arrangements, utils.StrToInt(a))
        }

        info = append(info, r)
    }
    return
}

func getNumberPossibleArrangements(r record) (res int) {


    return
}

func (r record) isValid() bool {
    index := 0
    cardinality := 0
    for _, a := range r.conditions {
        if rune(a) == '#' {
            cardinality++
        } else if rune(a) == '.' {
            if cardinality != 0 {
                if cardinality != r.arrangements[index] {
                    return false
                }
            }
            cardinality = 0
        }
    }

    return true
}

