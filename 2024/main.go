package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/orlando-pt/aoc/2024/internal"
	"github.com/orlando-pt/aoc/2024/solution/day01"
	"github.com/orlando-pt/aoc/2024/solution/day02"
	"github.com/orlando-pt/aoc/2024/solution/day03"
	"github.com/orlando-pt/aoc/2024/solution/day04"
	"github.com/orlando-pt/aoc/2024/solution/day05"
	"github.com/orlando-pt/aoc/2024/solution/day06"
	"github.com/orlando-pt/aoc/2024/solution/day07"
	"github.com/orlando-pt/aoc/2024/solution/day08"
	"github.com/orlando-pt/aoc/2024/solution/day09"
	"github.com/orlando-pt/aoc/2024/solution/day10"
	"github.com/orlando-pt/aoc/2024/solution/day11"
	"github.com/orlando-pt/aoc/2024/solution/day12"
    "github.com/orlando-pt/aoc/2024/solution/day13"
    "github.com/orlando-pt/aoc/2024/solution/day14"
    "github.com/orlando-pt/aoc/2024/solution/day15"
	"github.com/orlando-pt/aoc/2024/utils"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatalf("You need to provide both the day and the path to input!\n\t go run . <day> [debug (log level)]\n")
	}

	day := fmt.Sprintf("%02d", utils.StrToInt(os.Args[1]))
	inputPath := "input/" + day + ".txt"

	input, err := reader.ReadFileLines(inputPath)
	if err != nil {
		log.Fatalf("Error reading input file: %v\n", err)
	}

	setMainProgramLogger()
	switch day {
	case "01":
		runDay(day, day01.Part1, day01.Part2, input)
	case "02":
		runDay(day, day02.Part1, day02.Part2, input)
	case "03":
		runDay(day, day03.Part1, day03.Part2, input)
	case "04":
		runDay(day, day04.Part1, day04.Part2, input)
    case "05":
        runDay(day, day05.Part1, day05.Part2, input)
    case "06":
        runDay(day, day06.Part1, day06.Part2, input)
    case "07":
        runDay(day, day07.Part1, day07.Part2, input)
    case "08":
        runDay(day, day08.Part1, day08.Part2, input)
    case "09":
        runDay(day, day09.Part1, day09.Part2, input)
    case "10":
        runDay(day, day10.Part1, day10.Part2, input)
    case "11":
        runDay(day, day11.Part1, day11.Part2, input)
    case "12":
        runDay(day, day12.Part1, day12.Part2, input)
    case "13":
        runDay(day, day13.Part1, day13.Part2, input)
    case "14":
        runDay(day, day14.Part1, day14.Part2, input)
    case "15":
        runDay(day, day15.Part1, day15.Part2, input)
	default:
		log.Fatalf("Day %s not implemented yet! Or doesn't exist.\n", day)
	}
}

func runDay(day string, part1, part2 func([]string) int, input []string) {
	fmt.Printf("===== %s =====\n", day)
	if part1 != nil {
		start := time.Now()
		fmt.Printf("Part 1: %d\n", part1(input))
		elapsed := time.Since(start)
		fmt.Printf("Part 1 took %s\n", elapsed)
	}

	if part2 != nil {
		start := time.Now()
		fmt.Printf("Part 2: %d\n", part2(input))
		elapsed := time.Since(start)
		fmt.Printf("Part 2 took %s\n", elapsed)
	}
	fmt.Println()
}

func setMainProgramLogger() {

	var programLevel = new(slog.LevelVar)
	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: programLevel})

	slog.SetDefault(slog.New(h))

	if len(os.Args) == 3 && os.Args[2] == "debug" {
		programLevel.Set(slog.LevelDebug)
	} else {
		programLevel.Set(slog.LevelError)
	}
}
