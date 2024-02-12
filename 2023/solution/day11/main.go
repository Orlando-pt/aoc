package day11

import "github.com/orlando-pt/aoc/2023/utils"

var univ []int
var uw, uh int
var rows, cols [][]int

func Part1(lines []string) (sum int) {
	univ = []int{}
	return computeDistances(lines, 1)
}

// thanks to:
// https://github.com/ColasNahaboo/advent-of-code-my-solutions/blob/main/go/2023/days/d11/d11.go
func Part2(lines []string) int {
	univ = []int{}
	return computeDistances(lines, 1000000-1)
}

func computeDistances(lines []string, expansion int) (sum int) {
	parse(lines)
	expandUniv(expansion)
	// unordered pairs: we take all (g1, g2) with g2 < g1
	for g1 := 0; g1 < len(univ)-1; g1++ {
		for g2 := g1 + 1; g2 < len(univ); g2++ {
			sum += distance(g1, g2)
		}
	}
	return
}

func parse(lines []string) {
	uw = len(lines[0])
	uh = len(lines)
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			if lines[y][x] == '#' {
				galaxy := x + y*uw
				univ = append(univ, galaxy)
			}
		}
	}
}

func expandUniv(expansion int) {
	// compute rows and cols
	rows = make([][]int, uh)
	for y := range rows {
		rows[y] = []int{}
	}

	cols = make([][]int, uw)
	for x := range cols {
		cols[x] = []int{}
	}
	for gid, gal := range univ {
		rows[gal/uw] = append(rows[gal/uw], gid)
		cols[gal%uw] = append(cols[gal%uw], gid)
	}

	// delay expand for after we computed the new uw
	// lists of amounts to move down and right for all gids
	tmd := make([]int, len(univ), len(univ))
	tmr := make([]int, len(univ), len(univ))
	ouw := uw
	for y, exp := 0, 0; y < len(rows); y++ {
		if len(rows[y]) == 0 {
			exp += expansion // increase the expansion
			uh += expansion
			continue
		}
		if exp == 0 {
			continue
		}
		for _, gid := range rows[y] { // shift later, when new uw is known
			tmd[gid] = exp
		}
	}

	// expand cols rightwards
	for x, exp := 0, 0; x < len(cols); x++ {
		if len(cols[x]) == 0 {
			exp += expansion // increase the expansion
			uw += expansion
			continue
		}
		if exp == 0 {
			continue
		}
		for _, gid := range cols[x] { // shift right later
			tmr[gid] = exp
		}
	}

	// perform the moves
	for gid, pos := range univ {
		ox := pos % ouw
		oy := pos / ouw
		x := ox + tmr[gid]
		y := oy + tmd[gid]
		univ[gid] = x + y*uw
	}
}

func distance(g1, g2 int) int {
	return utils.IntAbs(univ[g2]%uw-univ[g1]%uw) + utils.IntAbs(univ[g2]/uw-univ[g1]/uw)
}
