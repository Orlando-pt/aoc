package day10

import (
	"testing"

	reader "github.com/orlando-pt/aoc/2023/internal"
	"github.com/orlando-pt/aoc/2023/solution/day10"
)

func TestPart1(t *testing.T) {
	t.Parallel()

	input, err := reader.ReadFileLines("01.txt")

	if err != nil {
		t.Fatalf("Error reading input file: %v\n", err)
	}

	result := day10.Part1(input)

	if result != 8 {
		t.Errorf("Expected 8, got %d", result)
	}
}

// func TestPart2(t *testing.T) {
// 	t.Parallel()
//
// 	input, err := reader.ReadFileLines("01.txt")
//
// 	if err != nil {
// 		t.Fatalf("Error reading input file: %v\n", err)
// 	}
//
// 	result := day10.Part2(input)
//
// 	if result != 2 {
// 		t.Errorf("Expected 2, got %d", result)
// 	}
// }
