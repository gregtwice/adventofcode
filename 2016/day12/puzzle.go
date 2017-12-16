package main

type testCase struct {
	input, expectedPart1, expectedPart2 string
}

var tests = []testCase{
	testCase{
		`cpy 41 a
inc a
inc a
dec a
jnz a 2
dec a`,
		`42`,
		``,
	},
}

var puzzle = `cpy 1 a
cpy 1 b
cpy 26 d
jnz c 2
jnz 1 5
cpy 7 c
inc d
dec c
jnz c -2
cpy a c
inc a
dec b
jnz b -2
cpy c b
dec d
jnz d -6
cpy 14 c
cpy 14 d
inc a
dec d
jnz d -2
dec c
jnz c -5`
