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
