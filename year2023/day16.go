package year2023

import (
	"fmt"
	"time"

	"github.com/fdfzcq/advent-of-code-go/utils"
)

type beam struct {
	x    int
	y    int
	xDir int
	yDir int
}

func Day16Part1() {
	input := utils.ReadStringFromFile(2023, 16)
	grid := utils.ParseGrid(input)
	energizeTiles(grid, beam{x: 0, y: 0, xDir: 1, yDir: 0})
}

func energizeTiles(grid [][]string, start beam) int {
	queue := make(chan beam)
	res := 0
	mem := make(map[beam]int)
	energized := make(map[utils.Pair]bool)
	go func() {
		queue <- start
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		close(queue)
	}()

	for p := range queue {
		v, ok := mem[p]
		if !ok {
			c := grid[p.y][p.x]
			if c == "." {
				newX := p.x + p.xDir
				newY := p.y + p.yDir
				if newX >= 0 && newX < len(grid[0]) && newY >= 0 && newY < len(grid) {
					go func(p beam, newX int, newY int) {
						queue <- beam{x: newX, y: newY, xDir: p.xDir, yDir: p.yDir}
					}(p, newX, newY)
				}
			}
			if c == "/" {
				var newXDir, newYDir int
				if p.xDir == 0 {
					newXDir = -p.yDir
					newYDir = 0
				} else if p.yDir == 0 {
					newXDir = 0
					newYDir = -p.xDir
				}
				newX := p.x + newXDir
				newY := p.y + newYDir
				if newX >= 0 && newX < len(grid[0]) && newY >= 0 && newY < len(grid) {
					go func(newX int, newY int, newXDir int, newYDir int) {
						queue <- beam{x: newX, y: newY, xDir: newXDir, yDir: newYDir}
					}(newX, newY, newXDir, newYDir)
				}
			}
			if c == "\\" {
				var newXDir, newYDir int
				if p.xDir == 0 {
					newXDir = p.yDir
					newYDir = 0
				} else if p.yDir == 0 {
					newXDir = 0
					newYDir = p.xDir
				}
				newX := p.x + newXDir
				newY := p.y + newYDir
				if newX >= 0 && newX < len(grid[0]) && newY >= 0 && newY < len(grid) {
					go func(newX int, newY int, newXDir int, newYDir int) {
						queue <- beam{x: newX, y: newY, xDir: newXDir, yDir: newYDir}
					}(newX, newY, newXDir, newYDir)
				}
			}
			if c == "|" {
				newX := p.x
				newY := p.y + 1
				if newX >= 0 && newX < len(grid[0]) && newY >= 0 && newY < len(grid) {
					go func(newX int, newY int) {
						queue <- beam{x: newX, y: newY, xDir: 0, yDir: 1}
					}(newX, newY)
				}
				newX = p.x
				newY = p.y - 1
				if newX >= 0 && newX < len(grid[0]) && newY >= 0 && newY < len(grid) {
					go func(newX int, newY int) {
						queue <- beam{x: newX, y: newY, xDir: 0, yDir: -1}
					}(newX, newY)
				}
			}
			if c == "-" {
				newX := p.x + 1
				newY := p.y
				if newX >= 0 && newX < len(grid[0]) && newY >= 0 && newY < len(grid) {
					go func(newX int, newY int) {
						queue <- beam{x: newX, y: newY, xDir: 1, yDir: 0}
					}(newX, newY)
				}
				newX = p.x - 1
				newY = p.y
				if newX >= 0 && newX < len(grid[0]) && newY >= 0 && newY < len(grid) {
					go func(newX int, newY int) {
						queue <- beam{x: newX, y: newY, xDir: -1, yDir: 0}
					}(newX, newY)
				}
			}
			_, ok := energized[utils.Pair{A: p.x, B: p.y}]
			if !ok {
				res++
				energized[utils.Pair{A: p.x, B: p.y}] = true
			}
			mem[p] = 1
		} else {
			mem[p] = v + 1
		}
	}
	return res
}

func Day16Part2() {
	input := utils.ReadStringFromFile(2023, 16)
	grid := utils.ParseGrid(input)
	res := -1
	for x := 0; x < len(grid[0]); x++ {
		res = utils.Max(res, energizeTiles(grid, beam{x: x, y: 0, xDir: 0, yDir: 1}))
		res = utils.Max(res, energizeTiles(grid, beam{x: x, y: len(grid) - 1, xDir: 0, yDir: -1}))
		fmt.Println(res)
	}
	for y := 0; y < len(grid); y++ {
		res = utils.Max(res, energizeTiles(grid, beam{x: 0, y: y, xDir: 1, yDir: 0}))
		res = utils.Max(res, energizeTiles(grid, beam{x: len(grid[0]) - 1, y: 0, xDir: -1, yDir: 0}))
		fmt.Println(res)
	}
	fmt.Println(res)
}
