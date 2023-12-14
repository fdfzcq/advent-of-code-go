package year2023

import (
	"fmt"
	"math"

	"github.com/fdfzcq/advent-of-code-go/utils"
)

func Day11Part1() {
	input := utils.ReadStringFromFile(2023, 11)
	universe := expandCols(expandRows(utils.ParseGrid(input)))
	fmt.Println(shortestDist(universe))
}

func shortestDist(universe [][]string) int {
	var stars []utils.Pair
	for y := 0; y < len(universe); y++ {
		for x := 0; x < len(universe[0]); x++ {
			if universe[y][x] == "#" {
				stars = append(stars, utils.Pair{A: x, B: y})
			}
		}
	}
	sum := 0
	for i, star := range stars {
		for j := i + 1; j < len(stars); j++ {
			star2 := stars[j]
			sum = sum + int(math.Abs(float64(star2.GetA()-star.GetA()))) + int(math.Abs(float64(star2.GetB()-star.GetB())))
		}
	}
	return sum
}

func expandRows(grid [][]string) [][]string {
	var newGrid [][]string
	for y := 0; y < len(grid); y++ {
		shouldExpand := true
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] != "." {
				shouldExpand = false
			}
		}
		newGrid = append(newGrid, grid[y])
		if shouldExpand {
			newGrid = append(newGrid, grid[y])
		}
	}
	return newGrid
}

func expandCols(grid [][]string) [][]string {
	newGrid := make([][]string, len(grid))
	for x := 0; x < len(grid[0]); x++ {
		shouldExpand := true
		for y := 0; y < len(grid); y++ {
			if grid[y][x] != "." {
				shouldExpand = false
			}
			newGrid[y] = append(newGrid[y], grid[y][x])
		}
		if shouldExpand {
			for y := 0; y < len(grid); y++ {
				newGrid[y] = append(newGrid[y], ".")
			}
		}
	}
	return newGrid
}

func Day11Part2() {
	input := utils.ReadStringFromFile(2023, 11)
	universe := utils.ParseGrid(input)
	var emptyRows []int
	var emptyColumns []int
	for y := 0; y < len(universe); y++ {
		shouldExpand := true
		for x := 0; x < len(universe[0]); x++ {
			if universe[y][x] != "." {
				shouldExpand = false
			}
		}
		if shouldExpand {
			emptyRows = append(emptyRows, y)
		}
	}
	for x := 0; x < len(universe[0]); x++ {
		shouldExpand := true
		for y := 0; y < len(universe); y++ {
			if universe[y][x] != "." {
				shouldExpand = false
			}
		}
		if shouldExpand {
			emptyColumns = append(emptyColumns, x)
		}
	}
	fmt.Println(shortestDist2(universe, 1000000-1, emptyRows, emptyColumns))
}

func shortestDist2(universe [][]string, v int, empRows []int, empCols []int) int {
	var stars []utils.Pair
	for y := 0; y < len(universe); y++ {
		for x := 0; x < len(universe[0]); x++ {
			if universe[y][x] == "#" {
				stars = append(stars, utils.Pair{A: x, B: y})
			}
		}
	}
	sum := 0
	for i, star := range stars {
		for j := i + 1; j < len(stars); j++ {
			star2 := stars[j]
			rows := 0
			cols := 0
			for _, r := range empRows {
				if r > utils.Min(star2.GetB(), star.GetB()) && r < utils.Max(star2.GetB(), star.GetB()) {
					rows++
				}
			}
			for _, e := range empCols {
				if e > utils.Min(star2.GetA(), star.GetA()) && e < utils.Max(star2.GetA(), star.GetA()) {
					cols++
				}
			}
			sum = sum + int(math.Abs(float64(star2.GetA()-star.GetA()))) + v*cols + v*rows + int(math.Abs(float64(star2.GetB()-star.GetB())))
		}
	}
	return sum
}
