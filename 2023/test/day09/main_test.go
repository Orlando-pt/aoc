package day09

import (
	"testing"

	reader "github.com/orlando-pt/aoc/2023/internal"
	"github.com/orlando-pt/aoc/2023/solution/day09"
)

func TestPart1(t *testing.T) {
	t.Parallel()

	input, err := reader.ReadFileLines("01.txt")

	if err != nil {
		t.Fatalf("Error reading input file: %v\n", err)
	}

	result := day09.Part1(input)

	if result != 114 {
		t.Errorf("Expected 114, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	t.Parallel()

	input, err := reader.ReadFileLines("01.txt")

	if err != nil {
		t.Fatalf("Error reading input file: %v\n", err)
	}

	result := day09.Part2(input)

	if result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}
}
