package year2023

import (
	"fmt"
	"slices"

	"github.com/fdfzcq/advent-of-code-go/utils"
)

func isIntersection(x int, y int, grid [][]string) bool {
	if x <= 0 || x > 139 || y <= 0 || y > 139 {
		return false
	}
	c := 0
	if grid[y-1][x] != "#" {
		c++
	}
	if grid[y+1][x] != "#" {
		c++
	}
	if grid[y][x-1] != "#" {
		c++
	}
	if grid[y][x+1] != "#" {
		c++
	}
	return c >= 3
}

func reachable(grid [][]string, p utils.Pair, prev utils.Pair) bool {
	if p.GetA() < 0 || p.GetB() < 0 {
		return false
	}
	s := grid[p.GetB()][p.GetA()]
	if s == "#" {
		return false
	} else {
		return true
	}
}

type node struct {
	x int
	y int
}

type state23 struct {
	n      node
	seen   []node
	length int
}

func Day23() {
	input := utils.ReadStringFromFile(2023, 23)
	grid := utils.ParseGrid(input)
	nodes := make(map[node]map[node]int)
	for y := 0; y < 141; y++ {
		for x := 0; x < 141; x++ {
			if grid[y][x] != "#" && isIntersection(x, y, grid) {
				nodes[node{x: x, y: y}] = make(map[node]int)
			}
		}
	}
	nodes[node{x: 1, y: 0}] = make(map[node]int)
	nodes[node{x: 139, y: 140}] = make(map[node]int)
	for n, m := range nodes {
		nexts := []node{
			{x: n.x - 1, y: n.y},
			{x: n.x + 1, y: n.y},
			{x: n.x, y: n.y - 1},
			{x: n.x, y: n.y + 1}}
		l := 1
		seen := []node{n}
		for true {
			var newNexts []node
			for _, next := range nexts {
				_, ok := nodes[next]
				if ok && next != n {
					m[next] = l
				} else if next.x >= 0 && next.y >= 0 && next.x < 141 && next.y < 141 &&
					grid[next.y][next.x] != "#" && !slices.Contains(seen, next) {
					seen = append(seen, next)
					newNexts = append(newNexts, []node{
						{x: next.x - 1, y: next.y},
						{x: next.x + 1, y: next.y},
						{x: next.x, y: next.y - 1},
						{x: next.x, y: next.y + 1}}...)
				}
			}
			l++
			if len(newNexts) == 0 {
				break
			}
			nexts = newNexts
		}
	}
	res := 0
	nexts := []state23{{n: node{x: 1, y: 0}, length: 0, seen: []node{{x: 1, y: 0}}}}
	for true {
		var newNexts []state23
		for _, next := range nexts {
			if next.n.x == 139 && next.n.y == 140 {
				res = utils.Max(res, next.length)
				continue
			}
			for neighbour, dist := range nodes[next.n] {
				if !slices.Contains(next.seen, neighbour) {
					newSeen := make([]node, len(next.seen))
					copy(newSeen, next.seen)
					newNexts = append(newNexts, state23{
						n:      neighbour,
						length: next.length + dist,
						seen:   append(newSeen, neighbour),
					})
				}
			}
		}
		if len(newNexts) == 0 {
			break
		}
		nexts = newNexts
	}
	fmt.Println(res)
}
