package day15

import (
	"fmt"
	"strings"

	"github.com/orlando-pt/aoc/2024/utils"
)

type Warehouse struct {
	environment       [][]rune
	robotInstructions [][]rune
	robotPosition     utils.Index
	width             int
	height            int
}

func Part1(lines []string) (sum int) {
	warehouse := parseInput(lines)
	utils.Debug(warehouse)

	warehouse.followInstructions()
	return warehouse.score()
}

func Part2(lines []string) (sum int) {
	return
}

func (w Warehouse) score() (score int) {
	for y, line := range w.environment {
		for x, char := range line {
			if char == 'O' {
				score += 100*y + x
			}
		}
	}

	return
}

func (w *Warehouse) followInstructions() {
	for _, instructionSet := range w.robotInstructions {
		for _, instruction := range instructionSet {
			if instruction == '>' {
				if w.environment[w.robotPosition.Y][w.robotPosition.X+1] == '.' {
					w.moveRobot(1, 0)
				} else if w.environment[w.robotPosition.Y][w.robotPosition.X+1] == 'O' {
					if canMove, _ := w.canMoveBoxRightWidthIdx(); canMove {
						w.moveBoxesRight()
						w.moveRobot(1, 0)
					}
				}
			} else if instruction == '<' {
				if w.environment[w.robotPosition.Y][w.robotPosition.X-1] == '.' {
					w.moveRobot(-1, 0)
				} else if w.environment[w.robotPosition.Y][w.robotPosition.X-1] == 'O' {
					if canMove, _ := w.canMoveBoxLeftWidthIdx(); canMove {
						w.moveBoxesLeft()
						w.moveRobot(-1, 0)
					}
				}
			} else if instruction == '^' {
				if w.environment[w.robotPosition.Y-1][w.robotPosition.X] == '.' {
					w.moveRobot(0, -1)
				} else if w.environment[w.robotPosition.Y-1][w.robotPosition.X] == 'O' {
					if canMove, _ := w.canMoveBoxUpWidthIdx(); canMove {
						w.moveBoxesUp()
						w.moveRobot(0, -1)
					}
				}
			} else if instruction == 'v' {
				if w.environment[w.robotPosition.Y+1][w.robotPosition.X] == '.' {
					w.moveRobot(0, 1)
				} else if w.environment[w.robotPosition.Y+1][w.robotPosition.X] == 'O' {
					if canMove, _ := w.canMoveBoxDownWidthIdx(); canMove {
						w.moveBoxesDown()
						w.moveRobot(0, 1)
					}
				}
			}
			w.printWarehouseLayout()
		}
	}
}

func (w Warehouse) printWarehouseLayout() {
	fmt.Println("..........")
	for y, line := range w.environment {
		if y == w.robotPosition.Y {
			line[w.robotPosition.X] = '@'
		}
		fmt.Println(string(line))
	}

	return
}

func (w *Warehouse) moveRobot(x, y int) {
	w.environment[w.robotPosition.Y][w.robotPosition.X] = '.'
	w.robotPosition.X += x
	w.robotPosition.Y += y
}

func (w Warehouse) canMoveBoxRightWidthIdx() (ret bool, idx int) {
	for i := w.robotPosition.X + 1; i < w.width; i++ {
		if w.environment[w.robotPosition.Y][i] == '#' {
			return false, -1
		}
		if w.environment[w.robotPosition.Y][i] == '.' {
			return true, i
		}
	}
	return false, -1
}

func (w *Warehouse) moveBoxesRight() {
	_, availableSpace := w.canMoveBoxRightWidthIdx()
	for i := w.robotPosition.X + 2; i <= availableSpace; i++ {
		w.environment[w.robotPosition.Y][i] = 'O'
	}
}

func (w Warehouse) canMoveBoxLeftWidthIdx() (ret bool, idx int) {
	for i := w.robotPosition.X - 1; i > 0; i-- {
		if w.environment[w.robotPosition.Y][i] == '#' {
			return false, -1
		}
		if w.environment[w.robotPosition.Y][i] == '.' {
			return true, i
		}
	}
	return false, -1
}

func (w *Warehouse) moveBoxesLeft() {
	_, availableSpace := w.canMoveBoxLeftWidthIdx()
	for i := w.robotPosition.X - 2; i >= availableSpace; i-- {
		w.environment[w.robotPosition.Y][i] = 'O'
	}
}

func (w Warehouse) canMoveBoxUpWidthIdx() (ret bool, idx int) {
	for i := w.robotPosition.Y - 1; i > 0; i-- {
		if w.environment[i][w.robotPosition.X] == '#' {
			return false, -1
		}
		if w.environment[i][w.robotPosition.X] == '.' {
			return true, i
		}
	}
	return false, -1
}

func (w *Warehouse) moveBoxesUp() {
	_, availableSpace := w.canMoveBoxUpWidthIdx()
	for i := w.robotPosition.Y - 2; i >= availableSpace; i-- {
		w.environment[i][w.robotPosition.X] = 'O'
	}
}

func (w Warehouse) canMoveBoxDownWidthIdx() (ret bool, idx int) {
	for i := w.robotPosition.Y + 1; i < w.height; i++ {
		if w.environment[i][w.robotPosition.X] == '#' {
			return false, -1
		}
		if w.environment[i][w.robotPosition.X] == '.' {
			return true, i
		}
	}
	return false, -1
}

func (w *Warehouse) moveBoxesDown() {
	_, availableSpace := w.canMoveBoxDownWidthIdx()
	for i := w.robotPosition.Y + 2; i <= availableSpace; i++ {
		w.environment[i][w.robotPosition.X] = 'O'
	}
}

func parseInput(lines []string) (warehouse Warehouse) {
	warehouse.width = len(lines[0])
	env := true
	for y, line := range lines {
		if line == "" {
			env = false
			warehouse.height = y
			continue
		}

		if env {
			if strings.Contains(line, "@") {
				warehouse.robotPosition = utils.Index{X: strings.Index(line, "@"), Y: y}
				line = strings.ReplaceAll(line, "@", ".")
			}
			warehouse.environment = append(warehouse.environment, []rune(line))
		} else {
			warehouse.robotInstructions = append(warehouse.robotInstructions, []rune(line))
		}
	}
	return
}
