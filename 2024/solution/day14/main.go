package day14

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"regexp"

	"github.com/orlando-pt/aoc/2024/utils"
)

const (
	// to test the solution
	MAX_X_TEST int = 11
	MAX_Y_TEST int = 7
	// to submit the solution
	MAX_X int = 101
	MAX_Y int = 103

	SECONDS int = 100
)

type Robot struct {
	position utils.Index
	velocity utils.Index
	testing  bool
}

type Quadrants struct {
	q1 int
	q2 int
	q3 int
	q4 int
}

func Part1(lines []string) (sum int) {
	robots := parseInput(lines)

	utils.Debug(robots)

	quadrants := Quadrants{}
	for _, robot := range robots {
		for j := 0; j < SECONDS; j++ {
			robot.move()
		}

		utils.Debug("Robot position: ", robot.position)
		quadrants.addQuadrant(robot)
		utils.Debug("Quadrant: ", quadrants)
	}

	return quadrants.calculateSafetyFactor()
}

func Part2(lines []string) (sum int) {
	robots := parseInput(lines)
	max_x, max_y := getMax(robots[0].testing)

	for j := 1; j < max_x*max_y; j++ {
		img := image.NewGray(image.Rect(0, 0, 8*MAX_X, 8*MAX_Y))

		tmpRobots := make([]Robot, len(robots))
		for i, robot := range robots {
			robot.move()
			tmpRobots[i] = robot

			for x := 0; x < 8; x++ {
				for y := 0; y < 8; y++ {
					img.Set(8*robot.position.X+x, 8*robot.position.Y+y, image.White)
				}
			}
		}
		robots = tmpRobots

		if has10InLine(robots, max_x, max_y) {
			file, err := os.Create(filepath.Join("image", fmt.Sprintf("%d.png", j)))
			if err != nil {
				println("Error creating file: ", err)
			}
			defer file.Close()

			err = png.Encode(file, img)
			if err != nil {
				println("Error encoding file: ", err)
			}

			return j
		}
	}

	return
}

func has10InLine(robots []Robot, max_x, max_y int) bool {
	positions := make([]bool, max_x*max_x+max_y*max_y)
	for _, robot := range robots {
		utils.Debug("Robot position: ", robot.position)
		positions[robot.position.X*max_x+robot.position.Y] = true
	}

	for i := 0; i < max_x*max_y; i++ {
		if positions[i] {
			count := 0
			for j := i; j < max_x*max_y; j++ {
				if positions[j] {
					count++
				} else {
					break
				}
			}
			if count >= 10 {
				return true
			}
		}
	}

	return false
}

func (q *Quadrants) addQuadrant(robot Robot) {
	switch robot.quadrant() {
	case 1:
		q.q1++
	case 2:
		q.q2++
	case 3:
		q.q3++
	case 4:
		q.q4++
	}
}

func (q Quadrants) calculateSafetyFactor() int {
	return q.q1 * q.q2 * q.q3 * q.q4
}

func (r *Robot) quadrant() int {
	max_x, max_y := getMax(r.testing)

	if r.position.X < max_x/2 {
		if r.position.Y < max_y/2 {
			return 1
		} else if r.position.Y > max_y/2 {
			return 3
		}
	} else if r.position.X > max_x/2 {
		if r.position.Y < max_y/2 {
			return 2
		} else if r.position.Y > max_y/2 {
			return 4
		}
	}
	return -1
}

func (r *Robot) move() {
	max_x, max_y := getMax(r.testing)

	r.position.X += r.velocity.X
	r.position.Y += r.velocity.Y

	if r.position.X < 0 {
		r.position.X += max_x
	} else if r.position.X >= max_x {
		r.position.X -= max_x
	}

	if r.position.Y < 0 {
		r.position.Y += max_y
	} else if r.position.Y >= max_y {
		r.position.Y -= max_y
	}
}

func parseInput(lines []string) (robots []Robot) {
	numberRegex := regexp.MustCompile(`-?\d+`)
	for _, line := range lines {
		numbers := numberRegex.FindAllString(line, -1)
		robots = append(robots, Robot{
			position: indexFromPosition(numbers[0], numbers[1]),
			velocity: indexFromPosition(numbers[2], numbers[3]),
			testing:  isTest(len(lines)),
		})
	}
	return
}

func indexFromPosition(x, y string) (index utils.Index) {
	index.X = utils.StrToInt(x)
	index.Y = utils.StrToInt(y)
	return
}

func isTest(robots int) bool {
	if robots < 20 {
		return true
	}
	return false
}

func getMax(testing bool) (int, int) {
	if testing {
		return MAX_X_TEST, MAX_Y_TEST
	}
	return MAX_X, MAX_Y
}
