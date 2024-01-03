package day04

import (
	"testing"

	reader "github.com/orlando-pt/aoc/2023/internal"
    "github.com/orlando-pt/aoc/2023/day04"
)

func TestPart1(t *testing.T) {
	input, err := reader.ReadFileLines("01.txt")

	if err != nil {
		t.Fatalf("Error reading input file: %v\n", err)
	}

	result := day04.Part1(input)

	if result != 13 {
		t.Errorf("Expected 13, got %d", result)
	}
}
