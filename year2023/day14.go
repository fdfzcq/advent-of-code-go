package year2023

import (
	"fmt"

	"github.com/fdfzcq/advent-of-code-go/utils"
)

func Day14() {
	input := utils.ReadStringFromFile(2023, 14)
	grid := utils.ParseGrid(input)
	res := spinRound(grid, 1000000000)
	fmt.Println(res)
}

func spinRound(grid [][]string, round int) int {
	res := 0
	for i := 0; i < round; i++ {
		grid = spin(grid)
		newres := 0
		for y := 0; y < len(grid); y++ {
			c := 0
			for x := 0; x < len(grid[0]); x++ {
				if grid[y][x] == "O" {
					c++
				}
			}
			newres = newres + c*(len(grid)-y)
		}
		fmt.Print(i + 1)
		fmt.Print(" : ")
		fmt.Println(newres)
		res = newres
	}
	return res
}

func spin(grid [][]string) [][]string {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == "O" {
				for n := 1; y-n >= -1; n++ {
					if y-n == -1 || grid[y-n][x] != "." {
						grid[y][x] = "."
						grid[y-n+1][x] = "O"
						break
					}
				}
			}
		}
	}
	for x := 0; x < len(grid[0]); x++ {
		for y := 0; y < len(grid); y++ {
			if grid[y][x] == "O" {
				for n := 1; x-n >= -1; n++ {
					if x-n == -1 || grid[y][x-n] != "." {
						grid[y][x] = "."
						grid[y][x-n+1] = "O"
						break
					}
				}
			}
		}
	}
	for y := len(grid) - 1; y >= 0; y-- {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == "O" {
				for n := 1; y+n <= len(grid); n++ {
					if y+n == len(grid) || grid[y+n][x] != "." {
						grid[y][x] = "."
						grid[y+n-1][x] = "O"
						break
					}
				}
			}
		}
	}
	for x := len(grid[0]) - 1; x >= 0; x-- {
		for y := 0; y < len(grid); y++ {
			if grid[y][x] == "O" {
				for n := 1; x+n <= len(grid[0]); n++ {
					if x+n == len(grid[0]) || grid[y][x+n] != "." {
						grid[y][x] = "."
						grid[y][x+n-1] = "O"
						break
					}
				}
			}
		}
	}

	return grid
}
