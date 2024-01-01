package day01

import (
	"testing"

	"github.com/orlando-pt/aoc/2023/day01"
	"github.com/orlando-pt/aoc/2023/internal"
)

func TestPart1(t *testing.T) {
	input, err := reader.ReadFileLines("01.txt")

	if err != nil {
		t.Fatalf("Error reading input file: %v\n", err)
	}

	result := day01.Part1(input)

	if result != 142 {
		t.Errorf("Expected 142, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	input, err := reader.ReadFileLines("02.txt")

	if err != nil {
		t.Fatalf("Error reading input file: %v\n", err)
	}

	result := day01.Part2(input)

	if result != 281 {
		t.Errorf("Expected 281, got %d", result)
	}
}
