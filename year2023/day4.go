package year2023

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/fdfzcq/advent-of-code-go/utils"
)

type card struct {
	id             int
	winningNumbers []int
	numbers        []int
}

func Day4Part1() {
	input := utils.ReadStringFromFile(2023, 4)
	cards := parseInputDay4(input)
	acc := 0
	for _, card := range cards {
		n := 0
		for _, num := range card.numbers {
			if slices.Contains(card.winningNumbers, num) {
				n++
			}
		}
		if n > 0 {
			acc = acc + int(math.Pow(2, float64(n-1)))
		}
	}
}

func parseInputDay4(str string) []card {
	lines := strings.Split(str, "\n")
	cards := make([]card, len(lines))
	for i, line := range lines {
		strs := strings.Split(line, ": ")
		r, _ := regexp.Compile(`[0-9]+`)
		idStr := r.FindAllString(strs[0], -1)
		id, _ := strconv.Atoi(idStr[0])
		numbersStr := strings.Split(strs[1], " | ")
		winNumsStr := r.FindAllString(numbersStr[0], -1)
		numsStr := r.FindAllString(numbersStr[1], -1)
		winNums := make([]int, len(winNumsStr))
		nums := make([]int, len(numsStr))
		for m, wns := range winNumsStr {
			wn, _ := strconv.Atoi(wns)
			winNums[m] = wn
		}
		for n, ns := range numsStr {
			num, _ := strconv.Atoi(ns)
			nums[n] = num
		}
		cards[i] = card{id: id, winningNumbers: winNums, numbers: nums}
	}
	return cards
}

func Day4Part2() {
	input := utils.ReadStringFromFile(2023, 4)
	cards := parseInputDay4(input)
	res := make(map[int]int)
	for i := 0; i < len(cards); i++ {
		card := cards[i]
		v, ok := res[card.id]
		if !ok {
			res[card.id] = 1
		} else {
			res[card.id] = v + 1
		}
		p := 0
		for _, num := range card.numbers {
			if slices.Contains(card.winningNumbers, num) {
				p++
				vv, ok := res[card.id+p]
				if !ok {
					res[card.id+p] = res[card.id]
				} else {
					res[card.id+p] = vv + res[card.id]
				}
			}
		}
	}
	acc := 0
	for _, v := range res {
		acc = acc + v
	}
	fmt.Println(acc)
}
