package day10

const (
	ns     = '|'
	ew     = '-'
	ne     = 'L'
	nw     = 'J'
	sw     = '7'
	se     = 'F'
	ground = '.'
	start  = 'S'
)

func Part1(lines []string) int {
	start := findStart(lines)

	surroundings := start.findPossibleSurroundings(lines)

	var paths [][]coordinate
	for _, surrounding := range surroundings {
		paths = append(paths, surrounding.findLoop(lines, start))

	}

	// find duplicates
	steps := 0
	for i := 0; i < len(paths); i++ {
		for j := i + 1; j < len(paths); j++ {
			if len(paths[i]) == len(paths[j]) {
				steps = len(paths[i])
				break
			}
		}
		if steps != 0 {
			break
		}
	}
	return steps / 2
}

// Thanks to:
// https://github.com/ColasNahaboo/advent-of-code-my-solutions/blob/main/go/2023/days/d10/d10.go
func Part2(lines []string) int {
	start := findStart(lines)

	surroundings := start.findPossibleSurroundings(lines)

	var paths [][]coordinate
	for _, surrounding := range surroundings {
		paths = append(paths, surrounding.findLoop(lines, start))

	}

	// find duplicates
	var path []coordinate
	for i := 0; i < len(paths); i++ {
		for j := i + 1; j < len(paths); j++ {
			if len(paths[i]) == len(paths[j]) {
				path = paths[i]
				break
			}
		}
		if len(path) != 0 {
			break
		}
	}

	if len(path) == 0 {
		panic("no path found")
	}

    // replace the loop with a start pipe
    line := []byte(lines[start.y])
    line[start.x] = findStartPipe(path[0], path[len(path)-2])
    lines[start.y] = string(line)

	return calculateContainedTiles(path, lines)
}

func findStartPipe(c1, c2 coordinate) byte {
    if c1.x == c2.x {
        return ns
    }
    if c1.y == c2.y {
        return ew
    }
    if c1.x < c2.x {
        if c2.y < c1.y {
            return nw
        }
        return sw
    }
    if c2.y < c1.y {
        return ne
    }
    return se
}

func calculateContainedTiles(path []coordinate, field []string) int {
	// mark all places occupied by the loop
	dim := len(field)
	board := make([]bool, dim*dim, dim*dim)
	for _, c := range path {
		board[c.y*dim+c.x] = true
	}

	insiders := 0
	for p := range board {
		if board[p] {
			continue
		}

		crosses := 0
		for i := p; i >= 0; i -= dim {
			if board[i] {
				crosses += isCrossing(i, field)
			}
		}
		if crosses%2 == 1 {
			insiders++
		}
	}

	return insiders
}

func isCrossing(pos int, field []string) int {
	dim := len(field)
	x := pos % dim
	y := pos / dim

	switch field[y][x] {
	case ew, ne, se:
		return 1
	default:
		return 0
	}
}

func (c coordinate) findPossibleSurroundings(lines []string) []coordinate {
	surroundings := []coordinate{
		{c.x, c.y - 1},
		{c.x, c.y + 1},
		{c.x - 1, c.y},
		{c.x + 1, c.y},
	}

	var possibleSurroundings []coordinate
	for _, surrounding := range surroundings {
		if surrounding.isValid(len(lines)) && lines[surrounding.x][surrounding.y] != ground {
			possibleSurroundings = append(possibleSurroundings, surrounding)
		}
	}

	return possibleSurroundings
}

func (c coordinate) findLoop(lines []string, last coordinate) []coordinate {
	if !c.isValid(len(lines)) || lines[c.y][c.x] == ground {
		return []coordinate{}
	}

	if lines[c.y][c.x] == start {
		return []coordinate{c}
	}

	return append([]coordinate{c}, c.nextMove(last, rune(lines[c.y][c.x])).findLoop(lines, c)...)
}

func findStart(lines []string) coordinate {
	for y, line := range lines {
		for x, char := range line {
			if char == start {
				return coordinate{x, y}
			}
		}
	}

	return coordinate{-1, -1}
}

func (c coordinate) isValid(max int) bool {
	return c.x >= 0 && c.y >= 0 && c.x < max && c.y < max
}

func (c coordinate) nextMove(last coordinate, direction rune) coordinate {
	switch direction {
	case ns:
		if last.y < c.y {
			c.y++
		} else {
			c.y--
		}
	case ew:
		if last.x < c.x {
			c.x++
		} else {
			c.x--
		}
	case ne:
		if last.y < c.y {
			c.x++
		} else {
			c.y--
		}
	case nw:
		if last.y < c.y {
			c.x--
		} else {
			c.y--
		}
	case sw:
		if last.y > c.y {
			c.x--
		} else {
			c.y++
		}
	case se:
		if last.y > c.y {
			c.x++
		} else {
			c.y++
		}
	}

	return c
}

type coordinate struct {
	x int
	y int
}
