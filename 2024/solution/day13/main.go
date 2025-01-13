package day13

import (
	"regexp"

	"github.com/orlando-pt/aoc/2024/utils"
)

const (
	COST_A          = 3
	COST_B          = 1
	CONVERSION_RATE = 10000000000000
)

type Machine struct {
	buttonA utils.Index
	buttonB utils.Index
	prize   utils.Index
}

func Part1(lines []string) (sum int) {
	machines := parseInput(lines, false)

	for _, machine := range machines {
		a, b := applyCramerTheorem(machine)
		if a != -1 {
			sum += cost(a, b)
		}
	}
	return
}

func Part2(lines []string) (sum int) {
	machines := parseInput(lines, true)

	for _, machine := range machines {
		a, b := applyCramerTheorem(machine)
		if a != -1 {
			sum += cost(a, b)
		}
	}

	return
}

func applyCramerTheorem(machine Machine) (x, y int) {
	det := machine.buttonA.X*machine.buttonB.Y - machine.buttonA.Y*machine.buttonB.X
	if det == 0 {
		return -1, -1
	}

	detX := machine.prize.X*machine.buttonB.Y - machine.prize.Y*machine.buttonB.X
	detY := machine.buttonA.X*machine.prize.Y - machine.buttonA.Y*machine.prize.X
	if detX%det != 0 || detY%det != 0 {
		return -1, -1
	}

	utils.Debug("Cramer Theorem: ", det, detX, detY)
	return detX / det, detY / det
}

func cost(a, b int) int {
	return a*COST_A + b*COST_B
}

func parseInput(lines []string, withConversion bool) (machines []Machine) {
	for i := 0; i < len(lines); i += 4 {
		decimals := regexp.MustCompile(`\d+`)
		a := decimals.FindAllString(lines[i], -1)
		b := decimals.FindAllString(lines[i+1], -1)
		p := decimals.FindAllString(lines[i+2], -1)
		utils.Debug(a, b, p)

		machines = append(machines, Machine{
			buttonA: utils.Index{X: utils.StrToInt(a[0]), Y: utils.StrToInt(a[1])},
			buttonB: utils.Index{X: utils.StrToInt(b[0]), Y: utils.StrToInt(b[1])},
			prize:   getPrize(p[0], p[1], withConversion),
		})
	}
	utils.Debug(machines)
	return
}

func getPrize(x, y string, withConversion bool) (prize utils.Index) {
	if withConversion {
		prize.X = utils.StrToInt(x) + CONVERSION_RATE
		prize.Y = utils.StrToInt(y) + CONVERSION_RATE
	} else {
		prize.X = utils.StrToInt(x)
		prize.Y = utils.StrToInt(y)
	}
	return
}
