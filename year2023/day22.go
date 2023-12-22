package year2023

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/fdfzcq/advent-of-code-go/utils"
)

func Day22Part1() {
	input := utils.ReadStringFromFile(2023, 22)
	var bricks []utils.Pair
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, "~")
		var ns []int
		for _, s := range strings.Split(parts[0], ",") {
			n, _ := strconv.Atoi(s)
			ns = append(ns, n)
		}
		for _, s := range strings.Split(parts[1], ",") {
			n, _ := strconv.Atoi(s)
			ns = append(ns, n)
		}
		bricks = append(bricks, utils.Pair{
			A: utils.Tuple{X: utils.Min(ns[0], ns[3]), Y: utils.Min(ns[1], ns[4]), Z: utils.Min(ns[2], ns[5])},
			B: utils.Tuple{X: utils.Max(ns[0], ns[3]), Y: utils.Max(ns[1], ns[4]), Z: utils.Max(ns[2], ns[5])}})
	}
	slices.SortFunc(bricks, func(a, b utils.Pair) int {
		return a.GetATuple().Z - b.GetATuple().Z
	})
	grid := make(map[int]map[utils.Pair]int)
	deps := make(map[int][]int)
	revDeps := make(map[int][]int)
	for i, b := range bricks {
		zMin := b.GetATuple().Z
		newZ := 1
		for iz := zMin; iz > 0; iz-- {
			xys, ok := grid[iz]
			if ok {
				l := overlaps(b, xys)
				if len(l) != 0 {
					newZ = iz + 1
					for _, n := range l {
						ll, ok := deps[i]
						if !ok {
							ll = []int{n}
							deps[i] = ll
						} else if !slices.Contains(ll, n) {
							deps[i] = append(ll, n)
						}
						ll, ok = revDeps[n]
						if !ok {
							ll = []int{i}
							revDeps[n] = ll
						} else if !slices.Contains(ll, i) {
							revDeps[n] = append(ll, i)
						}
					}
					break
				}
			}
		}
		for z := newZ; z <= newZ+b.GetBTuple().Z-b.GetATuple().Z; z++ {
			xys, ok := grid[z]
			if !ok {
				xys = make(map[utils.Pair]int)
				grid[z] = xys
			}
			for y := b.GetATuple().Y; y <= b.GetBTuple().Y; y++ {
				for x := b.GetATuple().X; x <= b.GetBTuple().X; x++ {
					xys[utils.Pair{A: x, B: y}] = i
				}
			}
		}
	}
	res := 0
	for i := 0; i < len(bricks); i++ {
		standAlone := true
		for _, l := range deps {
			if len(l) == 1 && l[0] == i {
				standAlone = false
				break
			}
		}
		if standAlone {
			res++
		}
	}
	fmt.Print("part1: ")
	fmt.Println(res)

	res2 := 0
	for i := len(bricks) - 1; i >= 0; i-- {
		nexts, ok := revDeps[i]
		removed := []int{i}
		for true {
			var newNexts []int
			if !ok {
				break
			} else {
				for _, d := range nexts {
					allRemoved := true
					for _, t := range deps[d] {
						if !slices.Contains(removed, t) {
							allRemoved = false
						}
					}
					if allRemoved {
						dps, ok := revDeps[d]
						if ok {
							for _, n := range dps {
								if !slices.Contains(newNexts, n) {
									newNexts = append(newNexts, n)
								}
							}
						}
						if !slices.Contains(removed, d) {
							removed = append(removed, d)
						}
					}
				}
			}
			if len(newNexts) == 0 {
				break
			}
			nexts = newNexts
		}
		res2 = res2 + len(removed) - 1
	}
	fmt.Println(res2)
}

func overlaps(p utils.Pair, grid map[utils.Pair]int) []int {
	var l []int
	for y := p.GetATuple().Y; y <= p.GetBTuple().Y; y++ {
		for x := p.GetATuple().X; x <= p.GetBTuple().X; x++ {
			v, ok := grid[utils.Pair{A: x, B: y}]
			if ok {
				l = append(l, v)
			}
		}
	}
	return l
}
