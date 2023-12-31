package day02

import (
	"strings"

	"github.com/orlando-pt/aoc/2023/utils"
)

func Part1(lines []string) int {
	totalIdGames := 0

	gameMap := make(map[string]int)
	validGame := true

	for _, line := range lines {
		game := strings.Split(line, ":")
		gameId := utils.StrToInt(strings.Split(game[0], " ")[1])
		gamePlays := game[1]

		plays := strings.Split(gamePlays, ";")

		for _, play := range plays {

			resetCubeNumbers(&gameMap)
			cubeSets := strings.Split(play, ",")

			for _, cubeSet := range cubeSets {
				cubeSet = strings.TrimSpace(cubeSet)
				cubeInfo := strings.Split(cubeSet, " ")
				numberCubes := utils.StrToInt(cubeInfo[0])
				cubeColor := cubeInfo[1]

				if !decrementCubes(&gameMap, cubeColor, numberCubes) {
					validGame = false
					break
				}
				updateMapCubes(&gameMap, cubeColor, numberCubes)
			}

			if !validGame {
				break
			}
			resetCubeNumbers(&gameMap)
		}

		if validGame {
		    totalIdGames += gameId
		}
		validGame = true
	}

	return totalIdGames
}

func Part2(lines []string) int {
	totalIdGames := 0

	gameMap := make(map[string]int)

	for _, line := range lines {
		resetCubeNumbersSol2(&gameMap)

		game := strings.Split(line, ":")
		gamePlays := game[1]

		plays := strings.Split(gamePlays, ";")

		for _, play := range plays {

			cubeSets := strings.Split(play, ",")

			for _, cubeSet := range cubeSets {
				cubeSet = strings.TrimSpace(cubeSet)
				cubeInfo := strings.Split(cubeSet, " ")
				numberCubes := utils.StrToInt(cubeInfo[0])
				cubeColor := cubeInfo[1]

				updateMapCubes(&gameMap, cubeColor, numberCubes)
			}

		}

		totalIdGames += multiplyCubes(&gameMap)
	}

	return totalIdGames

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
