package year2023

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/fdfzcq/advent-of-code-go/utils"
)

type storm struct {
	px int
	py int
	pz int
	vx int
	vy int
	vz int
}

func Day24Part1() {
	input := utils.ReadStringFromFile(2023, 24)
	var storms []storm
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " @ ")
		var ns []int
		re := regexp.MustCompile("[-0-9]+")
		matches := re.FindAllString(parts[0], -1)
		for _, m := range matches {
			n, _ := strconv.Atoi(m)
			ns = append(ns, n)
		}
		matches = re.FindAllString(parts[1], -1)
		for _, m := range matches {
			n, _ := strconv.Atoi(m)
			ns = append(ns, n)
		}
		storms = append(storms, storm{
			px: ns[0], py: ns[1], pz: ns[2],
			vx: ns[3], vy: ns[4], vz: ns[5],
		})
	}
	fmt.Println(storms)
	res := 0
	for i, st1 := range storms {
		for j := i + 1; j < len(storms); j++ {
			st2 := storms[j]
			a1 := float64(st1.vy) / float64(st1.vx)
			b1 := float64(st1.vx*st1.py-st1.vy*st1.px) / float64(st1.vx)
			a2 := float64(st2.vy) / float64(st2.vx)
			b2 := float64(st2.vx*st2.py-st2.vy*st2.px) / float64(st2.vx)
			x := (b2 - b1) / (a1 - a2)
			y := a1*x + b1
			n1 := (x - float64(st1.px)) / float64(st1.vx)
			n2 := (x - float64(st2.px)) / float64(st2.vx)
			if x >= float64(200000000000000) && x <= float64(400000000000000) && y >= float64(200000000000000) && y <= float64(400000000000000) && n1 >= 0 && n2 >= 0 {
				res++
			}
		}
	}
	fmt.Println(res)
}

type pos struct {
	x int
	y int
	z int
}

func Day24Part2() {
	input := utils.ReadStringFromFile(2023, 24)
	var storms []storm
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " @ ")
		var ns []int
		re := regexp.MustCompile("[-0-9]+")
		matches := re.FindAllString(parts[0], -1)
		for _, m := range matches {
			n, _ := strconv.Atoi(m)
			ns = append(ns, n)
		}
		matches = re.FindAllString(parts[1], -1)
		for _, m := range matches {
			n, _ := strconv.Atoi(m)
			ns = append(ns, n)
		}
		storm := storm{
			px: ns[0], py: ns[1], pz: ns[2],
			vx: ns[3], vy: ns[4], vz: ns[5],
		}
		storms = append(storms, storm)
	}
	acc := 0
	var poss []pos
	var velos []pos
	//var seen []pos

	for i := 0; i < len(storms); i++ {
		for n1 := 0; n1 < 10000000; n1++ {
			for j := i + 1; j < len(storms); j++ {
				for n2 := 0; n2 < 10000000; n2++ {
					if n2 != n1 {
						s1 := storms[i]
						s2 := storms[j]
						p1x := s1.px + n1*s1.vx
						p2x := s2.px + n2*s2.vx
						p1y := s1.py + n1*s1.vy
						p2y := s2.py + n2*s2.vy
						p1z := s1.pz + n1*s1.vz
						p2z := s2.pz + n2*s2.vz
						vx := (p2x - p1x) / (n2 - n1)
						vy := (p2y - p1y) / (n2 - n1)
						vz := (p2z - p1z) / (n2 - n1)
						for t := j + 1; t < len(storms); t++ {
							for n3 := 0; n3 < 10000000; n3++ {
								if acc%100000000 == 0 {
									fmt.Println(acc)
								}
								s3 := storms[t]
								p3x := s3.px + n3*s3.vx
								p3y := s3.py + n3*s3.vy
								p3z := s3.pz + n3*s3.vz
								if p3x == p2x+(n3-n2)*vx && p3y == p2y+(n3-n2)*vy && p3z == p2z+(n3-n2)*vz {
									fmt.Print("vx: ")
									fmt.Print(vx)
									fmt.Print(" vy: ")
									fmt.Print(vy)
									fmt.Print(" vz: ")
									fmt.Print(vz)
									fmt.Print(" start: ")
									fmt.Println(pos{x: p3x - vx*n3, y: p3y - vy, z: p3z - vz})
									poss = append(poss, pos{x: p3x - vx*n3, y: p3y - vy, z: p3z - vz})
									velos = append(velos, pos{x: vx, y: vy, z: vz})
								}
								acc++
							}
						}
					}
				}
			}
		}
	}

	// for i, v := range velos {
	// 	for _, s := range storms {
	// 		pos := poss[i]

	// 	}
	// }
}
