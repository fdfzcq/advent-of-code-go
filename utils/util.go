package utils

import (
	"math"
	"reflect"
	"strings"
)

type Pair struct {
	A, B interface{}
}

type Tuple struct {
	X, Y, Z int
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

func (p Pair) GetAString() string {
	return reflect.ValueOf(p.A).Interface().(string)
}

func (p Pair) GetATuple() Tuple {
	return reflect.ValueOf(p.A).Interface().(Tuple)
}

func (p Pair) GetB() int {
	return reflect.ValueOf(p.B).Interface().(int)
}

func (p Pair) GetBTuple() Tuple {
	return reflect.ValueOf(p.B).Interface().(Tuple)
}

func (p Pair) GetBList() []int {
	return reflect.ValueOf(p.B).Interface().([]int)
}

func Min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func Max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func abs(a, b int) int {
	return int(math.Abs(float64(b - a)))
}
