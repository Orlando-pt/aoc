package day03

import (
	"testing"

	"github.com/orlando-pt/aoc/2023/solution/day03"
	reader "github.com/orlando-pt/aoc/2023/internal"
)

func TestPart1(t *testing.T) {
	input, err := reader.ReadFileLines("01.txt")

	if err != nil {
		t.Fatalf("Error reading input file: %v\n", err)
	}

	result := day03.Part1(input)

	if result != 4361 {
		t.Errorf("Expected 4361, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	input, err := reader.ReadFileLines("01.txt")

	if err != nil {
		t.Fatalf("Error reading input file: %v\n", err)
	}

	result := day03.Part2(input)

	if result != 467835 {
		t.Errorf("Expected 467835, got %d", result)
	}
}
