package day10

// Source from:
// https://github.com/ColasNahaboo/advent-of-code-my-solutions/blob/main/go/2023/days/d10/d10.go

import "fmt"

var grid []byte      // the grid, gw x gh, the raw input, with S replaced
var gw, gh, area int // grid size
var startHelp int    // start position, pos of 'S'
var dirs [4]int      // the 4 directions as position offsets
const (
	UP      = 1
	RIGHT   = 2
	DOWN    = 4
	LEFT    = 8
	verbose = false
)

func Part1Help(lines []string) int {
	path := findLoop(lines)
	return (len(path) + 1) / 2
}

func Part2Help(lines []string) int {
	path := findLoop(lines)
	// mark all the places occupied by the loop
	board := make([]bool, area, area)
	inside := make([]bool, area, area) // DEBUG
	for _, pos := range path {
		board[pos] = true
	}
	// find places that can go to the top border crossing the loop an odd times
	insiders := 0
	for p := range board {
		if board[p] { // on the loop itself means not inside
			continue
		}
		crosses := 0
		for i := p; i >= 0; i -= gw {
			// we are moving up, so moving along '|' does not count as a crossing
			if board[i] {
				crosses += isCrossingHelp(i)
			}
		}
		if crosses%2 == 1 {
			insiders++
			inside[p] = true
		}
	}
	VPf("Start pipe is a: \"%s\"\n", string(grid[startHelp]))
	VPboard(board, inside)
	return insiders
}

// we count the pipes with only an horizontal part connecting to the right
// otherwise 7 above a L would count as a double crossing
func isCrossingHelp(pos int) int {
	switch grid[pos] {
	case '-', 'L', 'F':
		return 1
	default:
		return 0
	}
}

//////////// Common Parts code

// No real parsing, we just concatenate all the input strings (the rows) into
// one single byte array (a slice, actually), the grid of size (gw x gh).
// A position (x, y) is thus just an integer index x+gw*y in this array

func findLoop(lines []string) (lpath []int) {
	gw = len(lines[0])
	gh = len(lines)
	area = gw * gh
	grid = make([]byte, area, area)
	for y, line := range lines {
		for x, c := range line {
			grid[y*gw+x] = byte(c)
		}
	}
	dirs = [4]int{-gw, +1, gw, -1}
	startHelp = startPos()
	VPf("Grid %d x %d, start at %d\n", gw, gh, startHelp)
	ns := neighbours(startHelp)
	for _, n := range ns {
		path := followLoop(startHelp, n)
		if len(path) > len(lpath) {
			lpath = path
		}
	}
	grid[startHelp] = pipeConnecting(startHelp, lpath[0], lpath[len(lpath)-2])
	return
}

func startPos() int {
	for p, c := range grid {
		if c == 'S' {
			return p
		}
	}
	panic("S not found")
}

// follows a (potential) loop completely
func followLoop(pos, next int) (path []int) {
	for {
		path = append(path, next)
		if next == startHelp { // we went back to start
			return
		}
		from := pos
		pos = next
		next = nextStep(pos, from)
		if next == -1 { // not a loop!
			path = []int{}
			return
		}
	}
}

// follows the loop for one step from "pos", not going back via "from"
// we suppose we have one and only one possibility
// returns -1 if no neigbours can be found
func nextStep(pos, from int) (next int) {
	ns := neighbours(pos)
	for _, n := range ns {
		if n == from {
			continue
		}
		next = n
		return
	}
	next = -1
	return
}

// List the possible neighbours. S has all 4 dirs, other pipes only 2.
func neighbours(pos int) []int {
	switch grid[pos] {
	case '|':
		return validPair(pos-gw, pos+gw)
	case '-':
		return validPair(pos-1, pos+1)
	case 'L':
		return validPair(pos-gw, pos+1)
	case 'J':
		return validPair(pos-gw, pos-1)
	case '7':
		return validPair(pos-1, pos+gw)
	case 'F':
		return validPair(pos+1, pos+gw)
	case 'S': // S is special, be tolerant
		ns := []int{}
		for _, d := range dirs {
			p := pos + d
			if d >= 0 && d < area {
				ns = append(ns, p)
			}
		}
		return ns
	default:
		return []int{}
	}
}

func validPair(i, j int) (pair []int) {
	if i >= 0 && i < area && j >= 0 && j < area {
		pair = []int{i, j}
	}
	return
}

// return the symbol of the pipe at p that connects neighbours to and from
func pipeConnecting(p int, ns ...int) byte {
	sides := 0
	for _, n := range ns {
		switch n - p {
		case -gw:
			sides |= UP
		case 1:
			sides |= RIGHT
		case gw:
			sides |= DOWN
		case -1:
			sides |= LEFT
		default:
			panic(fmt.Sprintf("%d is not a neighbour of %d", n, p))
		}
	}
	switch sides {
	case 5:
		return '|'
	case 10:
		return '-'
	case 3:
		return 'L'
	case 9:
		return 'J'
	case 12:
		return '7'
	case 6:
		return 'F'
	default:
		panic(fmt.Sprintf("Cannot connect %d to %v", p, ns))
	}
}

//////////// Debug code

func VPboard(board, inside []bool) {
	if !verbose {
		return
	}
	for pos, inside := range inside {
		if inside {
			fmt.Print("#")
		} else if board[pos] {
			fmt.Print(string(grid[pos]))
		} else {
			fmt.Print(".")
		}
		if pos%gw == gw-1 {
			fmt.Println()
		}
	}
}

// VPf = fmt.Printf
func VPf(f string, v ...interface{}) {
	if verbose {
		fmt.Printf(f, v...)
	}
}
