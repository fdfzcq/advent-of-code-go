package year2023

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/fdfzcq/advent-of-code-go/utils"
)

type game struct {
	id       int
	cubeSets []*cube
}

type cube struct {
	red   int
	blue  int
	green int
}

func Day2Part1() {
	str := utils.ReadStringFromFile(2023, 2)
	games := parseInputDay2(str)
	acc := 0
	for _, game := range games {
		possible := true
		for _, cube := range game.cubeSets {
			if cube.red > 12 || cube.green > 13 || cube.blue > 14 {
				possible = false
			}
		}
		if possible {
			acc = acc + game.id
		}
	}
	fmt.Print(acc)
}

func parseInputDay2(str string) []*game {
	gameStrs := strings.Split(str, "\n")
	games := make([]*game, len(gameStrs))
	for i, gameStr := range gameStrs {
		s := strings.Split(gameStr, ": ")
		s1 := strings.Split(s[0], " ")
		gameId, _ := strconv.Atoi(s1[1])
		subsetStrs := strings.Split(s[1], "; ")
		cubeSets := make([]*cube, len(subsetStrs))
		for n, subsetStr := range subsetStrs {
			cubeStrs := strings.Split(subsetStr, ", ")
			cube := &cube{}
			for _, cubeStr := range cubeStrs {
				cubeS := strings.Split(cubeStr, " ")
				v, _ := strconv.Atoi(cubeS[0])
				if cubeS[1] == "red" {
					cube.red = v
				}
				if cubeS[1] == "blue" {
					cube.blue = v
				}
				if cubeS[1] == "green" {
					cube.green = v
				}
			}
			cubeSets[n] = cube
		}
		game := &game{
			id:       gameId,
			cubeSets: cubeSets,
		}
		games[i] = game
	}
	return games
}

func Day2Part2() {
	str := utils.ReadStringFromFile(2023, 2)
	games := parseInputDay2(str)
	acc := 0
	for _, game := range games {
		maxRed := 0
		maxBlue := 0
		maxGreen := 0
		for _, cube := range game.cubeSets {
			if cube.red != 0 {
				maxRed = int(math.Max(float64(maxRed), float64(cube.red)))
			}
			if cube.blue != 0 {
				maxBlue = int(math.Max(float64(maxBlue), float64(cube.blue)))
			}
			if cube.green != 0 {
				maxGreen = int(math.Max(float64(maxGreen), float64(cube.green)))
			}
		}
		acc = acc + maxRed*maxBlue*maxGreen
	}
	fmt.Print(acc)
}
