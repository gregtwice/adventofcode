package main

import "github.com/kindermoumoute/adventofcode/pkg"

var tests = pkg.TestCases{
	{
		`#ip 0
seti 5 0 1
seti 6 0 2
addi 0 1 0
addr 1 2 3
setr 1 0 0
seti 8 0 4
seti 9 0 5`,
		`7`,
		``,
	},
	{
		puzzle,
		`1392`,
		`15826992`,
	},
}

var puzzle = `#ip 3
addi 3 16 3
seti 1 3 2
seti 1 0 5
mulr 2 5 1
eqrr 1 4 1
addr 1 3 3
addi 3 1 3
addr 2 0 0
addi 5 1 5
gtrr 5 4 1
addr 3 1 3
seti 2 2 3
addi 2 1 2
gtrr 2 4 1
addr 1 3 3
seti 1 1 3
mulr 3 3 3
addi 4 2 4
mulr 4 4 4
mulr 3 4 4
muli 4 11 4
addi 1 4 1
mulr 1 3 1
addi 1 2 1
addr 4 1 4
addr 3 0 3
seti 0 2 3
setr 3 6 1
mulr 1 3 1
addr 3 1 1
mulr 3 1 1
muli 1 14 1
mulr 1 3 1
addr 4 1 4
seti 0 6 0
seti 0 9 3`