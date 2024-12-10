package day04

import (
	"github.com/orlando-pt/aoc/2024/utils"
)

const (
	XMAS = "XMAS"
	SAMX = "SAMX"
	MAS  = "MAS"
	SAM  = "SAM"
)

func Part1(lines []string) (sum int) {
	sum += checkRow(lines)
	sum += checkCol(lines)
	sum += checkDiagonal(lines)
	return
}

func Part2(lines []string) (sum int) {
	for row := 0; row < len(lines)-2; row++ {
		for col := 0; col < len(lines[0])-2; col++ {
			diagonal1 := []byte{lines[row][col], lines[row+1][col+1], lines[row+2][col+2]}
			diagonal2 := []byte{lines[row+2][col], lines[row+1][col+1], lines[row][col+2]}

			if str1, str2 := string(diagonal1), string(diagonal2); (str1 == MAS || str1 == SAM) && (str2 == MAS || str2 == SAM) {
				sum++
			}
		}
	}
	utils.Debug("X diagonal sum: ", sum)
	return
}

func checkRow(lines []string) (sum int) {
	for _, line := range lines {
		for i := 0; i < len(lines[0])-3; i++ {
			if line[i:i+4] == XMAS || line[i:i+4] == SAMX {
				sum++
			}
		}
	}
	utils.Debug("Row sum: ", sum)
	return
}

func checkCol(lines []string) int {
	return checkRow(mapToVerticalLines(lines))
}

func checkDiagonal(lines []string) (sum int) {
	for row := 0; row < len(lines)-3; row++ {
		for col := 0; col < len(lines[0])-3; col++ {
			diagonal1 := []byte{lines[row][col], lines[row+1][col+1], lines[row+2][col+2], lines[row+3][col+3]}
			diagonal2 := []byte{lines[row+3][col], lines[row+2][col+1], lines[row+1][col+2], lines[row][col+3]}

			if str := string(diagonal1); str == XMAS || str == SAMX {
				sum++
			}
			if str := string(diagonal2); str == XMAS || str == SAMX {
				sum++
			}
		}
	}
	utils.Debug("Diagonal sum: ", sum)
	return
}

func mapToVerticalLines(lines []string) (reversed []string) {
	for i := 0; i < len(lines[0]); i++ {
		line := make([]byte, len(lines))
		for j := 0; j < len(lines); j++ {
			line[j] = lines[j][i]
		}
		reversed = append(reversed, string(line))
	}

	utils.Debug("Vertical lines: ", reversed)
	return
}
