package year2023

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/fdfzcq/advent-of-code-go/utils"
)

func Day15Part1() {
	input := utils.ReadStringFromFile(2023, 15)
	strings := strings.Split(input, ",")
	res := 0
	for _, str := range strings {
		v := 0
		for _, r := range str {
			v = (v + int(r)) * 17 % 256
		}
		res = res + v
	}
	fmt.Print(res)
}

func Day15Part2() {
	input := utils.ReadStringFromFile(2023, 15)
	strs := strings.Split(input, ",")
	boxes := make(map[int][]utils.Pair)
	mem := make(map[string]int)
	vmem := make(map[string]int)
	for _, str := range strs {
		if strings.HasSuffix(str, "-") {
			str = strings.Split(str, "-")[0]
			i, ok := mem[str]
			if ok {
				v, ok := vmem[str]
				if ok {
					boxes[i] = slices.DeleteFunc(boxes[i], func(p utils.Pair) bool {
						return p == utils.Pair{A: str, B: v}
					})
					delete(vmem, str)
				}
			}
		} else {
			parts := strings.Split(str, "=")
			n, _ := strconv.Atoi(parts[1])
			i, ok := mem[parts[0]]
			if !ok {
				v := 0
				for _, r := range parts[0] {
					v = (v + int(r)) * 17 % 256
				}
				i = v
				mem[parts[0]] = i
			}
			v, ok := vmem[parts[0]]
			if !ok {
				boxes[i] = append(boxes[i], utils.Pair{A: parts[0], B: n})
			} else {
				x := slices.IndexFunc(boxes[i], func(p utils.Pair) bool {
					return p == utils.Pair{A: parts[0], B: v}
				})
				boxes[i][x] = utils.Pair{A: parts[0], B: n}
			}
			vmem[parts[0]] = n
		}
	}
	res := 0
	for i, b := range boxes {
		for j, p := range b {
			res = res + (i+1)*(j+1)*p.GetB()
		}
	}
	fmt.Print(res)
}
