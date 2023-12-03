package year2023

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/fdfzcq/advent-of-code-go/utils"
)

var symbols map[utils.Pair]interface{} = make(map[utils.Pair]interface{})
var numbers map[utils.Pair]int = make(map[utils.Pair]int)

func Day3Part1() {
	input := utils.ReadStringFromFile(2023, 3)
	parseInputDay3(input)
	acc := 0
	for p, len := range numbers {
		number := reflect.ValueOf(p.A).Interface().(int)
		coor := reflect.ValueOf(p.B).Interface().(utils.Pair)
		if isPartNumber(number, coor, len, symbols) {
			acc = acc + number
		}
	}
	fmt.Print(acc)
}

func Day3Part2() {
	input := utils.ReadStringFromFile(2023, 3)
	parseInputDay3(input)
	acc := 0
	for p, len := range numbers {
		number := reflect.ValueOf(p.A).Interface().(int)
		coor := reflect.ValueOf(p.B).Interface().(utils.Pair)
		isPartNumber(number, coor, len, symbols)
	}
	for _, v := range symbols {
		switch v := v.(type) {
		case []int:
			if len(v) == 2 {
				acc = acc + v[0]*v[1]
			}
		}
	}
	fmt.Print(acc)
}

func isPartNumber(n int, coor utils.Pair, length int, symbols map[utils.Pair]interface{}) bool {
	res := false
	for i := 0; i < length; i++ {
		x := reflect.ValueOf(coor.A).Interface().(int) - i
		y := reflect.ValueOf(coor.B).Interface().(int)
		if hasSymbol(x-1, y, symbols, n) || hasSymbol(x+1, y, symbols, n) || hasSymbol(x, y-1, symbols, n) ||
			hasSymbol(x, y+1, symbols, n) || hasSymbol(x-1, y-1, symbols, n) || hasSymbol(x-1, y+1, symbols, n) ||
			hasSymbol(x+1, y+1, symbols, n) || hasSymbol(x+1, y-1, symbols, n) {
			res = true
			break
		}
	}
	return res
}

func hasSymbol(x int, y int, symbols map[utils.Pair]interface{}, n int) bool {
	v, ok := symbols[utils.Pair{A: x, B: y}]
	if ok {
		switch v := v.(type) {
		case string:
			if v == "*" {
				symbols[utils.Pair{A: x, B: y}] = []int{n}
			}
		case []int:
			symbols[utils.Pair{A: x, B: y}] = append(v, n)
		}
	}
	return ok
}

func parseInputDay3(input string) {
	lines := strings.Split(input, "\n")
	numStr := ""
	numLen := 0
	p := utils.Pair{}
	for n, row := range lines {
		for m, char := range row {
			c := string(char)
			isNumber, _ := regexp.MatchString("[0-9]", c)
			if !isNumber && c != "." {
				symbols[utils.Pair{A: m, B: n}] = c
			} else if isNumber {
				numStr = numStr + c
				numLen = 1 + numLen
				p = utils.Pair{A: m, B: n}
			}
			if !isNumber && numLen != 0 {
				num, _ := strconv.Atoi(numStr)
				numbers[utils.Pair{A: num, B: p}] = numLen
				numStr = ""
				numLen = 0
				p = utils.Pair{}
			}
		}
	}
	if numLen != 0 {
		num, _ := strconv.Atoi(numStr)
		numbers[utils.Pair{A: num, B: p}] = numLen
	}
}
