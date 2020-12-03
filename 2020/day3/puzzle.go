package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

var tests = execute.TestCases{
	{
		`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`,
		`7`,
		`336`,
	},
}

//var puzzle = `..##.......
//#...#...#..
//.#....#..#.
//..#.#...#.#
//.#...##..#.
//..#.##.....
//.#.#.#....#
//.#........#
//#.##...#...
//#...##....#
//.#..#...#.#`

var puzzle = `.....#............#....#####.##
.#.#....#......#....##.........
......#.#.#.....###.#.#........
......#...#.....#####....#..##.
...#............##...###.##....
#.....#...#....#......##....##.
#...#.#....#..#..##.##...#.....
.......#..........#..#..#.#....
.#.....#.#.......#..#...#....#.
#..#.##.#..................###.
...#.#.##...##.###.....#..#...#
..#.#...#............#.......#.
#..#.#..#.#....#...#.#.....#..#
#......##....#..#.#.#........#.
....#..#.#.#.##............#..#
....#..#..#...#.#.##......#...#
##...#...........#.....###.#...
..#...#.#...#.#.....#....##.##.
....##...##.#....#.....#.##....
#........##......#......#.#.#.#
....#.#.#.........##......#....
.#......#...#.....##..#....#..#
....#..#.#.....#..........#..#.
..##...#..##................#.#
.....#....#.#..#......#........
........#..#.#......#.#........
.....#.#....##.###....#...#....
...##.#.......#....###..#......
............##.#..#...#........
#..###..#.....#.####...........
.......##.....#......#......#..
#........##..#.....##.......#.#
#.##...#...#...#......##..#.#.#
......#....##.#.#...#...##....#
#..#....##.#......#.......##...
.#..........#..........#....#.#
#.....##......##....#..........
..#.#.....#.#...#........#.....
...#........#..#..#.##..##.....
......###.....#..#...#.###...##
.##.##.......#.......###...#...
#.#..#.#.#....#.....###..#...##
......#.##..........#.......##.
#..#.#.........#.....##...##...
..#...#....#....###.#......#...
.....#..#.######.....#..#.#....
..#.#.....#.....##.#....##.#.##
...#.#.#....#....##..#..#.#.##.
...........#.#...#..#..####....
.........#####.#.#.#...#.##.#..
.......#...#......#.##.#.##....
....#.....#.....###..........#.
.#.###....##.#..#..........#...
#...#.........##.....####....#.
##....##...#..........#........
...#.#.#.#....#..........#.....
.......#....#......##.......#..
.#.#..#.........#.#.##....#....
..#.............#..##.#.##..###
.#.##..............#..#..##..#.
..##.#..#......#...##..#.##...#
......#..#....#....#....##..#..
...#...##.............#..###...
...##....#.#.##........#.....##
....#.#.......#..###..#....####
#...#...##..#.####..#...##....#
.......#..#.##..#...#.#........
###.#......#..##..#......#.##..
#....#............#.....#......
..##...#..##......#..#....#....
.#..##...#....#.#...#...#..#..#
........#....###...#..##..###.#
.........#....#....#..#.#.#...#
.#....###.##...#.#...........##
..#..#.#..#.#.##..#...##.......
##..#.#.#....#...#..#..........
#..#.......#....#..##...####...
............#.#..........##.##.
#...#..#.#....#..#.#....##.....
......#...#...#.##............#
#.....##..###..#.#..#.#.##..#.#
#..#.#..#......#.......##.#....
##..#.#..#...#......#.##...###.
.#....#..............#....#.#..
..#.#..##....#....#..##........
.#.#...#..#.....#.#..##........
.....#..#.#......#....#.#..#.#.
....#.###...###.#.#.....#......
...........#.#....##....##.....
..#..#.##..........#...#...#..#
.....#.###.#..........#........
....#....##........###...#.....
.#.....##.......#....#..##..###
#.....#...............##......#
#..#.#..#.#.#.....#.#...#......
.##.###...#....#..........##...
.#.......#.....................
.#.#....#...##..#...#...##.....
.#.#...#.......#.......#...#...
....#.#..#.#..#...#....##......
....##.....#.##....#.##..##..##
..#............#...###.##..#...
.#..#.........#.##....#....#..#
.#..##..#..#........#.#.##.#.##
.###.#...#...............#...#.
...#.##.##.#......#...#....##.#
#......##.......##...###....#.#
#..##.....##......#.#.##....#.#
...#.#....#.#.#...........##..#
#.....##......##.#..........##.
###....#.#...#.#..####.........
.##.#.#...##..#.....#..#...#...
#.....#.#......#..........#.#..
..###.##.#...................#.
#.............#..#........#.##.
#.#.#.#..#.....##..##.#....#...
...#...#..#...#..##..##........
...##...##..#...##...........#.
.####..#.#.#.##.#.......#......
...#....#.......#......#.......
.....#.#...#...#..##..#..#.....
......#.....###.#..#..#.#..###.
.#....#....#..#..##.....##...#.
.#.............##.###.#...#.#..
#..#..#......#.###............#
##.#..##....#..........#.#.#...
......#........#...#.......##..
....#.#..#..........#.....#.#..
...#..#...#.#...#........#.....
.....##...#....#.........##.##.
....#...#...#.##.##...#....#...
.#..#.....##......#..#.#..#....
........##...##.##......#.#.#.#
.................#..#.....##.#.
...#.....#...#.........#..#.#.#
....##.#.....#........#...#..#.
#...............#..#.....#...#.
.....#..#....#...#.####.#.#....
####.#..#.##...#....#...##.....
#...##..#...####..#....#.#...#.
..#.......#.##..##...#.#.......
...........##.......#....#..#..
#.##....#...#.....#....##......
....##.#.......#..#...##.......
...#.........##.#..#......#.###
.#..#..#....#.#.##....###..###.
....#.#........##........##....
....########....#.#.#.###.#...#
...#.###.###.##......##.......#
.#...#.###.......#..........#..
..#..##.........#............#.
.......##.#...#...##...#...#..#
#.##....#.#...#.....#..#.#.....
..#........#..#.#.#.#....#.##..
...#...#.#.........#...#.#..##.
#....#......#.#...........#..##
...#.#.#..#...##...#...#...#...
###..........#.#..........#....
..#....#.#.#.#............#.#..
....#...#..###...#.#....#......
#...........####......##.#.....
..#..##.#...#.....#..#.......##
#.....#..###.....#...##..##....
##..###..##...........#.#...#..
.....#......#..............#...
#..#.##.###.......#.......#...#
#........#....##......#.#......
.#.#.#...#.......#........#.##.
#..#..##.....#...#.#.#.#..###..
.#.#....#..#..#.#....##.#.#....
..#.#.........####.#...#.#.###.
....##........##....#........#.
................#..........#...
..#...................###.##..#
.........#..#..#.#...#....#.#.#
......#.....###.....#.#..#...#.
.#.#.....#..##............##...
...##......##.#....#...........
...##..##..###.#...##..........
....###...#..#.#......#......#.
....##..............#..#..#.#..
####.......#...##.##..#.#......
.#......#.....#....###..#....#.
.#.......#...##...#..##.#......
#.......#.......#.#....#.#.#..#
........#..#..#............##.#
#.#...#.#..##..#.......##..#...
...#....#...#..........##..#...
#.#...#.##....###......##....#.
#..#...###........#..#....#..#.
#....#....###....#..#.......#..
....#.#........#.............#.
.#.##........##...#...#...#...#
#.....##.....#.......#.#.#.....
.#.##..........##..#....#......
.#..##.##.#...##....#.#....##..
........#.#.##.#....#.#..#....#
..#...........................#
.#...........#....#....#.#..#..
........##...........#...#...#.
..#.....#..#......#..##.......#
..#....###..###...#.#.#..#....#
#..#.#...#......##......#......
...........#...##..##....##....
#.#......###..#.....#.......#.#
#.....#....#....#.#...#...#....
....#...#.......#....##.#..#...
.####..##......##.#........#..#
..###..#.#.....#...........##..
..##.#.#..#....#..#..#.........
..........#.#.#####...#........
.###......##.#....#........#...
.....#..#..#.#..#.........#....
..#....#...#...#...##..........
....#..##.#.........##.#..##...
##.####..#...#.#...#.....#..###
..#..#...#...#.....##....#..#.#
#..##..#.....#....#.#.....##..#
...#...........##.....#......#.
......#...#.....#.#..###.......
.........#.....###.##..#...#...
.#...#.##...#..........#.#..##.
......#.......##.....#.....##..
........###..........#...#.....
##.......###..###...##...#.....
#.#.............#..#..#.#......
..##........#.###.....#....##..
......#...#......#....##......#
..#.....#...##...#.......#..#..
..#.###..##.##...#....#...##.#.
........##...#..#.#..##.....#.#
.......................#......#
..##.###......#.#.............#
....#...........###............
##...##.....#.......##.......#.
...#..##..##..#.#.###..#......#
........#........#.#..#..#.....
.#......#....##..........#...#.
.##...........##....#..........
.#..#....###.......#....#..##..
.....###..........#....#.#.#...
...#....###.#.#......#......#..
#.#.##.#.....#..#........#...#.
...#.##.........#..#.....#.....
.##...##......##...###...#.....
...#.....#.##..#...#..#........
........#............#.#.#..##.
###...#.....#...#..#........##.
##...#..#.....#.#....#.#.#.....
#..##.......#...#.#...##..#....
#...#.##.....#.#..#.##......#.#
..#......#.#.#.##.##..........#
..#.##......#.#.#..##..........
....#..#....#..#..............#
..........###.....##..#........
...#.....##.....#..#.#..#...##.
.#..##.#..#....#.#......#.##...
...#.....#..#.#...#..#.....#.#.
#...#.#......##...#..#...#....#
..#.......##...#..#.......#...#
#.....#...........##.#.........
.#......##.....####...#.......#
........#..#.....#.......#..#..
....#.#...##..##...#..#....#...
#.#......#...#.#.###.....#.....
..##...#.#........#.##....#.#.#
.#....#......#.#...###.#.......
.......#.#...##....#.#....#....
.....##..##...#..#.#.....##..#.
.##..#.#.#....##.#...#.....#...
.#..#..##....#.##.......#...#..
....#.##...#..##......#.....#..
.#..#....##....#...............
..##...#.....###...............
..............#.#.##........#.#
.#.#....#....#...#.#........#..
.##...#...#.#....#....#.#.....#
#..............#......#.####.#.
......#...........#..#.....##..
#.#..##.##.....#......#..#.#..#
##.##..#.##.#.............#...#
...#..#......#....#............
........###.#.#..#...#.....#.##
..#.......#.##.........#..#....
...##.#........##...#.#.##..#..
...#..#......#...#....#........
...........#..#..#...##...#....
...#.....#....#..####..##.....#
.......#..#..#......#.........#
#......#........###.....##....#
..#..#..#.#.#....##...##......#
#.#..#..###.#..#.....####......
.#................#####....#...
.#.........#...#.......#......#
..#.......#######........#.....
..#........#.....#..#...#..#..#
.#..#.#..#....#.#..##...#..#.#.
..#...........#.#...#.#.##.....
...#.#.#....##.###....#...####.
.....#..#.....#..#.#.........#.
......##...#...###............#
..#.#......###..####..#......#.
###.##.#..#......##.#..##.....#
....###...##............#.#....
..#.....##...#...##....#...#...
#.....#.....#...#...#.#..#.....
####..........##.#.#..#.....##.
...#..........#...#...##..##.#.
..........#.........#.#..#..#..
#....###.....#.#...#.......##.#
#..#.##.....#..........#...#...
...#.#.###.......##..#.....#...
#...#.#..#.............#..#.#..
#........#.................#..#
..#.#....#.#..##.#...#..#....#.
#...#..........#...###....#...#
......#.............#....#....#
#.#.......##.......#.#....##..#
##...#....#.............#..#...
........#...###.##.#..###.#...#
...##...#..#..#...##..##......#
..#.......##....#.#.##....#....
.....#....#..#.#...##.#.#.....#`
