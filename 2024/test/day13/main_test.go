package day13

import (
	"testing"

	reader "github.com/orlando-pt/aoc/2024/internal"
	day "github.com/orlando-pt/aoc/2024/solution/day13"
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

    expected := 480
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

    expected := 875318608908
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
