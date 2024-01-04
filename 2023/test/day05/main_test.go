package day05

import (
	"testing"

	reader "github.com/orlando-pt/aoc/2023/internal"
    "github.com/orlando-pt/aoc/2023/solution/day05"
)

func TestPart1(t *testing.T) {
	input, err := reader.ReadFileLines("01.txt")

	if err != nil {
		t.Fatalf("Error reading input file: %v\n", err)
	}

	result := day05.Part1(input)

	if result != 35 {
		t.Errorf("Expected 35, got %d", result)
	}
}

// func TestPart2(t *testing.T) {
// 	input, err := reader.ReadFileLines("01.txt")
//
// 	if err != nil {
// 		t.Fatalf("Error reading input file: %v\n", err)
// 	}
//
// 	result := day05.Part2(input)
//
// 	if result != 30 {
// 		t.Errorf("Expected 30, got %d", result)
// 	}
// }
