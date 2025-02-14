package day12

import (
	"testing"

	reader "github.com/orlando-pt/aoc/2024/internal"
	day "github.com/orlando-pt/aoc/2024/solution/day12"
	"github.com/orlando-pt/aoc/2024/utils"
)

func TestPart1(t *testing.T) {
	t.Parallel()

	input, err := reader.ReadFileLines("input.txt")

	if err != nil {
		t.Fatalf("Error reading input file: %v\n", err)
	}

    utils.SetTestLogger()
	result := day.Part1(input)

    expected := 1930
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

    utils.SetTestLogger()
	result := day.Part2(input)

    expected := 1206
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
