package day02

import (
	"testing"

	reader "github.com/orlando-pt/aoc/2024/internal"
	day "github.com/orlando-pt/aoc/2024/solution/day02"
)

func TestPart1(t *testing.T) {
	t.Parallel()

	input, err := reader.ReadFileLines("input.txt")

	if err != nil {
		t.Fatalf("Error reading input file: %v\n", err)
	}

	result := day.Part1(input)

    expected := 2
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	t.Parallel()

	input, err := reader.ReadFileLines("input.txt")

	if err != nil {
		t.Fatalf("Error reading input file: %v\n", err)
	}

	result := day.Part2(input)

    expected := 4
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
