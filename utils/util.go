package utils

import (
	"reflect"
	"strings"
)

type Pair struct {
	A, B interface{}
}

func ParseGrid(input string) [][]string {
	lines := strings.Split(input, "\n")
	grid := make([][]string, len(lines))
	for y, line := range lines {
		chars := strings.Split(line, "")
		row := make([]string, len(chars))
		for x, char := range chars {
			row[x] = char
		}
		grid[y] = row
	}
	return grid
}

func (p Pair) GetA() int {
	return reflect.ValueOf(p.A).Interface().(int)
}

func (p Pair) GetB() int {
	return reflect.ValueOf(p.B).Interface().(int)
}
