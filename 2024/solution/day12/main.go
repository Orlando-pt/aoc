package day12

import "github.com/orlando-pt/aoc/2024/utils"

type VisitedPositions map[utils.Index]bool

var CornerToOrthogonal = map[utils.Index][]utils.Index{
	{X: 1, Y: 1}:   {{X: 1, Y: 0}, {X: 0, Y: 1}},
	{X: -1, Y: 1}:  {{X: -1, Y: 0}, {X: 0, Y: 1}},
	{X: 1, Y: -1}:  {{X: 1, Y: 0}, {X: 0, Y: -1}},
	{X: -1, Y: -1}: {{X: -1, Y: 0}, {X: 0, Y: -1}},
}

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
	visited := make(VisitedPositions)
	maxX, maxY := len(lines[0]), len(lines)

	for y := range lines {
		for x := range lines[y] {
			index := utils.Index{X: x, Y: y}
			if _, ok := visited[index]; !ok {
				area, corners := traversePlantRegionWithCorners(index, lines, maxX, maxY, &visited)
				sum += area * corners
			}
		}
	}
	return
}

func traversePlantRegion(index utils.Index, lines []string, maxX, maxY int, visited *VisitedPositions) (area, perimeter int) {
	if (*visited)[index] {
		return
	}
	(*visited)[index] = true

	plant := lines[index.Y][index.X]
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

func traversePlantRegionWithCorners(index utils.Index, lines []string, maxX, maxY int, visited *VisitedPositions) (area, corners int) {
	if (*visited)[index] {
		return
	}
	(*visited)[index] = true

	plant := lines[index.Y][index.X]
	for _, direction := range utils.Directions {
		neighbour := utils.Index{X: index.X + direction.X, Y: index.Y + direction.Y}

		if checkBonds(neighbour, maxX, maxY) && lines[neighbour.Y][neighbour.X] == plant {
			a, c := traversePlantRegionWithCorners(neighbour, lines, maxX, maxY, visited)
			area += a
			corners += c
		}
	}

	for corner, pair := range CornerToOrthogonal {
		c := utils.Index{X: index.X + corner.X, Y: index.Y + corner.Y}
		i1 := utils.Index{X: index.X + pair[0].X, Y: index.Y + pair[0].Y}
		i2 := utils.Index{X: index.X + pair[1].X, Y: index.Y + pair[1].Y}

		if !match(i1, index, lines, maxX, maxY) && !match(i2, index, lines, maxX, maxY) {
			corners++
		}
		if match(i1, index, lines, maxX, maxY) && match(i2, index, lines, maxX, maxY) && !match(index, c, lines, maxX, maxY) {
			corners++
		}
	}

	return area + 1, corners
}

func match(i1, i2 utils.Index, lines []string, maxX, maxY int) (match bool) {
	if !checkBonds(i1, maxX, maxY) && !checkBonds(i2, maxX, maxY) {
		return true
	} else if checkBonds(i1, maxX, maxY) && checkBonds(i2, maxX, maxY) {
		p1, p2 := lines[i1.Y][i1.X], lines[i2.Y][i2.X]
		return p1 == p2
	}

	return false
}

func checkBonds(index utils.Index, maxX, maxY int) (outOfBounds bool) {
	return index.X >= 0 && index.X < maxX && index.Y >= 0 && index.Y < maxY
}
