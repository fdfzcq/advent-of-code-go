package year2023

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/fdfzcq/advent-of-code-go/utils"
)

func Day25() {
	input := utils.ReadStringFromFile(2023, 25)
	m := make(map[string][]string)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ": ")
		key := parts[0]
		for _, v := range strings.Split(parts[1], " ") {
			_, ok := m[v]
			if !ok {
				m[v] = make([]string, 0)
			}
			m[v] = append(m[v], key)
			_, ok = m[key]
			if !ok {
				m[key] = make([]string, 0)
			}
			m[key] = append(m[key], v)
		}
	}
	i := 0
	k1 := "tvf"
	v1 := "tqn"
	k2 := "vzb"
	v2 := "tnr"
	for k3, l3 := range m {
		for _, v3 := range l3 {
			if !(k1 == k2 && v1 == v2) && !(k1 == v2 && v1 == k2) && !(k1 == k3 && v1 == v3) &&
				!(k1 == v3 && v1 == k3) && !(k2 == k3 && v2 == v3) && !(k2 == v3 && v2 == k3) {
				groups := make([][]string, 0)
				var seen []string
				acc := 0
				i++
				if i%10 == 0 {
					fmt.Println(i)
				}
				for k, _ := range m {
					if !slices.Contains(seen, k) {
						nexts := []string{k}
						group := []string{}
						for true {
							var newNexts []string
							//fmt.Println(nexts)
							for _, next := range nexts {
								if !slices.Contains(seen, next) {
									seen = append(seen, next)
									group = append(group, next)
									for _, v := range m[next] {
										if !(next == k1 && v == v1) && !(next == k2 && v == v2) && !(next == k3 && v == v3) &&
											!(next == v1 && v == k1) && !(next == v2 && v == k2) && !(next == v3 && v == k3) {
											newNexts = append(newNexts, v)
										}
									}
								}
							}
							if len(newNexts) == 0 {
								break
							}
							nexts = newNexts
						}
						acc++
						if acc > 3 {
							break
						}
						groups = append(groups, group)
					}
				}
				if len(groups) == 2 {
					fmt.Println(len(groups[0]) * len(groups[1]))
					os.Exit(1)
				}
			}

		}
	}
}
