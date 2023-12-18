package year2023

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/fdfzcq/advent-of-code-go/utils"
)

type heatstate struct {
	x     int
	y     int
	dirX  int
	dirY  int
	block int
}

type ht struct {
	state heatstate
	heat  int
}

func Day17Part1() {
	input := utils.ReadStringFromTestFile(2023, 17)
	grid := utils.ParseGrid(input)
	queue := []ht{
		{state: heatstate{
			x: 1, y: 0, dirX: 1, dirY: 0, block: 0,
		}, heat: 0},
		{state: heatstate{
			x: 0, y: 1, dirX: 0, dirY: 1, block: 0,
		}, heat: 0},
	}
	fmt.Println(minHeatLoss(queue, grid, make(map[heatstate]int)))
}

func minHeatLoss(queue []ht, grid [][]string, cache map[heatstate]int) int {
	min := 999999
	for true {
		var newQueue []ht
		for _, q := range queue {
			v, ok := cache[q.state]
			if ok && v <= q.heat || q.heat >= min {
				continue
			}
			if q.state.x == len(grid[0])-1 && q.state.y == len(grid)-1 {
				n, _ := strconv.Atoi(grid[q.state.y][q.state.x])
				min = utils.Min(min, q.heat+n)
				continue
			}
			n, _ := strconv.Atoi(grid[q.state.y][q.state.x])
			cache[q.state] = q.heat
			if q.state.block < 2 {
				newX := q.state.x + q.state.dirX
				newY := q.state.y + q.state.dirY
				if newX >= 0 && newX < len(grid[0]) && newY >= 0 && newY < len(grid) {
					newQueue = append(newQueue, ht{
						state: heatstate{
							x:     newX,
							y:     newY,
							dirX:  q.state.dirX,
							dirY:  q.state.dirY,
							block: q.state.block + 1,
						},
						heat: q.heat + n,
					})
				}
			}
			var nextDirs []utils.Pair
			if q.state.dirX == 0 {
				nextDirs = append(nextDirs, utils.Pair{A: 1, B: 0}, utils.Pair{A: -1, B: 0})
			} else if q.state.dirY == 0 {
				nextDirs = append(nextDirs, utils.Pair{A: 0, B: 1}, utils.Pair{A: 0, B: -1})
			}
			for _, p := range nextDirs {
				newX := q.state.x + p.GetA()
				newY := q.state.y + p.GetB()
				if newX >= 0 && newX < len(grid[0]) && newY >= 0 && newY < len(grid) {
					newQueue = append(newQueue, ht{
						state: heatstate{
							x:     newX,
							y:     newY,
							dirX:  p.GetA(),
							dirY:  p.GetB(),
							block: 0,
						},
						heat: q.heat + n,
					})
				}
			}
		}
		if len(newQueue) == 0 {
			break
		}
		slices.SortFunc(newQueue, func(a, b ht) int {
			return a.heat - b.heat
		})
		queue = newQueue
	}
	return min
}

func Day17Part2() {
	input := utils.ReadStringFromFile(2023, 17)
	grid := utils.ParseGrid(input)
	queue := []ht{
		{state: heatstate{
			x: 1, y: 0, dirX: 1, dirY: 0, block: 0,
		}, heat: 0},
		{state: heatstate{
			x: 0, y: 1, dirX: 0, dirY: 1, block: 0,
		}, heat: 0},
	}
	fmt.Println(minHeatLoss2(queue, grid, make(map[heatstate]int)))
}

func minHeatLoss2(queue []ht, grid [][]string, cache map[heatstate]int) int {
	min := 999999
	for true {
		var newQueue []ht
		for _, q := range queue {
			v, ok := cache[q.state]
			if ok && v <= q.heat || q.heat >= min {
				continue
			}
			if q.state.x == len(grid[0])-1 && q.state.y == len(grid)-1 {
				if q.state.block >= 3 {
					n, _ := strconv.Atoi(grid[q.state.y][q.state.x])
					min = utils.Min(min, q.heat+n)
				}
				continue
			}
			n, _ := strconv.Atoi(grid[q.state.y][q.state.x])
			cache[q.state] = q.heat
			if q.state.block < 9 {
				newX := q.state.x + q.state.dirX
				newY := q.state.y + q.state.dirY
				if newX >= 0 && newX < len(grid[0]) && newY >= 0 && newY < len(grid) {
					newQueue = append(newQueue, ht{
						state: heatstate{
							x:     newX,
							y:     newY,
							dirX:  q.state.dirX,
							dirY:  q.state.dirY,
							block: q.state.block + 1,
						},
						heat: q.heat + n,
					})
				}
			}
			if q.state.block >= 3 {
				var nextDirs []utils.Pair
				if q.state.dirX == 0 {
					nextDirs = append(nextDirs, utils.Pair{A: 1, B: 0}, utils.Pair{A: -1, B: 0})
				} else if q.state.dirY == 0 {
					nextDirs = append(nextDirs, utils.Pair{A: 0, B: 1}, utils.Pair{A: 0, B: -1})
				}
				for _, p := range nextDirs {
					newX := q.state.x + p.GetA()
					newY := q.state.y + p.GetB()
					if newX >= 0 && newX < len(grid[0]) && newY >= 0 && newY < len(grid) {
						newQueue = append(newQueue, ht{
							state: heatstate{
								x:     newX,
								y:     newY,
								dirX:  p.GetA(),
								dirY:  p.GetB(),
								block: 0,
							},
							heat: q.heat + n,
						})
					}
				}
			}
		}
		if len(newQueue) == 0 {
			break
		}
		slices.SortFunc(newQueue, func(a, b ht) int {
			return a.heat - b.heat
		})
		queue = newQueue
	}
	return min
}
