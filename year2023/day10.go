package year2023

import (
	"fmt"
	"slices"

	"github.com/fdfzcq/advent-of-code-go/utils"
)

func Day10Part1() {
	input := utils.ReadStringFromFile(2023, 10)
	grid := utils.ParseGrid(input)
	start := findStart(grid)
	findLoop(grid, start, make([]utils.Pair, 0))
}

func findStart(grid [][]string) utils.Pair {
	for y, r := range grid {
		for x, c := range r {
			if c == "S" {
				return utils.Pair{A: x, B: y}
			}
		}
	}
	return utils.Pair{}
}

func findLoop(grid [][]string, start utils.Pair, loop []utils.Pair) []utils.Pair {
	next := start
	for true {
		if !slices.Contains(loop, next) {
			loop = append(loop, next)
		}
		neighbours := connectedNodes(next, grid[next.GetB()][next.GetA()], grid)
		cont := false
		for _, n := range neighbours {
			if !slices.Contains(loop, n) {
				loop = append(loop, n)
				next = n
				cont = true
				break
			}
		}
		if !cont {
			fmt.Println(float64(len(loop)) / 2.0)
			break
		}
	}
	return loop
}

func connectedNodes(n utils.Pair, c string, grid [][]string) []utils.Pair {
	var nodes []utils.Pair
	if c == "-" {
		nodes = rightNode(n, grid, nodes)
		nodes = leftNode(n, grid, nodes)
	}
	if c == "|" {
		nodes = upNode(n, grid, nodes)
		nodes = downNode(n, grid, nodes)
	}
	if c == "F" {
		nodes = downNode(n, grid, nodes)
		nodes = rightNode(n, grid, nodes)
	}
	if c == "7" {
		nodes = downNode(n, grid, nodes)
		nodes = leftNode(n, grid, nodes)
	}
	if c == "J" {
		nodes = upNode(n, grid, nodes)
		nodes = leftNode(n, grid, nodes)
	}
	if c == "L" {
		nodes = upNode(n, grid, nodes)
		nodes = rightNode(n, grid, nodes)
	}
	if c == "S" {
		nodes = downNode(n, grid, nodes)
		nodes = upNode(n, grid, nodes)
		nodes = rightNode(n, grid, nodes)
		nodes = leftNode(n, grid, nodes)
	}

	return nodes
}

func rightNode(n utils.Pair, grid [][]string, nodes []utils.Pair) []utils.Pair {
	if n.GetA()+1 < len(grid[0]) {
		if grid[n.GetB()][n.GetA()+1] == "-" || grid[n.GetB()][n.GetA()+1] == "J" || grid[n.GetB()][n.GetA()+1] == "7" {
			return append(nodes, utils.Pair{A: n.GetA() + 1, B: n.GetB()})
		}
	}
	return nodes
}

func leftNode(n utils.Pair, grid [][]string, nodes []utils.Pair) []utils.Pair {
	if n.GetA()-1 >= 0 {
		if grid[n.GetB()][n.GetA()-1] == "-" || grid[n.GetB()][n.GetA()-1] == "F" || grid[n.GetB()][n.GetA()-1] == "L" {
			return append(nodes, utils.Pair{A: n.GetA() - 1, B: n.GetB()})
		}
	}
	return nodes
}

func upNode(n utils.Pair, grid [][]string, nodes []utils.Pair) []utils.Pair {
	if n.GetB()-1 >= 0 {
		if grid[n.GetB()-1][n.GetA()] == "|" || grid[n.GetB()-1][n.GetA()] == "F" || grid[n.GetB()-1][n.GetA()] == "7" {
			return append(nodes, utils.Pair{A: n.GetA(), B: n.GetB() - 1})
		}
	}
	return nodes
}

func downNode(n utils.Pair, grid [][]string, nodes []utils.Pair) []utils.Pair {
	if n.GetB()+1 < len(grid) {
		if grid[n.GetB()+1][n.GetA()] == "|" || grid[n.GetB()+1][n.GetA()] == "J" || grid[n.GetB()+1][n.GetA()] == "L" {
			return append(nodes, utils.Pair{A: n.GetA(), B: n.GetB() + 1})
		}
	}
	return nodes
}

func Day10Part2() {
	input := utils.ReadStringFromFile(2023, 10)
	grid := utils.ParseGrid(input)
	start := findStart(grid)
	loop := findLoop(grid, start, make([]utils.Pair, 0))
	c := 0
	for y := 0; y < len(grid); y++ {
		isIn := false
		prev := ""
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == "S" {
				grid[y][x] = "|"
			}
			if !slices.Contains(loop, utils.Pair{A: x, B: y}) && isIn {
				c++
			} else if slices.Contains(loop, utils.Pair{A: x, B: y}) {
				if grid[y][x] == "|" {
					isIn = !isIn
				} else if grid[y][x] == "L" || grid[y][x] == "F" {
					isIn = !isIn
					prev = grid[y][x]
				} else if grid[y][x] == "J" || grid[y][x] == "7" {
					if (grid[y][x] == "J" && prev == "L") || (grid[y][x] == "7" && prev == "F") {
						isIn = !isIn
					}
				}
			}
		}
	}
	fmt.Println(c)
}
