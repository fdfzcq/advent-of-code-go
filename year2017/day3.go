package year2017

import (
	"fmt"
	"math"
	"strconv"

	"github.com/fdfzcq/advent-of-code-go/utils"
)

func Day3Part1() {
	num, _ := strconv.Atoi(utils.ReadStringFromFile(2017, 3))
	i := 1
	start := 2
	end := 9
	for end < num {
		i++
		newEnd := end + end - start + 9
		start = end + 1
		end = newEnd
	}
	// start 10 end 25
	// end - start + 1 / 4
	n := start - 1
	for true {
		if num <= n+((end-start+1)/4) {
			fmt.Print(math.Abs(float64(n+((end-start+1)/8)-num)) + float64(i))
			break
		}
		n = n + ((end - start + 1) / 4)
	}
}
