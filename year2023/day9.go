package year2023

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fdfzcq/advent-of-code-go/utils"
)

func Day9Part1() {
	input := utils.ReadStringFromFile(2023, 9)
	hists := parseInputDay9(input)
	res := 0
	for _, hist := range hists {
		res = res + findNextPred(hist)
	}
	fmt.Println(res)
}

func findNextPred(nums []int) int {
	var history [][]int
	history = append(history, nums)
	for i := 0; i >= 0; i++ {
		newHist := make([]int, len(history[i])-1)
		cont := false
		for n := 1; n < len(history[i]); n++ {
			newHist[n-1] = history[i][n] - history[i][n-1]
			if newHist[n-1] != 0 {
				cont = true
			}
		}
		history = append(history, newHist)
		if !cont {
			break
		}
	}
	v := 0
	for i := len(history) - 1; i >= 0; i-- {
		v = history[i][len(history[i])-1] + v
	}
	return v
}

func findPrevPred(nums []int) int {
	var history [][]int
	history = append(history, nums)
	for i := 0; i >= 0; i++ {
		newHist := make([]int, len(history[i])-1)
		cont := false
		for n := 1; n < len(history[i]); n++ {
			newHist[n-1] = history[i][n] - history[i][n-1]
			if newHist[n-1] != 0 {
				cont = true
			}
		}
		history = append(history, newHist)
		if !cont {
			break
		}
	}
	v := 0
	for i := len(history) - 1; i >= 0; i-- {
		v = history[i][0] - v
	}
	return v
}

func parseInputDay9(input string) [][]int {
	lines := strings.Split(input, "\n")
	hists := make([][]int, len(lines))
	for i, line := range lines {
		numStrs := strings.Split(line, " ")
		nums := make([]int, len(numStrs))
		for n, numStr := range numStrs {
			v, _ := strconv.Atoi(numStr)
			nums[n] = v
		}
		hists[i] = nums
	}
	return hists
}

func Day9Part2() {
	input := utils.ReadStringFromFile(2023, 9)
	hists := parseInputDay9(input)
	res := 0
	for _, hist := range hists {
		res = res + findPrevPred(hist)
	}
	fmt.Println(res)
}
