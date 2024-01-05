package day07

import (
	"testing"

	reader "github.com/orlando-pt/aoc/2023/internal"
    "github.com/orlando-pt/aoc/2023/solution/day07"
)

func TestPart1(t *testing.T) {
    t.Parallel()

	input, err := reader.ReadFileLines("01.txt")

	if err != nil {
		t.Fatalf("Error reading input file: %v\n", err)
	}

	result := day07.Part1(input)

	if result != 6440 {
		t.Errorf("Expected 6440, got %d", result)
	}
}

func TestPart2(t *testing.T) {
    t.Parallel()

	input, err := reader.ReadFileLines("01.txt")

	if err != nil {
		t.Fatalf("Error reading input file: %v\n", err)
	}

	result := day07.Part2(input)

	if result != 0 {
		t.Errorf("Expected 71503, got %d", result)
	}
}
