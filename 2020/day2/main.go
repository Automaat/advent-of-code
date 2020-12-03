package main

import (
	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"strings"
)

type Password struct {
	minRange     int
	maxRange     int
	policyLetter string
	password     string
}

func parsePassword(line string) Password {
	parts := strings.Split(line, ":")
	policy := strings.Split(parts[0], " ")
	passRange := strings.Split(policy[0], "-")
	return Password{
		minRange:     pkg.MustAtoi(passRange[0]),
		maxRange:     pkg.MustAtoi(passRange[1]),
		policyLetter: policy[1],
		password:     strings.TrimSpace(parts[1]),
	}
}

func parseInputs(inputs string) []Password {
	lines := strings.Split(inputs, "\n")
	passwords := make([]Password, len(lines))
	for i, line := range lines {
		passwords[i] = parsePassword(line)
	}
	return passwords
}

func isCorrect(password Password) bool {
	passLetters := strings.Split(password.password, "")
	occurrences := 0
	for _, letter := range passLetters {
		if letter == password.policyLetter {
			occurrences++
		}
	}
	if occurrences < password.minRange || occurrences > password.maxRange {
		return false
	}
	return true
}

func isCorrectWithTobogganPolicy(password Password) bool {
	passLetters := strings.Split(password.password, "")
	if passLetters[password.minRange-1] == password.policyLetter && passLetters[password.maxRange-1] == password.policyLetter {
		return false
	}
	if passLetters[password.minRange-1] != password.policyLetter && passLetters[password.maxRange-1] != password.policyLetter {
		return false
	}

	return true
}

func run(input string) (interface{}, interface{}) {
	part1, part2 := 0, 0

	passwords := parseInputs(input)

	correctPasswords := 0
	for _, pass := range passwords {
		if isCorrect(pass) {
			correctPasswords++
		}
	}
	part1 = correctPasswords

	correctPasswordsWithTobogganPolicy := 0
	for _, pass := range passwords {
		if isCorrectWithTobogganPolicy(pass) {
			correctPasswordsWithTobogganPolicy++
		}
	}
	part2 = correctPasswordsWithTobogganPolicy

	return part1, part2
}

func main() {
	execute.Run(run, nil, puzzle, true)
}
