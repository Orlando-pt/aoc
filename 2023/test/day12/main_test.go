package day12

import (
	"testing"

	reader "github.com/orlando-pt/aoc/2023/internal"
	day "github.com/orlando-pt/aoc/2023/solution/day12"
)

func TestPart1(t *testing.T) {
	input, err := reader.ReadFileLines("01.txt")

	if err != nil {
		t.Fatalf("Error reading input file: %v\n", err)
	}

	result := day.Part1(input)

    expected := 21
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input, err := reader.ReadFileLines("01.txt")

	if err != nil {
		t.Fatalf("Error reading input file: %v\n", err)
	}

	result := day.Part2(input)

    expected := 525152
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
