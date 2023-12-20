package year2023

import (
	"fmt"
	"strings"

	"github.com/fdfzcq/advent-of-code-go/utils"
)

type input struct {
	source      string
	destination string
	isHigh      bool
}

type module interface {
	process(from string, impulse bool) []input
	connect(m string)
	dests() []string
	on() bool
}

type flipflop struct {
	name         string
	isOn         bool
	destinations []string
}

type conjuction struct {
	name         string
	mem          map[string]bool
	destinations []string
}

type broadcaster struct {
	name         string
	destinations []string
}

func (m *flipflop) process(_ string, impulse bool) []input {
	if impulse {
		return []input{}
	} else {
		if !m.isOn {
			m.isOn = true
			var res []input
			for _, d := range m.destinations {
				res = append(res, input{source: m.name, destination: d, isHigh: true})
			}
			return res
		} else {
			m.isOn = false
			var res []input
			for _, d := range m.destinations {
				res = append(res, input{source: m.name, destination: d, isHigh: false})
			}
			return res
		}
	}
}

func (m *flipflop) connect(from string) {
}

func (m *flipflop) dests() []string {
	return m.destinations
}

func (m *flipflop) on() bool {
	return m.isOn
}

func (m *conjuction) process(from string, impulse bool) []input {
	m.mem[from] = impulse
	allHigh := true
	for _, v := range m.mem {
		if !v {
			allHigh = false
			break
		}
	}
	if allHigh {
		var res []input
		for _, d := range m.destinations {
			res = append(res, input{source: m.name, destination: d, isHigh: false})
		}
		return res
	} else {
		var res []input
		for _, d := range m.destinations {
			res = append(res, input{source: m.name, destination: d, isHigh: true})
		}
		return res
	}
}

func (m *conjuction) connect(from string) {
	m.mem[from] = false
}

func (m *conjuction) dests() []string {
	return m.destinations
}

func (m *conjuction) on() bool {
	return false
}

func (m *broadcaster) process(from string, impulse bool) []input {
	var outputs []input
	for _, d := range m.destinations {
		outputs = append(outputs, input{source: m.name, destination: d, isHigh: impulse})
	}
	return outputs
}

func (m *broadcaster) connect(from string) {
}

func (m *broadcaster) dests() []string {
	return m.destinations
}

func (m *broadcaster) on() bool {
	return false
}

func Day20() {
	inputStr := utils.ReadStringFromFile(2023, 20)
	fmt.Println(process(parseModules(inputStr), "bm") *
		process(parseModules(inputStr), "cl") *
		process(parseModules(inputStr), "tn") *
		process(parseModules(inputStr), "dr"))
}

func process(modules map[string]module, source string) int {
	i := 1
	for true {
		newOnC, _, _ := processOnce(modules, source)
		if newOnC == -1 {
			fmt.Println(i)
			break
		}
		i++
	}
	return i
}

func processOnce(modules map[string]module, source string) (int, int, int) {
	lowC := 1
	highC := 0
	inputs := []input{{source: "", destination: "broadcaster", isHigh: false}}
	t := false
	for true {
		var newInputs []input
		for _, i := range inputs {
			if i.source == source && i.isHigh && i.destination == "vr" {
				t = true
			}
			mod, ok := modules[i.destination]
			if ok {
				outputs := mod.process(i.source, i.isHigh)
				for _, o := range outputs {
					if o.isHigh {
						highC++
					} else {
						lowC++
					}
				}
				newInputs = append(newInputs, outputs...)
			} else {
				if !i.isHigh {
					return -1, 0, 0
				}
			}
		}
		if len(newInputs) == 0 {
			break
		}
		inputs = newInputs
	}
	onC := 0
	for _, m := range modules {
		if m.on() {
			onC++
		}
	}
	if t {
		return -1, 0, 0
	}
	return onC, lowC, highC
}

func parseModules(input string) map[string]module {
	m := make(map[string]module)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " -> ")
		if parts[0] == "broadcaster" {
			m["broadcaster"] = &broadcaster{
				name:         "broadcaster",
				destinations: strings.Split(parts[1], ", "),
			}
		} else if parts[0][0:1] == "%" {
			name := parts[0][1:len(parts[0])]
			m[name] = &flipflop{
				name:         name,
				destinations: strings.Split(parts[1], ", "),
				isOn:         false,
			}
		} else {
			name := parts[0][1:len(parts[0])]
			m[name] = &conjuction{
				name:         name,
				destinations: strings.Split(parts[1], ", "),
				mem:          make(map[string]bool),
			}
		}
	}
	for name, mod := range m {
		for _, d := range mod.dests() {
			mod, ok := m[d]
			if ok {
				mod.connect(name)
			}
		}
	}
	return m
}
