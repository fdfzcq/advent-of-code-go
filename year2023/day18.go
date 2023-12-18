package year2023

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/fdfzcq/advent-of-code-go/utils"
)

func toDir(d string) (int, int) {
	if d == "U" || d == "3" {
		return 0, -1
	}
	if d == "D" || d == "1" {
		return 0, 1
	}
	if d == "R" || d == "0" {
		return 1, 0
	}
	return -1, 0
}

func Day18() {
	input := utils.ReadStringFromFile(2023, 18)
	minX := 0
	minY := 0
	maxX := 0
	maxY := 0
	p := utils.Pair{A: 0, B: 0}
	g := make(map[int][]utils.Pair)
	for b, l := range strings.Split(input, "\n") {
		parts := strings.Split(l, " ")
		n, dir := parseInstruction(parts[2])
		dirX, dirY := toDir(dir)
		for i := 1; i <= n; i++ {
			x := p.GetA() + dirX
			y := p.GetB() + dirY
			minX = utils.Min(minX, x)
			minY = utils.Min(minY, y)
			maxX = utils.Max(maxX, x)
			maxY = utils.Max(maxY, y)
			p = utils.Pair{A: x, B: y}
			rs, ok := g[y]
			if ok {
				found := false
				for k, r := range rs {
					if x == r.GetA()-1 {
						rs[k] = utils.Pair{A: x, B: r.GetB()}
						found = true
						break
					} else if x == r.GetB()+1 {
						rs[k] = utils.Pair{A: r.GetA(), B: x}
						found = true
						break
					}
				}
				if !found {
					rs = append(rs, utils.Pair{A: x, B: x})
				}
			} else {
				rs = []utils.Pair{{A: x, B: x}}
			}
			g[y] = rs
		}
		fmt.Println(b)
	}
	fmt.Println("res")

	xs := make(map[int]bool)
	res := 0
	extra := 0
	for y := minY; y <= maxY; y++ {
		if y%1000 == 0 {
			fmt.Println(y)
		}
		slices.SortFunc(g[y], func(a, b utils.Pair) int {
			return a.GetA() - b.GetA()
		})
		for _, k := range g[y] {
			for x := k.GetA(); x <= k.GetB(); x++ {
				_, ok0 := xs[x]
				ok := isIn(g[y], x)
				if ok {
					res++
					ok1 := isIn(g[y], x-1)
					ok2 := isIn(g[y], x+1)
					_, ok3 := xs[x-1]
					_, ok4 := xs[x+1]
					if ok0 && ((ok1 && ok2) || (ok1 && ok4) || (ok2 && ok3)) {
						delete(xs, x)
					} else if (ok1 && ok2) || (ok1 && ok4) || (ok2 && ok3) {
						xs[x] = true
						extra++
					} else if ok0 {
						extra++
					}
				}
			}
		}
		res = res + len(xs) - extra
		extra = 0
	}

	fmt.Println(res)
}

func isIn(rs []utils.Pair, v int) bool {
	for _, r := range rs {
		if v >= r.GetA() && v <= r.GetB() {
			return true
		}
	}
	return false
}

func parseInstruction(ins string) (int, string) {
	distStr := ins[2:7]
	dist, _ := strconv.ParseUint(distStr, 16, 64)
	dirStr := ins[7:8]
	return int(dist), dirStr
}
