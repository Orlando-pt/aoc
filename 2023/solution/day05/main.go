package day05

import (
	"math"
	"strconv"
	"strings"
	"unicode"

	"github.com/orlando-pt/aoc/2023/utils"
)

func Part1(lines []string) int {
	parseInput := parseInput(lines)

	minLocation := math.MaxInt32

	for _, seed := range parseInput.seeds {

		seed = getNext(seed, parseInput.seed2soil)
		seed = getNext(seed, parseInput.soil2fertilizer)
		seed = getNext(seed, parseInput.fertilizer2water)
		seed = getNext(seed, parseInput.water2light)
		seed = getNext(seed, parseInput.light2temperature)
		seed = getNext(seed, parseInput.temperature2humidity)
		seed = getNext(seed, parseInput.humidity2location)

		if seed < minLocation {
			minLocation = seed
		}
	}

	return minLocation
}

func Part2(lines []string) int {
	parseInput := parseInput(lines)

	minLocation := math.MaxInt32


	for i := 0; i < len(parseInput.seeds); i += 2 {
		for lower := parseInput.seeds[i]; lower < parseInput.seeds[i]+parseInput.seeds[i+1]; lower++ {
            seed := lower

			seed = getNext(seed, parseInput.seed2soil)
			seed = getNext(seed, parseInput.soil2fertilizer)
			seed = getNext(seed, parseInput.fertilizer2water)
			seed = getNext(seed, parseInput.water2light)
			seed = getNext(seed, parseInput.light2temperature)
			seed = getNext(seed, parseInput.temperature2humidity)
			seed = getNext(seed, parseInput.humidity2location)

			if seed < minLocation {
				minLocation = seed
			}
		}
	}

	return minLocation
}

type SourceDestination struct {
	source      int
	destination int
	interval    int
}

type InputData struct {
	seeds                []int
	seed2soil            []SourceDestination
	soil2fertilizer      []SourceDestination
	fertilizer2water     []SourceDestination
	water2light          []SourceDestination
	light2temperature    []SourceDestination
	temperature2humidity []SourceDestination
	humidity2location    []SourceDestination
}

func getNext(seed int, attribution []SourceDestination) int {
	for _, attr := range attribution {
		if seed >= attr.source && seed < attr.source+attr.interval {
			seed = attr.destination + (seed - attr.source)
			break
		}
	}

	return seed
}

func parseInput(lines []string) InputData {
	ret := InputData{}

	ret.seeds = parseSeeds(lines[0])

	currentAttribution := -1
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}

		if unicode.IsLetter(rune(line[0])) {
			currentAttribution++
			continue
		}

		switch currentAttribution {
		case 0:
			ret.seed2soil = append(ret.seed2soil, parseSourceDestination(line))
		case 1:
			ret.soil2fertilizer = append(ret.soil2fertilizer, parseSourceDestination(line))
		case 2:
			ret.fertilizer2water = append(ret.fertilizer2water, parseSourceDestination(line))
		case 3:
			ret.water2light = append(ret.water2light, parseSourceDestination(line))
		case 4:
			ret.light2temperature = append(ret.light2temperature, parseSourceDestination(line))
		case 5:
			ret.temperature2humidity = append(ret.temperature2humidity, parseSourceDestination(line))
		case 6:
			ret.humidity2location = append(ret.humidity2location, parseSourceDestination(line))
		}
	}

	return ret
}

func parseSourceDestination(line string) SourceDestination {
	ret := SourceDestination{}

	parts := strings.Split(strings.TrimSpace(line), " ")
	ret.source = utils.StrToInt(parts[1])
	ret.destination = utils.StrToInt(parts[0])
	ret.interval = utils.StrToInt(parts[2])

	return ret
}

func parseSeeds(line string) []int {
	ret := []int{}

	seeds := strings.TrimSpace(strings.Split(line, ":")[1])
	for _, seed := range strings.Split(seeds, " ") {
		number, err := strconv.Atoi(seed)

		if err != nil {
			panic(err)
		}

		ret = append(ret, number)
	}

	return ret
}
