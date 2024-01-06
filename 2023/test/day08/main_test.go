package day08

import (
	"testing"

	reader "github.com/orlando-pt/aoc/2023/internal"
	"github.com/orlando-pt/aoc/2023/solution/day08"
)

func TestPart1(t *testing.T) {
	t.Parallel()

	input, err := reader.ReadFileLines("01.txt")
	input2, err2 := reader.ReadFileLines("02.txt")

	if err != nil || err2 != nil {
		t.Fatalf("Error reading input file: %v\n", err)
	}

	result := day08.Part1(input)
	result2 := day08.Part1(input2)

	if result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}
	if result2 != 6 {
		t.Errorf("Expected 6, got %d", result2)
	}
}

func TestPart2(t *testing.T) {
	t.Parallel()

	input, err := reader.ReadFileLines("03.txt")

	if err != nil {
		t.Fatalf("Error reading input file: %v\n", err)
	}

	result := day08.Part2(input)

	if result != 6 {
		t.Errorf("Expected 6, got %d", result)
	}
}
