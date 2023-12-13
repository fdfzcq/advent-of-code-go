package year2023

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/fdfzcq/advent-of-code-go/utils"
)

type camelGame struct {
	cards []string
	bid   int
}

func Day7Part1() {
	input := utils.ReadStringFromFile(2023, 7)
	camelGames := parseInputDay7(input)
	sort.SliceStable(camelGames, func(i, j int) bool {
		if gameType1(camelGames[i].cards) < gameType1(camelGames[j].cards) {
			return true
		} else if gameType1(camelGames[i].cards) == gameType1(camelGames[j].cards) {
			for n, _ := range camelGames[i].cards {
				if camelGames[i].cards[n] == camelGames[j].cards[n] {
					continue
				} else {
					return cardValue1(camelGames[i].cards[n]) < cardValue1(camelGames[j].cards[n])
				}
			}
		}
		return false
	})
	res := 0
	for i, cg := range camelGames {
		res = res + (i+1)*cg.bid
	}
	fmt.Println(res)
}

func cardValue1(s string) int {
	switch s {
	case "2":
		return 0
	case "3":
		return 1
	case "4":
		return 2
	case "5":
		return 3
	case "6":
		return 4
	case "7":
		return 5
	case "8":
		return 6
	case "9":
		return 7
	case "T":
		return 8
	case "J":
		return 9
	case "Q":
		return 10
	case "K":
		return 11
	case "A":
		return 12
	}
	return 0
}

func gameType1(cards []string) int {
	m := make(map[string]int)
	max := 0
	for _, c := range cards {
		v, ok := m[c]
		if !ok {
			m[c] = 1
		} else {
			m[c] = v + 1
		}
		max = int(math.Max(float64(m[c]), float64(max)))
	}
	if len(m) == 5 {
		return 0
	}
	if len(m) == 4 {
		return 1
	}
	if len(m) == 3 && max == 2 {
		return 2
	}
	if len(m) == 3 && max == 3 {
		return 3
	}
	if len(m) == 2 && max == 3 {
		return 4
	}
	if len(m) == 2 {
		return 5
	}
	return 6
}

func parseInputDay7(input string) []camelGame {
	lines := strings.Split(input, "\n")
	games := make([]camelGame, len(lines))
	for i, line := range lines {
		cardStr := strings.Split(line, " ")
		cards := strings.Split(cardStr[0], "")
		bid, _ := strconv.Atoi(cardStr[1])
		games[i] = camelGame{cards: cards, bid: bid}
	}
	return games
}

func Day7Part2() {
	input := utils.ReadStringFromFile(2023, 7)
	camelGames := parseInputDay7(input)
	sort.SliceStable(camelGames, func(i, j int) bool {
		if gameType2(camelGames[i].cards) < gameType2(camelGames[j].cards) {
			return true
		} else if gameType2(camelGames[i].cards) == gameType2(camelGames[j].cards) {
			for n, _ := range camelGames[i].cards {
				if camelGames[i].cards[n] == camelGames[j].cards[n] {
					continue
				} else {
					return cardValue2(camelGames[i].cards[n]) < cardValue2(camelGames[j].cards[n])
				}
			}
		}
		return false
	})
	res := 0
	for i, cg := range camelGames {
		res = res + (i+1)*cg.bid
	}
	fmt.Println(res)
}

func cardValue2(s string) int {
	switch s {
	case "2":
		return 0
	case "3":
		return 1
	case "4":
		return 2
	case "5":
		return 3
	case "6":
		return 4
	case "7":
		return 5
	case "8":
		return 6
	case "9":
		return 7
	case "T":
		return 8
	case "J":
		return -1
	case "Q":
		return 10
	case "K":
		return 11
	case "A":
		return 12
	}
	return 0
}

func gameType2(cards []string) int {
	m := make(map[string]int)
	max := 0
	for _, c := range cards {
		v, ok := m[c]
		if !ok {
			m[c] = 1
		} else {
			m[c] = v + 1
		}
		max = int(math.Max(float64(m[c]), float64(max)))
	}
	if len(m) == 5 {
		if m["J"] > 0 {
			return 1
		}
		return 0
	}
	if len(m) == 4 {
		if m["J"] > 0 {
			return 3
		}
		return 1
	}
	if len(m) == 3 && max == 2 {
		if m["J"] == 2 {
			return 5
		} else if m["J"] == 1 {
			return 4
		}
		return 2
	}
	if len(m) == 3 && max == 3 {
		if m["J"] > 0 {
			return 5
		}
		return 3
	}
	if len(m) == 2 && max == 3 {
		if m["J"] > 0 {
			return 6
		}
		return 4
	}
	if len(m) == 2 {
		if m["J"] > 0 {
			return 6
		}
		return 5
	}
	return 6
}
