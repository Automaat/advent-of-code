package main

import (
	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"strings"
)

func part1Solution(numbers []int, moves int) int {
	lastSpokenTurn := make(map[int]int)
	last := numbers[0]

	for i, num := range numbers[1:] {
		lastSpokenTurn[last] = i + 1
		last = num
	}

	for i := len(numbers) - 1; i < moves; i++ {
		if val, ok := lastSpokenTurn[last]; ok {
			lastSpokenTurn[last] = i + 1
			last = (i + 1) - val
		} else {
			lastSpokenTurn[last] = i + 1
			last = 0
		}
	}

	return last
}

func run(input string) (interface{}, interface{}) {
	part1, part2 := 0, 0

	numString := strings.Split(input, ",")

	numbers := make([]int, len(numString))

	for i, num := range numString {
		numbers[i] = pkg.MustAtoi(num)
	}

	part1 = part1Solution(numbers, 2019)
	part2 = part1Solution(numbers, 30000000-1)

	return part1, part2
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
