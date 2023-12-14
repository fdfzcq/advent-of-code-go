package year2023

import (
	"fmt"
	"strings"

	"github.com/fdfzcq/advent-of-code-go/utils"
)

func Day13Part1() {
	input := utils.ReadStringFromFile(2023, 13)
	parts := strings.Split(input, "\n\n")
	res := 0
	for _, p := range parts {
		v, t := findMirror(utils.ParseGrid(p))
		if t == "h" {
			res = res + 100*v
		} else {
			res = res + v
		}
	}
	fmt.Println(res)
}

func findMirror(grid [][]string) (int, string) {
	for x := 1; x < len(grid[0]); x++ {
		isMirror := true
		for y := 0; y < len(grid); y++ {
			for i := 1; x-i >= 0 && x+i-1 < len(grid[0]); i++ {
				if grid[y][x-i] != grid[y][x+i-1] {
					isMirror = false
					break
				}
			}
			if !isMirror {
				break
			}
		}
		if isMirror {
			return x, "v"
		}
	}
	for y := 1; y < len(grid); y++ {
		isMirror := true
		for x := 0; x < len(grid[0]); x++ {
			for i := 1; y-i >= 0 && y+i-1 < len(grid); i++ {
				if grid[y-i][x] != grid[y+i-1][x] {
					isMirror = false
					break
				}
			}
			if !isMirror {
				break
			}
		}
		if isMirror {
			return y, "h"
		}
	}
	return 0, ""
}

func Day13Part2() {
	input := utils.ReadStringFromFile(2023, 13)
	parts := strings.Split(input, "\n\n")
	res := 0
	for _, p := range parts {
		v, t := findSmudgeMirror(utils.ParseGrid(p))
		if t == "h" {
			res = res + 100*v
		} else {
			res = res + v
		}
	}
	fmt.Println(res)
}

func findSmudgeMirror(grid [][]string) (int, string) {
	for x := 1; x < len(grid[0]); x++ {
		noOfSumdge := 0
		for y := 0; y < len(grid); y++ {
			for i := 1; x-i >= 0 && x+i-1 < len(grid[0]); i++ {
				if grid[y][x-i] != grid[y][x+i-1] {
					noOfSumdge++
					break
				}
			}
			if noOfSumdge > 1 {
				break
			}
		}
		if noOfSumdge == 1 {
			return x, "v"
		}
	}
	for y := 1; y < len(grid); y++ {
		noOfSumdge := 0
		for x := 0; x < len(grid[0]); x++ {
			for i := 1; y-i >= 0 && y+i-1 < len(grid); i++ {
				if grid[y-i][x] != grid[y+i-1][x] {
					noOfSumdge++
					break
				}
			}
			if noOfSumdge > 1 {
				break
			}
		}
		if noOfSumdge == 1 {
			return y, "h"
		}
	}
	return 0, ""
}
