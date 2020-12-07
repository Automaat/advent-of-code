package main

import (
	"fmt"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"strings"
)

func parseInputs(inputs string) []string {
	return strings.Split(inputs, "\n")
}

func run(input string) (interface{}, interface{}) {
	part1, part2 := 0, 0

	lines := parseInputs(input)

	highestId := 0
	seats := make([]bool, 878)
	fmt.Println(len(lines))

	for _, line := range lines {
		rowMin := 0
		rowMax := 127
		seatMin := 0
		seatMax := 7
		for _, letter := range strings.Split(line, "") {
			if letter == "F" {
				rowMax -= ((rowMax - rowMin) / 2) + 1
			}
			if letter == "B" {
				rowMin += ((rowMax - rowMin) / 2) + 1
			}
			if letter == "R" {
				seatMin += ((seatMax - seatMin) / 2) + 1
			}
			if letter == "L" {
				seatMax -= ((seatMax - seatMin) / 2) + 1
			}
		}
		id := rowMin*8 + seatMin
		seats[id-1] = true
		if id > highestId {
			highestId = id
		}
	}

	for i, seat := range seats {
		if !seat {
			if i > 1 && (seats[i-1] && seats[i+1]) {
				part2 = i + 1
			}
		}
	}

	part1 = highestId

	return part1, part2
}

func main() {
	execute.Run(run, nil, puzzle, true)
}
