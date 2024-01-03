package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/orlando-pt/aoc/2023/solution/day01"
	"github.com/orlando-pt/aoc/2023/solution/day02"
	"github.com/orlando-pt/aoc/2023/solution/day03"
	"github.com/orlando-pt/aoc/2023/solution/day04"
	"github.com/orlando-pt/aoc/2023/internal"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatalf("You need to provide both the day and the path to input!\n\t go run . <day>\n")
	}

	day := os.Args[1]
	inputPath := "input/" + day + ".txt"

	input, err := reader.ReadFileLines(inputPath)
	if err != nil {
		log.Fatalf("Error reading input file: %v\n", err)
	}

	switch day {
	case "01":
		runDay(day, day01.Part1, day01.Part2, input)
	case "02":
		runDay(day, day02.Part1, day02.Part2, input)
	case "03":
		runDay(day, day03.Part1, day03.Part2, input)
    case "04":
        runDay(day, day04.Part1, day04.Part2, input)
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
