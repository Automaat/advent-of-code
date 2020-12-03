package main

import (
	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

func findMatching(numbers []int, sum int) (bool, int) {
	supplements := make(map[int]bool)
	for _, line := range numbers {
		if _, ok := supplements[sum-line]; ok {
			return true, line * (sum - line)
		}
		supplements[line] = true
	}
	return false, 0
}

func run(input string) (interface{}, interface{}) {
	part1, part2 := 0, 0

	inputs := pkg.ParseIntList(input, "\n")
	expectedSum := 2020

	if ok, result := findMatching(inputs, expectedSum); ok {
		part1 = result
	}

	for _, line := range inputs {
		if ok, result := findMatching(inputs, expectedSum-line); ok {
			part2 = line * result
			break
		}
	}

	return part1, part2
}

func main() {
	execute.Run(run, nil, puzzle, true)
}
