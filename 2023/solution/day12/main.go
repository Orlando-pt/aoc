package day12

import (
	"fmt"
	"regexp"

	"github.com/orlando-pt/aoc/2023/utils"
)

var reline = regexp.MustCompile("^([.#?]+) ([,0-9]+)")
var renum = regexp.MustCompile("[0-9]+")
var respan = regexp.MustCompile("[#]+")

func Part1(lines []string) (sum int) {
	for _, line := range lines {
		sum += arrangements(line)
	}
	return
}

func arrangements(line string) (na int) {
	lineparts := reline.FindStringSubmatch(line)
	conditions := lineparts[1]
	spans := []int{}

	for _, num := range renum.FindAllString(lineparts[2], -1) {
		spans = append(spans, utils.StrToInt(num))
	}

	arr := []byte(conditions)
	unkpos := []int{}

	for i, b := range line {
		if b == '?' {
			unkpos = append(unkpos, i)
		}
	}

	nrCombs := 1 << len(unkpos)

testCombination:
	for comb := 0; comb < nrCombs; comb++ {
		fmtStr := fmt.Sprintf("%%0%db", len(unkpos))
		binStr := fmt.Sprintf(fmtStr, comb)

		for i, b := range binStr {
			if b == '1' {
				arr[unkpos[i]] = '#'
			} else {
				arr[unkpos[i]] = '.'
			}
		}

		brokenSpans := respan.FindAll(arr, -1)
		if len(brokenSpans) != len(spans) {
			continue
		}

		for i, l := range spans {
			if len(brokenSpans[i]) != l {
				continue testCombination
			}
		}
		na++
	}

	return
}

func Part2(lines []string) (sum int) {
    for _, line := range lines {
        sum += arrNumsAll(line)
    }

	return
}

const TIMES = 5
var maxSpan int
func arrNumsAll(line string) int {
    lineparts := reline.FindStringSubmatch(line)
    conditions := lineparts[1]
    rec1 := []byte(conditions)
    spans1 := []int{}

    for _, num := range renum.FindAllString(lineparts[2], -1) {
        span := utils.StrToInt(num)
        spans1 = append(spans1, span)
        if span > maxSpan {
            maxSpan = span
        }
    }

    maxSpan++

    recb := []byte{}
    spans := []int{}
    for i := 0; i < TIMES; i++ {
        if len(recb) > 0 {
            recb = append(recb, '?')
        }
        recb = append(recb, rec1...)
        spans = append(spans, spans1...)
    }
    rec := string(recb)
    skb := make([]byte, len(spans), len(spans))

    for i, s := range spans {
        skb[i] = '0' + byte(s)
    }
    return arrNums(rec, spans)
}

func arrNums(rec string, groups []int) int {
    groupCount := len(groups)
    arrangements := []int{1}

    for _, c := range rec {
        newArrangements := []int{}
        for idx, count := range arrangements {
            if count < 0 {
                continue
            }
            groupIndex := idx / maxSpan
            currentGroupSize := idx % maxSpan
            valid := true

            if c == '#' {
                currentGroupSize++
                valid = isValid(groups, groupIndex, currentGroupSize, false)
            } else if c == '?' {
                if currentGroupSize > 0 {
                    if isValid(groups, groupIndex, currentGroupSize, true) {
                        newArrangements = sparseInc(newArrangements, groupIndex + 1, 0, count)
                    }
                } else {
                    newArrangements = sparseInc(newArrangements, groupIndex, currentGroupSize, count)
                }
                if isValid(groups, groupIndex, currentGroupSize + 1, false) {
                    newArrangements = sparseInc(newArrangements, groupIndex, currentGroupSize + 1, count)
                }
                continue
            } else if currentGroupSize > 0 {
                valid = isValid(groups, groupIndex, currentGroupSize, true)
                currentGroupSize = 0
                groupIndex++
            }
            if valid {
                newArrangements = sparseInc(newArrangements, groupIndex, currentGroupSize, count)
            }
        }
        arrangements = newArrangements
    }

    c := 0
    for idx, count := range arrangements {
        if count < 0 {
            continue
        }
        groupIndex := idx / maxSpan
        currentGroupSize := idx % maxSpan

        if currentGroupSize > 0 {
            if groupIndex >= len(groups) {
                continue
            }
            if currentGroupSize != groups[groupIndex] {
                continue
            }
            groupIndex++
        }
        if groupIndex == groupCount {
            c += count
        }
    }
    return c
}

func isValid(groups []int, groupIndex, currentGroupSize int, strict bool) bool {
    if groupIndex >= len(groups) {
        return false
    }
    if strict {
        return currentGroupSize == groups[groupIndex]
    } else {
        return currentGroupSize <= groups[groupIndex]
    }
}

func sparseInc(a []int, spanIdx, span, incr int) []int {
    v := sparseGet(a, spanIdx, span)
    newa := sparseSet(a, spanIdx, span, v + incr)

    for _, n := range newa {
        if n != -1 {
            goto OK
        }
    }
    panic("sparseInc: no non-empty elements")

    OK:
    return newa
}

func sparseSet(a []int, spanIdx, span, v int) []int {
    i := spanIdx * maxSpan + span
    if i >= len(a) {
        for j := len(a); j <= i; j++ {
            a = append(a, -1)
        }
    }
    a[i] = v
    return a
}

func sparseGet(a []int, spanIdx, span int) int {
    i := spanIdx * maxSpan + span
    if i >= len(a) || a[i] == - 1 {
        return 0
    } else {
        return a[i]
    }
}
