package day06

import (
	"strings"

	"github.com/orlando-pt/aoc/2023/utils"
)

func Part1(lines []string) int {
	races := parseLines(lines)

	var racesPossibleWins []int
	for _, race := range races {
		racesPossibleWins = append(racesPossibleWins, getNrPossibleWins(race))
	}

	return mul(racesPossibleWins)
}

func Part2(lines []string) int {
	race := parseLinesPart2(lines)
	return getNrPossibleWins(race)
}

func parseLinesPart2(lines []string) Race {

	timeString := strings.Split(lines[0], ":")[1]
	time := strings.ReplaceAll(timeString, " ", "")

	recordString := strings.Split(lines[1], ":")[1]
	record := strings.ReplaceAll(recordString, " ", "")

	return Race{time: utils.StrToInt(time), record: utils.StrToInt(record)}
}

type Race struct {
	time   int
	record int
}

func getNrPossibleWins(race Race) int {
	wins := 0

	for pressing := 1; pressing <= race.time; pressing++ {
		remainingPressing := race.time - pressing

		distance := pressing * remainingPressing

		if distance > race.record {
			wins++
		}
	}

	return wins
}

func mul(list []int) int {
	mul := 1

	for _, i := range list {
		mul *= i
	}

	return mul
}

func parseLines(lines []string) []Race {
	var races []Race

	timesString := strings.TrimSpace(strings.Split(lines[0], ":")[1])
	times := strings.Split(timesString, " ")

	distancesString := strings.TrimSpace(strings.Split(lines[1], ":")[1])
	distances := strings.Split(distancesString, " ")

	times = removeEmpty(times)
	distances = removeEmpty(distances)

	for i := 0; i < len(times); i++ {
		time := utils.StrToInt(times[i])
		record := utils.StrToInt(distances[i])

		races = append(races, Race{time: time, record: record})
	}

	return races
}

func removeEmpty(list []string) []string {
	var result []string

	for _, s := range list {
		if s != "" {
			result = append(result, s)
		}
	}

	return result
}
