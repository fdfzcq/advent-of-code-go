package year2023

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/fdfzcq/advent-of-code-go/utils"
)

var strToNumMap = map[string]string{
	"one":   "one1one",
	"two":   "two2two",
	"three": "three3three",
	"four":  "four4four",
	"five":  "five5five",
	"six":   "six6six",
	"seven": "seven7seven",
	"eight": "eight8eight",
	"nine":  "nine9nine",
}

func Day1Part1() {
	lines := utils.ReadStringFromFile(2023, 1)
	list := strings.Split(lines, "\n")
	part1(list)
}

func part1(list []string) {
	acc := 0
	for _, str := range list {
		re := regexp.MustCompile("[0-9]")
		matches := re.FindAllString(str, -1)
		n1, _ := strconv.Atoi(matches[0])
		n2, _ := strconv.Atoi(matches[len(matches)-1])
		acc = acc + 10*n1 + n2
	}
	fmt.Print(acc)
}

func Day1Part2() {
	lines := utils.ReadStringFromFile(2023, 1)
	list := strings.Split(lines, "\n")
	newlist := make([]string, len(list))
	for i, str := range list {
		for k, v := range strToNumMap {
			str = strings.ReplaceAll(str, k, v)
		}
		newlist[i] = str
	}
	part1(newlist)
}
