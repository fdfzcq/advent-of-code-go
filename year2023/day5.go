package year2023

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/fdfzcq/advent-of-code-go/utils"
)

type almanac struct {
	seeds         []int
	seedSoil      []resMap
	soilFert      []resMap
	fertWater     []resMap
	waterLight    []resMap
	lightTemp     []resMap
	tempHumid     []resMap
	humidLocation []resMap
}

type resMap struct {
	from int
	to   int
	rang int
}

type rang struct {
	v    int
	rang int
}

func Day5Part1() {
	input := utils.ReadStringFromFile(2023, 5)
	alma := parseInputDay5(input)
	res := 99999999999
	for _, seed := range alma.seeds {
		soil := readFromResMapWithDefault(alma.seedSoil, seed, seed)
		fert := readFromResMapWithDefault(alma.soilFert, soil, soil)
		water := readFromResMapWithDefault(alma.fertWater, fert, fert)
		light := readFromResMapWithDefault(alma.waterLight, water, water)
		temp := readFromResMapWithDefault(alma.lightTemp, light, light)
		humid := readFromResMapWithDefault(alma.tempHumid, temp, temp)
		location := readFromResMapWithDefault(alma.humidLocation, humid, humid)
		res = int(math.Min(float64(res), float64(location)))
	}
	fmt.Println(res)
}

func parseInputDay5(input string) almanac {
	paras := strings.Split(input, "\n\n")
	seedsStr := strings.Split(paras[0], ": ")[1]
	seedStrs := strings.Split(seedsStr, " ")
	seeds := make([]int, len(seedStrs))
	for i, s := range seedStrs {
		n, _ := strconv.Atoi(s)
		seeds[i] = n
	}
	return almanac{
		seeds:         seeds,
		seedSoil:      parseMapDay5(paras[1]),
		soilFert:      parseMapDay5(paras[2]),
		fertWater:     parseMapDay5(paras[3]),
		waterLight:    parseMapDay5(paras[4]),
		lightTemp:     parseMapDay5(paras[5]),
		tempHumid:     parseMapDay5(paras[6]),
		humidLocation: parseMapDay5(paras[7]),
	}
}

func readFromResMapWithDefault(maps []resMap, k int, d int) int {
	res := d
	for _, m := range maps {
		if k >= m.from && k <= m.from+m.rang {
			res = k - m.from + m.to
		}
	}
	return res
}

func parseMapDay5(input string) []resMap {
	lines := strings.Split(input, "\n")
	res := make([]resMap, len(lines)-1)
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		numStrs := strings.Split(line, " ")
		toNum, _ := strconv.Atoi(numStrs[0])
		froNum, _ := strconv.Atoi(numStrs[1])
		rang, _ := strconv.Atoi(numStrs[2])
		res[i-1] = resMap{from: froNum, to: toNum, rang: rang}
	}
	return res
}

func Day5Part2() {
	input := utils.ReadStringFromFile(2023, 5)
	alma := parseInputDay5(input)
	res := 99999999999
	i := 0
	for i+1 < len(alma.seeds) {
		seedRange := rang{v: alma.seeds[i], rang: alma.seeds[i+1]}
		soilR := overlapResMap(seedRange.v, seedRange.rang, alma.seedSoil)
		fertR := overlapResMaps(soilR, alma.soilFert)
		waterR := overlapResMaps(fertR, alma.fertWater)
		lightR := overlapResMaps(waterR, alma.waterLight)
		tempR := overlapResMaps(lightR, alma.lightTemp)
		humidR := overlapResMaps(tempR, alma.tempHumid)
		locationR := overlapResMaps(humidR, alma.humidLocation)
		sort.SliceStable(locationR, func(x, y int) bool {
			return locationR[x].v < locationR[y].v
		})
		res = int(math.Min(float64(res), float64(locationR[0].v)))

		i = i + 2
	}
	fmt.Println(res)
}

func overlapResMaps(ranges []rang, ms []resMap) []rang {
	var res []rang
	for _, r := range ranges {
		res = append(res, overlapResMap(r.v, r.rang, ms)...)
	}
	return res
}

func overlapResMap(start int, r int, ms []resMap) []rang {
	res := make([]rang, 0)
	sort.SliceStable(ms, func(i, j int) bool {
		return ms[i].from < ms[j].from
	})
	ceiling := start

	for _, m := range ms {
		if start > m.from+m.rang {
			continue
		}
		if start+r < m.from {
			res = append(res, rang{v: ceiling, rang: start + r - ceiling})
			ceiling = start + r
			break
		} else if start < m.from && start+r-1 < m.from+m.rang {
			res = append(res, rang{v: m.to, rang: start + r - m.from})
			res = append(res, rang{v: ceiling, rang: m.from - ceiling})
			ceiling = start + r
			break
		} else if start < m.from {
			res = append(res, rang{v: m.to, rang: m.rang})
			res = append(res, rang{v: ceiling, rang: m.from - ceiling})
			ceiling = m.from + m.rang
		} else if start >= m.from && start+r-1 < m.from+m.rang {
			res = append(res, rang{v: start - m.from + m.to, rang: r})
			ceiling = start + r
			break
		} else if start >= m.from {
			res = append(res, rang{v: start - m.from + m.to, rang: m.from + m.rang - start})
			ceiling = m.from + m.rang
		}
	}

	if ceiling < start+r {
		res = append(res, rang{v: ceiling, rang: start + r - ceiling})
	}

	var newRes []rang
	for _, r := range res {
		if r.rang != 0 {
			newRes = append(newRes, r)
		}
	}

	return newRes
}
