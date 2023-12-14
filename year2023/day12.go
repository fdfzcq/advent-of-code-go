package year2023

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fdfzcq/advent-of-code-go/utils"
)

type state struct {
	current bool
	i       int
	l       int
}

func Day12Part1() {
	input := utils.ReadStringFromTestFile(2023, 12)
	c := make(chan int)
	res := 0
	for _, line := range strings.Split(input, "\n") {
		go func(line string) {
			c <- arrangements(line)
		}(line)
	}
	for v := range c {
		res = res + v
		fmt.Println(res)
	}
}

func arrangements(line string) int {
	parts := strings.Split(line, " ")
	numStrs := strings.Split(parts[1], ",")
	nums := make([]int, len(numStrs))
	for i, numStr := range numStrs {
		n, _ := strconv.Atoi(numStr)
		nums[i] = n
	}
	hs := strings.Split(parts[0], "")

	states := make(map[state]int)

	for i, e := range hs {
		if i == 0 {
			if e == "?" {
				s0 := state{current: true, i: 0, l: 1}
				s1 := state{current: false, i: 0, l: 0}
				states[s0] = 1
				states[s1] = 1
			} else if e == "#" {
				states[state{current: true, i: 0, l: 1}] = 1
			} else {
				states[state{current: false, i: 0, l: 0}] = 1
			}
		} else {
			newstates := make(map[state]int)
			for s, v := range states {
				if e == "#" {
					s.l = s.l + 1
					s.current = true
					if s.i < len(nums) && s.l <= nums[s.i] {
						val, ok := newstates[s]
						if !ok {
							newstates[s] = v
						} else {
							newstates[s] = v + val
						}
					}
				}
				if e == "." {
					if s.current {
						if s.i < len(nums) && s.l == nums[s.i] {
							s.current = false
							s.i = s.i + 1
							s.l = 0
							val, ok := newstates[s]
							if !ok {
								newstates[s] = v
							} else {
								newstates[s] = v + val
							}
						}
					} else {
						val, ok := newstates[s]
						if !ok {
							newstates[s] = v
						} else {
							newstates[s] = v + val
						}
					}
				}
				if e == "?" {
					news := s
					if s.i < len(nums) && s.l+1 <= nums[s.i] {
						s.l = s.l + 1
						s.current = true
						val, ok := newstates[s]
						if !ok {
							newstates[s] = v
						} else {
							newstates[s] = v + val
						}
					}
					if news.current {
						if news.i < len(nums) && news.l == nums[news.i] {
							news.current = false
							news.i = news.i + 1
							news.l = 0
							val, ok := newstates[news]
							if !ok {
								newstates[news] = v
							} else {
								newstates[news] = v + val
							}
						}
					} else {
						val, ok := newstates[news]
						if !ok {
							newstates[news] = v
						} else {
							newstates[news] = v + val
						}
					}
				}
			}
			states = newstates
		}
	}
	res := 0

	for s, v := range states {
		if (s.i == len(nums) && s.l == 0) || (s.i == len(nums)-1 && s.l == nums[s.i]) {
			res = res + v
		}
	}

	return res
}

func Day12Part2() {
	input := utils.ReadStringFromFile(2023, 12)
	c := make(chan int)
	res := uint64(0)
	for _, line := range strings.Split(input, "\n") {
		go func(line string) {
			parts := strings.Split(line, " ")
			// a0 := arrangements(parts[0] + "?" + parts[0] + " " + parts[1] + "," + parts[1])
			// a := arrangements(line)
			// c <- a * (a0 / a) * (a0 / a) * (a0 / a) * (a0 / a)
			c <- arrangements(parts[0] + "?" + parts[0] + "?" + parts[0] + "?" + parts[0] + "?" + parts[0] + " " + parts[1] + "," + parts[1] + "," + parts[1] + "," + parts[1] + "," + parts[1])
		}(line)
	}
	for v := range c {
		res = res + uint64(v)
		fmt.Println(res)
	}
}
