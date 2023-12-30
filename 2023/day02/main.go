package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	totalIdGames := 0

	gameMap := make(map[string]int)
	// validGame := true
	fileScanner := readFile("input.txt")

	for fileScanner.Scan() {
        resetCubeNumbersSol2(&gameMap)
		line := fileScanner.Text()

		game := strings.Split(line, ":")
		// gameId, _ := strconv.Atoi(strings.Split(game[0], " ")[1])
		gamePlays := game[1]

		plays := strings.Split(gamePlays, ";")

		for _, play := range plays {

			// resetCubeNumbers(&gameMap)
			cubeSets := strings.Split(play, ",")

			for _, cubeSet := range cubeSets {
				cubeSet = strings.TrimSpace(cubeSet)
				cubeInfo := strings.Split(cubeSet, " ")
				numberCubes, _ := strconv.Atoi(cubeInfo[0])
				cubeColor := cubeInfo[1]

				// if !decrementCubes(&gameMap, cubeColor, numberCubes) {
				// 	validGame = false
				// 	break
				// }
                updateMapCubes(&gameMap, cubeColor, numberCubes)
			}

			// if !validGame {
			// 	break
			// }
            // resetCubeNumbers(&gameMap)
		}

        // if validGame {
        //     totalIdGames += gameId
        // }
        // validGame = true

        totalIdGames += multiplyCubes(&gameMap)
	}

    fmt.Println("Total Id Games:", totalIdGames)
}

func multiplyCubes(gameMap *map[string]int) int {
    return (*gameMap)["red"] * (*gameMap)["green"] * (*gameMap)["blue"]
}

func decrementCubes(gameMap *map[string]int, cubeColor string, numberCubes int) bool {
	(*gameMap)[cubeColor] = (*gameMap)[cubeColor] - numberCubes

	return (*gameMap)[cubeColor] >= 0
}

func updateMapCubes(gameMap *map[string]int, cubeColor string, numberCubes int) {
    if (*gameMap)[cubeColor] < numberCubes {
        (*gameMap)[cubeColor] = numberCubes
    }
}

func resetCubeNumbers(gameMap *map[string]int) {
	(*gameMap)["red"], (*gameMap)["green"], (*gameMap)["blue"] = 12, 13, 14
}

func resetCubeNumbersSol2(gameMap *map[string]int) {
	(*gameMap)["red"], (*gameMap)["green"], (*gameMap)["blue"] = 0, 0, 0
}

func readFile(fileName string) *bufio.Scanner {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	return fileScanner
}
