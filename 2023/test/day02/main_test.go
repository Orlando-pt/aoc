package day02

import (
	"testing"

	"github.com/orlando-pt/aoc/2023/solution/day02"
	reader "github.com/orlando-pt/aoc/2023/internal"
)

func TestPart1(t *testing.T) {
	input, err := reader.ReadFileLines("01.txt")

	if err != nil {
		t.Fatalf("Error reading input file: %v\n", err)
	}

	result := day02.Part1(input)

	if result != 8 {
		t.Errorf("Expected 8, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	input, err := reader.ReadFileLines("01.txt")

	if err != nil {
		t.Fatalf("Error reading input file: %v\n", err)
	}

	result := day02.Part2(input)

	if result != 2286 {
		t.Errorf("Expected 2286, got %d", result)
	}
}

