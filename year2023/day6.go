package year2023

import "fmt"

func Day6Part1() {
	//fmt.Println(winTimes(7, 9) * winTimes(15, 40) * winTimes(30, 200))
	fmt.Println(winTimes(48, 261) * winTimes(93, 1192) * winTimes(84, 1019) * winTimes(66, 1063))
}

func winTimes(time int, dist int) int {
	t := 0
	for i := 0; i <= time; i++ {
		if i*(time-i) > dist {
			t++
		}
	}
	return t
}

func Day6Part2() {
	time := 48938466
	dist := 261119210191063
	var min int
	var max int
	for i := 0; i <= time; i++ {
		if i*(time-i) > dist {
			min = i
			break
		}
	}
	for i := time; i >= 0; i-- {
		if i*(time-i) > dist {
			max = i
			break
		}
	}
	fmt.Println(max - min + 1)
}
