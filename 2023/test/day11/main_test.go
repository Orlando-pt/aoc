package day11

import (
	"testing"

	reader "github.com/orlando-pt/aoc/2023/internal"
	"github.com/orlando-pt/aoc/2023/solution/day11"
)

func TestPart1(t *testing.T) {
	input, err := reader.ReadFileLines("01.txt")

	if err != nil {
		t.Fatalf("Error reading input file: %v\n", err)
	}

	result := day11.Part1(input)

	if result != 374 {
		t.Errorf("Expected 374, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	input, err := reader.ReadFileLines("01.txt")

	if err != nil {
		t.Fatalf("Error reading input file: %v\n", err)
	}

	result := day11.Part2(input)

	if result != 1030 {
		t.Errorf("Expected 1030, got %d", result)
	}
}