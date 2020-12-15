package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

var tests = execute.TestCases{
	{
		`0,3,6`,
		`436`,
		`175594`,
	},
	{
		`1,3,2`,
		`1`,
		`2578`,
	},
	{
		`2,1,3`,
		`10`,
		`3544142`,
	},
}

var puzzle = `12,1,16,3,11,0`
