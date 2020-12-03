package main

import (
	"fmt"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"strings"
)

func parseInputs(inputs string) [][]bool {
	lines := strings.Split(inputs, "\n")
	ySize := len(lines)
	xSize := len(lines[0])

	treeMap := make([][]bool, ySize)
	for i := range treeMap {
		treeMap[i] = make([]bool, xSize)
	}

	for i, line := range lines {
		letters := strings.Split(line, "")
		for j, letter := range letters {
			if letter == "." {
				treeMap[i][j] = true
			}
			if letter == "#" {
				treeMap[i][j] = false
			}
		}
	}

	return treeMap
}

func traverseForrest(treeMap [][]bool, xJump int, yJump int) int {
	treesEncountered := 0
	xSize := len(treeMap[0])

	j := xJump
	for i := yJump; i < len(treeMap); i += yJump {
		if !treeMap[i][j] {
			treesEncountered++
		}
		j += xJump
		if j >= xSize {
			j = j % xSize
		}
	}

	return treesEncountered
}

func run(input string) (interface{}, interface{}) {
	part1, part2 := 0, 0

	treeMap := parseInputs(input)

	part1 = traverseForrest(treeMap, 3, 1)

	aTrees := traverseForrest(treeMap, 1, 1)
	cTrees := traverseForrest(treeMap, 5, 1)
	dTrees := traverseForrest(treeMap, 7, 1)
	eTrees := traverseForrest(treeMap, 1, 2)
	fmt.Printf("1-1: %d, 3-1: %d, 5-1: %d, 7-1: %d, 1-2: %d \n", aTrees, part1, cTrees, dTrees, eTrees)

	part2 = aTrees * part1 * cTrees * dTrees * eTrees

	return part1, part2
}

func main() {
	execute.Run(run, nil, puzzle, true)
}
