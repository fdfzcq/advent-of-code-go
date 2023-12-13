package year2023

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/fdfzcq/advent-of-code-go/utils"
)

type nodesMap struct {
	instruction []string
	nodes       map[string][]string
	aNodes      []string
}

func Day8Part1() {
	input := utils.ReadStringFromFile(2023, 8)
	nodesMap := parseInputDay8(input)
	steps := 0
	node := "AAA"
	for true {
		i := steps % len(nodesMap.instruction)
		dir := nodesMap.instruction[i]
		nodes := nodesMap.nodes[node]
		if dir == "L" {
			node = nodes[0]
		} else {
			node = nodes[1]
		}
		steps++
		if node == "ZZZ" {
			break
		}
	}
	fmt.Print(steps)
}

func parseInputDay8(input string) nodesMap {
	insStr := strings.Split(input, "\n\n")
	instruction := strings.Split(insStr[0], "")
	nodesStr := strings.Split(insStr[1], "\n")
	nodes := make(map[string][]string)
	var aNodes []string
	for _, nodeStr := range nodesStr {
		re := regexp.MustCompile("[0-9A-Z]+")
		matches := re.FindAllString(nodeStr, -1)
		nodes[matches[0]] = []string{matches[1], matches[2]}
		if strings.HasSuffix(matches[0], "A") {
			aNodes = append(aNodes, matches[0])
		}
	}
	return nodesMap{instruction: instruction, nodes: nodes, aNodes: aNodes}
}

func Day8Part2() {
	input := utils.ReadStringFromFile(2023, 8)
	nodesMap := parseInputDay8(input)
	res := 1
	for _, aNode := range nodesMap.aNodes {
		res = LCM(res, stepsUntilZs(aNode, nodesMap))
	}
	fmt.Print(res)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func stepsUntilZs(aNode string, nodesMap nodesMap) int {
	steps := 0
	node := aNode
	for true {
		i := steps % len(nodesMap.instruction)
		dir := nodesMap.instruction[i]
		nodes := nodesMap.nodes[node]
		if dir == "L" {
			node = nodes[0]
		} else {
			node = nodes[1]
		}
		steps++
		if strings.HasSuffix(node, "Z") {
			break
		}
	}
	return steps
}
