package main

import (
	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"regexp"
	"strings"
)

func parseInputs(inputs string) []map[string]string {
	passports := strings.Split(inputs, "\n\n")

	parsedPassports := make([]map[string]string, len(passports))
	for i, passport := range passports {
		flaPassportData := strings.ReplaceAll(passport, "\n", " ")
		parsedPassports[i] = make(map[string]string)
		for _, flatPassportEntry := range strings.Split(flaPassportData, " ") {
			entry := strings.Split(flatPassportEntry, ":")
			parsedPassports[i][entry[0]] = entry[1]
		}
	}

	return parsedPassports
}

func validPassportStructure(passport map[string]string) bool {
	if len(passport) == 8 {
		return true
	}
	if _, ok := passport["cid"]; !ok && len(passport) == 7 {
		return true
	}
	return false
}

func validByr(passport map[string]string) bool {
	byr := passport["byr"]
	byrInt := pkg.MustAtoi(byr)
	return byrInt >= 1920 && byrInt <= 2002
}

func validIyr(passport map[string]string) bool {
	iyr := passport["iyr"]
	iyrInt := pkg.MustAtoi(iyr)
	return iyrInt >= 2010 && iyrInt <= 2020
}

func validEyr(passport map[string]string) bool {
	eyr := passport["eyr"]
	eyrInt := pkg.MustAtoi(eyr)
	return eyrInt >= 2020 && eyrInt <= 2030
}

func validHgt(passport map[string]string) bool {
	hgt := passport["hgt"]
	if strings.HasSuffix(hgt, "in") {
		number := pkg.MustAtoi(hgt[0 : len(hgt)-2])
		if number >= 59 && number <= 76 {
			return true
		}
	} else if strings.HasSuffix(hgt, "cm") {
		number := pkg.MustAtoi(hgt[0 : len(hgt)-2])
		if number >= 150 && number <= 193 {
			return true
		}
	}
	return false
}

func validHcl(passport map[string]string) bool {
	hcl := passport["hcl"]
	hclParts := strings.Split(hcl, "#")
	if len(hclParts) != 2 {
		return false
	} else {
		if len(hclParts[1]) != 6 {
			return false
		}
		match, _ := regexp.MatchString("[0-9]", hclParts[1])
		match2, _ := regexp.MatchString("[a-f]", hclParts[1])
		return match || match2
	}
}

func validEcl(passport map[string]string) bool {
	ecls := map[string]bool{"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true}
	ecl := passport["ecl"]
	if _, ok := ecls[ecl]; ok {
		return true
	}
	return false
}

func validPid(passport map[string]string) bool {
	pid := passport["pid"]
	return len(pid) == 9
}

func validPassportValues(passport map[string]string, validators []func(map[string]string) bool) bool {
	valid := true
	for _, isValid := range validators {
		if !isValid(passport) {
			valid = false
		}
	}
	return valid
}

func run(input string) (interface{}, interface{}) {
	part1, part2 := 0, 0
	validators := []func(map[string]string) bool{validByr, validIyr, validEyr, validHgt, validHcl, validEcl, validPid}

	passports := parseInputs(input)

	validPassports := 0
	for _, passport := range passports {
		if validPassportStructure(passport) {
			if validPassportValues(passport, validators) {
				validPassports++
			}
		}
	}
	part1 = validPassports
	return part1, part2
}

func main() {
	execute.Run(run, nil, puzzle, true)
}
