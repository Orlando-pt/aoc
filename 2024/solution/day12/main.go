package day12

import "github.com/orlando-pt/aoc/2024/utils"

type VisitedPositions map[utils.Index]bool

func Part1(lines []string) (sum int) {
	visited := make(VisitedPositions)
	maxX, maxY := len(lines[0]), len(lines)

	for y, line := range lines {
		for x := range line {
			index := utils.Index{X: x, Y: y}
			if _, ok := visited[index]; !ok {
				area, perimeter := traversePlantRegion(index, lines, maxX, maxY, &visited)
				sum += area * perimeter
			}
		}
	}
	return
}

func Part2(lines []string) (sum int) {
	return
}

func traversePlantRegion(index utils.Index, lines []string, maxX, maxY int, visited *VisitedPositions) (area, perimeter int) {
	plant := lines[index.Y][index.X]

	if (*visited)[index] {
		return
	}

	(*visited)[index] = true

	neightbours := 0
	for _, direction := range utils.Directions {
		neighbour := utils.Index{X: index.X + direction.X, Y: index.Y + direction.Y}

		if checkBonds(neighbour, maxX, maxY) && lines[neighbour.Y][neighbour.X] == plant {
			neightbours++
			a, p := traversePlantRegion(neighbour, lines, maxX, maxY, visited)
			area += a
			perimeter += p
		}
	}

	return area + 1, perimeter + 4 - neightbours
}

func checkBonds(index utils.Index, maxX, maxY int) (outOfBounds bool) {
	return index.X >= 0 && index.X < maxX && index.Y >= 0 && index.Y < maxY
}
