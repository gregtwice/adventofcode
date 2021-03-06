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
	}, {
		puzzle,
		`240`,
		`2832009600`,
	},
}

var puzzle = `.........#....#.###.........##.
..###.#......#......#.......##.
##....#.#.......#.....#........
#........#..........#.#...#...#
#....###...##.....#........#...
###..............##..#.....#...
.................##.#..........
.........##......#..###.....#.#
..#..#...#.#.#...#.#.#.##...#..
..............#.#.#..#..#..#...
.#.#.#....#.........#.........#
..#.#....##..#...#.....#..##..#
............#.....#.........##.
...#.#..#..........#.##.....#..
#......#...##.......###........
.....#....#.#..............#...
.....#.......#..##.###.....#.#.
.#.....#....##.#......##......#
..##...##.........#..#.#..#....
............#.......#.....#....
.......................#...####
.#.#.......#..#....#....##..#.#
..#.##.#.#...............#.....
#..##..#...#.....##..#...##.#..
##...#....#####.....##...#...##
.#..##..#..#.#..##.#.#........#
....#..#........##......#.#....
..#......#...##.#..##.......#..
.#.#....#.#..#.....#..#...#....
.....#.#.................#.....
##.#........#.....#...........#
#............#.....#..#.#...##.
..#.#..#......#.......#......##
....#.#....#...#....#..........
.........#...#.##..#...#...#...
....#...#...#..................
..##...#.#....#...#......#.....
#....#.......##..#...#..#......
.#....##..#.#....#...##...#....
#.................#...#.#...#..
.#.....#..........#.......#....
.#..........#.##....#.##......#
#.#.....##.##..#.......#..#....
.....#...#............#..##....
...#.#.##.#..........#.#....#..
.......#...#............#.....#
..........#...##..#....###....#
............#....#......###....
...................####..#.##..
...#.#.##.........#..#.#.......
...........#.........#..###....
.........#.........#...#...###.
.#.#.##....#.#...........#.#...
..###.....#...#.#.......##..#..
.....#.#........#.#....#....##.
......#...#........#.........#.
...............#.........#.#...
..#...#...#...#.###..###....#..
#..#.......#..###.##.......#...
#.#.........##..#.....###..#..#
...#....#....#.#..#............
..###..##....................#.
..#.......#..........#.##......
..##........#...###..##.#......
......#.#...#.....#..##..#.....
#........#......#..............
........#........#.......#....#
.....#.......#......#........##
#.#......#.#...##.#.......#....
#...................##...#.....
..#.#...#..#...#..#.....#..#...
.......##..................##..
.............#..##.#......##..#
###...........##.#....###..#...
.#..........##...#..#......#...
..#.###.#....#........#........
....#....###.....#.......#.....
.....##....#..##...#..#........
.##...#..#....#..#.........#...
#.........##....#..##..##......
.#.#.............#.....##......
..#.#............#.......##.#..
..#.##..#.........#......#.#.#.
.#...#...#..#....#...#....#.#.#
....#...#..#.##..##.......##...
.#.....#.....#............#....
..........#....#..#..#......#..
.............#....#..#.........
....#.#.#.......#....#.#.......
..##....#.#....#...#........###
#...#..........#..........###..
...#..#...#...#..##......#...#.
.....................#...#..#.#
#..#............#.........#.##.
..#...#...........#.......#....
.....##..........#...##.....##.
.#.....#.#........##...........
..#....#..#.#..##...#.........#
.........#.###.##....#..###....
.........#...##...##.#.#....##.
...#..##.#...........#....#....
..........#.#...........##.....
...........#..........#...##...
.........#..........#...#.#...#
......#..#.................#...
.....#...........#..#...#..###.
.....#..#....#.#.##...##....##.
...##.###.#.#..#.#....##.#.#...
###.....#.....#........#...#...
.#....#........#.#....#..#...##
##.....#.....#......#.#..#...##
.##....#...............#.......
#.#.....#.#....#.........#.....
..#...............#.......##...
#...#.###..#....#.#.#..#.#.....
##.###....##...#....#.....#....
.......#................##.....
....#....##..#....#..#......#..
...#.........#...#........##.##
.#..............#..............
..##.......#.###..##.#.........
..#...#...#...#...#...#.....##.
.....#..##...#.....##..#.#.....
..#.............#...#.........#
#.........................#..##
.............#..###........#...
......#..#....#.##.......#..#..
...#..#..#...#....#..#...##...#
.##............#.......###.#...
.#........#..#.................
#.#.#.....##....##...#.....#.#.
...##.......#.#......#...#.#...
#.##....#.........#.....##...#.
#...#..#....#.......#.##...##..
.................##.#..##.#.#.#
..#.............#.......#.#.##.
#....##..................#...##
..........#.......#..##........
......#.#..#......#.#.........#
#.#........#####......#.#.#....
#..#........#.#..#.....#.......
...#.............#.............
.....#.......#......#..#.##..#.
..#.........#..........#.##...#
#.....#.#####......#.......###.
.......#.....#...#.....#.#..#..
#...#.#........#.#..#..#...#..#
...#....#....##.....#..........
.#.......##.......##...........
...#.##.#.#..#....#...##....##.
.#...#...#.........#...........
.#.#.##..#.......#.#...#..#....
.#....##.#.#...#......#......#.
##..#..#..#..#.......#......#..
.........#.#...........#....#.#
........#....#.#...#.#..#......
#.......#.#.................##.
.....#..#..#....#.#........#.##
.#..###..#....#..#........#.#..
#...........#...#........#.....
........#..#.#.#.#.......#....#
....#..#..........#.#..#.....#.
..####..........#..............
....###.#..#........##..#......
.#..#......##..........#...#...
.#.....#....#......##.##...#.##
..##.#.#......#.......##.......
....##.......#..............#..
........#.....##..............#
.#...#....##.....#....#.......#
....#.......###.......#.#.#....
##.....##........#.....##......
..........#.....#...##.#..#.#..
..........#.#......#..........#
..........#...#..#...#...#.....
.#.......#..##.................
.#........##..............#..#.
.##...................#...#....
.##....#.##.##........#........
...##.....####.....#..#.......#
...##.#...##...#.##............
##.......#.....###..#..#...##..
#.####...#...#...##..#..#....#.
...#........##........#........
#....#.#....##..#..#.##...#....
...##....##....#.......#..###..
..........#..#..........#..#.#.
#..#....#.......#.......#....#.
......#......#.....##..##.#..#.
##.#.....#....#.......#...#...#
..##..#.#...#...#.....###..##..
....#..#.......#............##.
#..##.#.#.....#####....#....#..
.#........#...#.#..##.#.####.#.
#...#...#.............#.#......
.........#.....##..........#...
.##....#....#........#......#.#
#..###...#....#..........#.....
.#...##.........#..#..#.#...#..
#.#.#.......................#..
#.....#..#.#............###....
#...#.....#.....###..#..#.#.##.
............#.........#.#.##...
...#.......###......##......#..
.#....#.#....##......##.#...#..
.........#.......#....#...##..#
................#.#.....#....#.
.##......#....#..#..###..#...#.
....##....#..#....#.##..#......
.......#.#.....#..#............
..........#....#....#..#..#....
..#....#.....#.......#.........
......#.........#.##..#....##.#
..#..#.#.......##..#...##......
...##..#.#.#...............##..
..#.#.#......#....#.....#.##...
..#.....#.#...........#....##..
###.....##.....................
.......#..#.................#..
.#.#..#..#.........#......#...#
##.......#.##.......#..........
#..#.....#.....#.....#.......#.
#..#.....#.....#..........#.#..
.#....##....#.....#.......#.##.
.....#.#......##..##.#.........
#....##......#..#....#..##..#..
#.##..#..#..............#...#..
.#......#......#..#...........#
..#.......#........#....#..#...
.....##.#.......##....#.#....#.
........#....................#.
........#..#..........#........
......#.#.....#.#.....#......#.
#......##......###.##......#...
...........###..#...#.......#..
..#...###...#.....#....#...#..#
.....##......#.#......##..#.#..
#.#......##...#.....##...#...#.
.#.#........#.......#.........#
....#....#...##..........#.#...
.#..##.#...#.#.....##......##..
.#.....#..##....#....#....#..#.
..#.......#.#.#..........#..#..
#.#..#....#.##....#.......#....
........##.........#..#.#......
.......#.....#.##..##......#.#.
.........................#.#.#.
..#..............##.........#..
.......###.#.#.......#.........
#........#.....#.......#..#...#
##....#..#....#...........#....
..#..#.#.#.....#.#..#....#.....
.#..##....#.##..#..............
...#....##..#..#.##....#.#.....
...##....#......##..#........#.
....#......#....#....#........#
...#..#...#.#...........#..#..#
....#.#.#.........#...#...#....
..........#....#......##....#.#
..##..#...##.#...###.#.##......
#........#.##......##.#........
..#...#.##...#..........#.#...#
...........#...........#.......
......###....#..###..##........
...#........#..#.#.............
....#.#.....#.#............#..#
##.#.....#........#....##.....#
.......#.#...#..##.......#.#.#.
#......##..#..#...#.....#..#.#.
.#......###.....#..#.....#...#.
....#.#.............#.##.......
......#....#.....#.......#..#..
#..#.#.#..#......#...#..#.#....
#..............#.#....#...#....
..#......#...##.#......#..#..#.
.......##..#.##..#.#...#.....##
.....#...........#....#.......#
.#.........#..#..........##....
#..##..#.#......##.......#..#..
...#....#...........##.#.#.#..#
#..#..#..#...........#....#.#.#
.....##......#......#.#...#....
.....#..##....###.....#....##..
........#...##......#.....#..#.
..#.#..#.#....#...#.......#....
#.....#...#.....#.#.......#....
......#...#.......#..##....#...
#............#.....#....#......
..##...#.....#..#......#...#...
...#..#...#..#.......#........#
...##.#.#........#........#....
#.#..#......##.#..#..#......#..
#.......#..#...................
#.....#....#......#...........#
.##.#...#.#...#..............##
...###........#........#..##.#.
..##....#.#.#.##..#.#......#...
..#.#........#..............#..
.......#.................##....
..................#............
....#.....#.#..............#.##
......#.....#..#......#...#....
..#....##......#...####....#.#.
..........#...##...........##.#
...#.#..#....#......#..#....#..
#.........###...#.....#..#....#
.#.#......##.#.....#...........
...............#.#....##..#..#.
..........#.#.#.#...#....##...#
...#....##.................##..
#..##..#...##.##.#......#.#.###
#..#...#..#.....#...#.#..##...#
..#................#...##....#.
...#.....#.##.......##....#.#..
....#.....#..#....##...........
...............#..........#....
....#...#........##...#........
...#....#...#.###..............
#.#....#.......#..#.##.##......
.#.....#..#..#......#....#.#...
...#..........................#
............#.#..#.##......#...
.....#..........#.#........#.#.`
