package day01

import (
	"regexp"

	"github.com/orlando-pt/aoc/2023/utils"
)

func Part1(lines []string) int {
    total := 0
    firstLocation := 0
    lastLocation := 0
    line := ""

    for _, line = range lines {
        firstLocation = getFirstLocation(line)
        lastLocation = getLastLocation(line)

        total += (utils.StrToInt(string(line[firstLocation])) * 10) + utils.StrToInt(string(line[lastLocation]))

    }

    return total
}

func Part2(lines []string) int {
	total := 0
	firstLocation := 0
	lastLocation := 0
	line := ""

	for _, line = range lines {
		firstLocation = getFirstLocation(line)
		lastLocation = getLastLocation(line)
        numbers := getRegexNumbers(line, firstLocation, lastLocation)

		total += (numbers[0] * 10) + numbers[1]

	}

    return total
}

func getFirstLocation(line string) int {
	for i := 0; i < len(line); i++ {
		if rune(line[i]) >= '0' && rune(line[i]) <= '9' {
			return i
		}
	}

	return len(line) - 1
}

func getLastLocation(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		if line[i] >= '0' && line[i] <= '9' {
			return i
		}
	}

	return 0
}

func getRegexNumbers(line string, firstLocation int, lastLocation int) []int {
	numbersLetters := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	numbersDigits := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

    firstNumber := utils.StrToInt(string(line[firstLocation]))
    lastNumber := utils.StrToInt(string(line[lastLocation]))

	for i := 0; i < len(numbersDigits); i++ {

		re := regexp.MustCompile(numbersLetters[i])

		locations := re.FindAllStringIndex(line, -1)

		if len(locations) > 0 {
			if locations[0][0] < firstLocation {
				firstLocation = locations[0][0]
				firstNumber = numbersDigits[i]
			}

            resLastLocation := locations[len(locations)-1][0]
			if resLastLocation > lastLocation {
				lastLocation = resLastLocation
				lastNumber = numbersDigits[i]
			}
		}
	}

	return []int{firstNumber, lastNumber}
}

