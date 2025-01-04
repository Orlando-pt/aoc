package day06

import "github.com/orlando-pt/aoc/2024/utils"

type Index struct {
	r, c int
}

var DIRECTIONS = [...]Index{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func Part1(lines []string) (sum int) {
	return len(getVisitedIndices(lines))
}

func Part2(lines []string) (sum int) {
	visited := getVisitedIndices(lines)
	startIndex := findGuard(lines)

	for _, idx := range visited {
		if (idx == startIndex) || lines[idx.r][idx.c] == '#' {
			continue
		}

		lines[idx.r] = lines[idx.r][:idx.c] + "#" + lines[idx.r][idx.c+1:]
		if hasCycle(lines, startIndex) {
			sum++
		}
		lines[idx.r] = lines[idx.r][:idx.c] + "." + lines[idx.r][idx.c+1:]
	}

	return
}

func getVisitedIndices(lines []string) (indices []Index) {
	maxRow := len(lines)
	maxCol := len(lines[0])

	visited := make(map[Index]bool)
	facing := 0

	at := findGuard(lines)
	utils.Debug("Guard at:", at)

	for isValidIndex(at, maxRow, maxCol) {
		if lines[at.r][at.c] == '#' {
			at.r -= DIRECTIONS[facing].r
			at.c -= DIRECTIONS[facing].c
			facing = (facing + 1) % 4
			continue
		}
		visited[at] = true
		at.r += DIRECTIONS[facing].r
		at.c += DIRECTIONS[facing].c
	}

	for k := range visited {
		indices = append(indices, k)
	}
	return
}

func hasCycle(lines []string, startIndex Index) bool {
	visited := make(map[Index]Index)
	maxRow := len(lines)
	maxCol := len(lines[0])

	at := startIndex
	facing := 0

	for isValidIndex(at, maxRow, maxCol) {
		if visited[at] == DIRECTIONS[facing] {
			return true
		}

		visited[at] = DIRECTIONS[facing]

		if lines[at.r][at.c] == '#' {
			at.r -= DIRECTIONS[facing].r
			at.c -= DIRECTIONS[facing].c
			facing = (facing + 1) % 4
			continue
		}

		lines[at.r] = lines[at.r][:at.c] + "X" + lines[at.r][at.c+1:]
		at.r += DIRECTIONS[facing].r
		at.c += DIRECTIONS[facing].c
	}

	return false
}

func isValidIndex(idx Index, maxRow, maxCol int) bool {
	return idx.r >= 0 && idx.c >= 0 && idx.r < maxRow && idx.c < maxCol
}

func findGuard(lines []string) (guard Index) {
	for i, line := range lines {
		for j, c := range line {
			if c == '^' {
				return Index{i, j}
			}
		}
	}
	return
}
