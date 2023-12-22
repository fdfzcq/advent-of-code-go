package year2023

import (
	"fmt"

	"github.com/fdfzcq/advent-of-code-go/utils"
)

func Day21Part1() {
	input := utils.ReadStringFromFile(2023, 21)
	grid := utils.ParseGrid(input)
	var s utils.Pair
	mem := make(map[utils.Pair]bool)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == "S" {
				s = utils.Pair{A: x, B: y}
				mem[utils.Pair{A: x, B: y}] = false
			} else if grid[y][x] == "." {
				mem[utils.Pair{A: x, B: y}] = false
			}
		}
	}
	res := 0
	queue := []utils.Pair{}
	_, ok := mem[utils.Pair{A: s.GetA() - 1, B: s.GetB()}]
	if ok {
		mem[utils.Pair{A: s.GetA() - 1, B: s.GetB()}] = true
		queue = append(queue, utils.Pair{A: s.GetA() - 1, B: s.GetB()})
		res++
	}
	_, ok = mem[utils.Pair{A: s.GetA() + 1, B: s.GetB()}]
	if ok {
		mem[utils.Pair{A: s.GetA() + 1, B: s.GetB()}] = true
		queue = append(queue, utils.Pair{A: s.GetA() + 1, B: s.GetB()})
		res++
	}
	_, ok = mem[utils.Pair{A: s.GetA(), B: s.GetB() - 1}]
	if ok {
		mem[utils.Pair{A: s.GetA(), B: s.GetB() - 1}] = true
		queue = append(queue, utils.Pair{A: s.GetA(), B: s.GetB() - 1})
		res++
	}
	_, ok = mem[utils.Pair{A: s.GetA(), B: s.GetB() + 1}]
	if ok {
		mem[utils.Pair{A: s.GetA(), B: s.GetB() + 1}] = true
		queue = append(queue, utils.Pair{A: s.GetA(), B: s.GetB() + 1})
		res++
	}
	n := 851 - 1
	for true {
		if n == 0 {
			break
		}
		fmt.Print("res: ")
		fmt.Print(res)
		fmt.Print(" n: ")
		fmt.Println(851 - n)
		res, queue = step(queue, grid, mem, res)
		n = n - 2
	}
	fmt.Println(res)
}

func step(queue []utils.Pair, grid [][]string, mem map[utils.Pair]bool, res int) (int, []utils.Pair) {
	var newqueue []utils.Pair
	for _, p := range queue {
		x := p.GetA()
		y := p.GetB()
		next := utils.Pair{A: x - 2, B: y}
		xl := len(grid[0])
		yl := len(grid)
		if isReachable(next, mem, []utils.Pair{{A: x - 1, B: y}}, xl, yl) {
			newqueue = append(newqueue, next)
			mem[next] = true
			res++
		}
		next = utils.Pair{A: x + 2, B: y}
		if isReachable(next, mem, []utils.Pair{{A: x + 1, B: y}}, xl, yl) {
			newqueue = append(newqueue, next)
			mem[next] = true
			res++
		}
		next = utils.Pair{A: x + 1, B: y - 1}
		if isReachable(next, mem, []utils.Pair{{A: x + 1, B: y}, {A: x, B: y - 1}}, xl, yl) {
			newqueue = append(newqueue, next)
			mem[next] = true
			res++
		}
		next = utils.Pair{A: x + 1, B: y + 1}
		if isReachable(next, mem, []utils.Pair{{A: x, B: y + 1}, {A: x + 1, B: y}}, xl, yl) {
			newqueue = append(newqueue, next)
			mem[next] = true
			res++
		}
		next = utils.Pair{A: x, B: y - 2}
		if isReachable(next, mem, []utils.Pair{{A: x, B: y - 1}}, xl, yl) {
			newqueue = append(newqueue, next)
			mem[next] = true
			res++
		}
		next = utils.Pair{A: x, B: y + 2}
		if isReachable(next, mem, []utils.Pair{{A: x, B: y + 1}}, xl, yl) {
			newqueue = append(newqueue, next)
			mem[next] = true
			res++
		}
		next = utils.Pair{A: x - 1, B: y + 1}
		if isReachable(next, mem, []utils.Pair{{A: x, B: y + 1}, {A: x - 1, B: y}}, xl, yl) {
			newqueue = append(newqueue, next)
			mem[next] = true
			res++
		}
		next = utils.Pair{A: x - 1, B: y - 1}
		if isReachable(next, mem, []utils.Pair{{A: x, B: y - 1}, {A: x - 1, B: y}}, xl, yl) {
			newqueue = append(newqueue, next)
			mem[next] = true
			res++
		}
	}
	return res, newqueue
}

func isReachable(p utils.Pair, mem map[utils.Pair]bool, ns []utils.Pair, xl int, yl int) bool {
	v, ok := mem[p]
	_, ok2 := mem[normalise(p, xl, yl)]
	if (ok && !v) || (!ok && ok2) {
		isReachable := false
		for _, n := range ns {
			_, ok := mem[n]
			_, ok2 := mem[normalise(n, xl, yl)]
			if ok || (!ok && ok2) {
				isReachable = true
				break
			}
		}
		return isReachable
	}
	return false
}

func normalise(p utils.Pair, xl int, yl int) utils.Pair {
	var newX, newY int
	if p.GetA() >= 0 {
		newX = p.GetA() % xl
	} else {
		newX = (xl + (p.GetA() % xl)) % xl
	}
	if p.GetB() >= 0 {
		newY = p.GetB() % yl
	} else {
		newY = (yl + (p.GetB() % yl)) % yl
	}
	return utils.Pair{A: newX, B: newY}
}
