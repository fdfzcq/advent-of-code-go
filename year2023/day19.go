package year2023

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fdfzcq/advent-of-code-go/utils"
)

type workflow struct {
	rules       []rule
	defaultRule string
}

type rule struct {
	source      string
	operator    string
	number      int
	destination string
}

type rule2 struct {
	rmap        map[string][]utils.Pair
	destination string
}

func Day19Part1() {
	inputStr := utils.ReadStringFromFile(2023, 19)
	sections := strings.Split(inputStr, "\n\n")
	workflows := parseWorkflows(sections[0])

	res := 0

	for _, input := range strings.Split(sections[1], "\n") {
		valueStrs := strings.Split(input[1:len(input)-1], ",")
		m := make(map[string]int)
		for _, valueStr := range valueStrs {
			parts := strings.Split(valueStr, "=")
			n, _ := strconv.Atoi(parts[1])
			m[parts[0]] = n
		}
		if checkWorkflows(m, workflows) == "A" {
			for _, mv := range m {
				res = res + mv
			}
		}
	}

	fmt.Println(res)
}

func checkWorkflows(m map[string]int, workflows map[string]workflow) string {
	cur := "in"
	for true {
		if cur == "A" || cur == "R" {
			return cur
		}
		workflow := workflows[cur]
		found := false
		for _, r := range workflow.rules {
			if fulfills(m, r) {
				cur = r.destination
				found = true
				break
			}
		}
		if !found {
			cur = workflow.defaultRule
		}
	}
	return "R"
}

func fulfills(m map[string]int, r rule) bool {
	v, _ := m[r.source]
	c := r.number
	if r.operator == ">" {
		return v > c
	} else if r.operator == "<" {
		return v < c
	}
	return false
}

func parseWorkflows(input string) map[string]workflow {
	workflows := make(map[string]workflow)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line[0:len(line)-1], "{")
		name := parts[0]
		ruleStrs := strings.Split(parts[1], ",")
		var rules []rule
		var defaultRule string
		for i, ruleStr := range ruleStrs {
			if i == len(ruleStrs)-1 {
				defaultRule = ruleStr
			} else {
				ps := strings.Split(ruleStr, ":")
				n, _ := strconv.Atoi(ps[0][2:len(ps[0])])
				rules = append(rules, rule{
					source:      ps[0][0:1],
					operator:    ps[0][1:2],
					number:      n,
					destination: ps[1],
				})
			}
		}
		workflows[name] = workflow{
			rules:       rules,
			defaultRule: defaultRule,
		}
	}
	return workflows
}

func Day19Part2() {
	inputStr := utils.ReadStringFromFile(2023, 19)
	sections := strings.Split(inputStr, "\n\n")
	workflows := parseWorkflows(sections[0])

	rA := []utils.Pair{{A: 1, B: 4000}}
	rM := []utils.Pair{{A: 1, B: 4000}}
	rX := []utils.Pair{{A: 1, B: 4000}}
	rS := []utils.Pair{{A: 1, B: 4000}}
	rmap := map[string][]utils.Pair{"a": rA, "m": rM, "x": rX, "s": rS}
	queue := []rule2{{rmap: rmap, destination: "in"}}

	res := 0
	for true {
		var newQueue []rule2
		for _, q := range queue {
			rules := workflows[q.destination].rules
			newRR := copyMap(q.rmap)
			if q.destination == "R" {
				continue
			} else if q.destination == "A" {
				acc := 1
				for _, ps := range q.rmap {
					a := 0
					for _, p := range ps {
						a = a + (p.GetB() - p.GetA() + 1)
					}
					acc = acc * a
				}
				res = res + acc
			} else {
				for _, r := range rules {
					newR := copyMap(newRR)
					rr := newR[r.source]
					n := r.number
					isValid := false
					for i, vr := range rr {
						if r.operator == "<" {
							if vr.GetB() < n {
								isValid = true
								newRR[r.source] = mergeRange(newRR[r.source], utils.Pair{A: n, B: 4001})
							} else if vr.GetA() < n {
								isValid = true
								rr[i] = utils.Pair{A: vr.GetA(), B: n - 1}
								newRR[r.source] = mergeRange(newRR[r.source], utils.Pair{A: n, B: vr.GetB()})
							} else {
								newRR[r.source] = mergeRange(newRR[r.source], vr)
							}
						} else {
							if vr.GetA() > n {
								isValid = true
								newRR[r.source] = mergeRange(newRR[r.source], utils.Pair{A: 1, B: n})
							} else if vr.GetB() > n {
								isValid = true
								rr[i] = utils.Pair{A: n + 1, B: vr.GetB()}
								newRR[r.source] = mergeRange(newRR[r.source], utils.Pair{A: vr.GetA(), B: n})
							} else {
								newRR[r.source] = mergeRange(newRR[r.source], vr)
							}
						}
					}
					if isValid {
						newQueue = append(newQueue, rule2{rmap: newR, destination: r.destination})
					}
				}
				isValid := true
				for _, s := range []string{"a", "m", "x", "s"} {
					l, ok := newRR[s]
					isValid = ok && len(l) != 0
				}
				if isValid {
					newQueue = append(newQueue, rule2{rmap: newRR, destination: workflows[q.destination].defaultRule})
				}
			}
		}
		if len(newQueue) == 0 {
			break
		}
		queue = newQueue
	}

	fmt.Println(res)
}

func copyMap(m map[string][]utils.Pair) map[string][]utils.Pair {
	newR := make(map[string][]utils.Pair)
	for _, s := range []string{"a", "m", "x", "s"} {
		list := make([]utils.Pair, len(m[s]))
		copy(list, m[s])
		newR[s] = list
	}
	return newR
}

func mergeRange(ls []utils.Pair, r utils.Pair) []utils.Pair {
	var newl []utils.Pair
	for _, l := range ls {
		if r.GetA() <= l.GetB() && r.GetB() >= l.GetA() {
			newl = append(newl, utils.Pair{A: utils.Max(r.GetA(), l.GetA()), B: utils.Min(r.GetB(), l.GetB())})
		}
	}
	return newl
}
