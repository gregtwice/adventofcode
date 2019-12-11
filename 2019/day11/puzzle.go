package main

import "github.com/kindermoumoute/adventofcode/pkg"

var tests = pkg.TestCases{
	{
		puzzle,
		`2129`,
		`PECKRGZL`,
	},
}

var puzzle = `3,8,1005,8,337,1106,0,11,0,0,0,104,1,104,0,3,8,1002,8,-1,10,101,1,10,10,4,10,108,1,8,10,4,10,1002,8,1,28,2,1,15,10,2,2,10,10,1,1107,0,10,2,1105,18,10,3,8,102,-1,8,10,1001,10,1,10,4,10,1008,8,1,10,4,10,101,0,8,67,1,1003,4,10,2,1007,14,10,1006,0,64,3,8,102,-1,8,10,1001,10,1,10,4,10,1008,8,0,10,4,10,102,1,8,100,2,102,15,10,3,8,1002,8,-1,10,1001,10,1,10,4,10,108,0,8,10,4,10,1001,8,0,125,2,1003,7,10,1006,0,10,2,1007,13,10,2,103,14,10,3,8,102,-1,8,10,1001,10,1,10,4,10,1008,8,1,10,4,10,101,0,8,163,1006,0,5,3,8,1002,8,-1,10,1001,10,1,10,4,10,1008,8,0,10,4,10,102,1,8,188,1,1101,2,10,1006,0,82,3,8,1002,8,-1,10,101,1,10,10,4,10,1008,8,0,10,4,10,101,0,8,217,1,1109,1,10,1,109,9,10,1,1009,9,10,1006,0,41,3,8,102,-1,8,10,1001,10,1,10,4,10,1008,8,1,10,4,10,102,1,8,254,2,104,1,10,2,8,15,10,3,8,1002,8,-1,10,1001,10,1,10,4,10,1008,8,1,10,4,10,101,0,8,284,1,1107,11,10,3,8,102,-1,8,10,101,1,10,10,4,10,108,0,8,10,4,10,101,0,8,309,2,1001,10,10,1006,0,49,101,1,9,9,1007,9,1058,10,1005,10,15,99,109,659,104,0,104,1,21101,937267929896,0,1,21101,0,354,0,1106,0,458,21102,1,936995566336,1,21102,1,365,0,1106,0,458,3,10,104,0,104,1,3,10,104,0,104,0,3,10,104,0,104,1,3,10,104,0,104,1,3,10,104,0,104,0,3,10,104,0,104,1,21101,3263269979,0,1,21101,0,412,0,1106,0,458,21102,1,46174071899,1,21101,0,423,0,1106,0,458,3,10,104,0,104,0,3,10,104,0,104,0,21101,825544561428,0,1,21102,446,1,0,1105,1,458,21102,1,867966018404,1,21101,457,0,0,1106,0,458,99,109,2,21202,-1,1,1,21102,40,1,2,21101,489,0,3,21102,1,479,0,1105,1,522,109,-2,2106,0,0,0,1,0,0,1,109,2,3,10,204,-1,1001,484,485,500,4,0,1001,484,1,484,108,4,484,10,1006,10,516,1102,0,1,484,109,-2,2105,1,0,0,109,4,2102,1,-1,521,1207,-3,0,10,1006,10,539,21101,0,0,-3,21201,-3,0,1,22101,0,-2,2,21101,1,0,3,21102,558,1,0,1105,1,563,109,-4,2105,1,0,109,5,1207,-3,1,10,1006,10,586,2207,-4,-2,10,1006,10,586,22102,1,-4,-4,1106,0,654,22101,0,-4,1,21201,-3,-1,2,21202,-2,2,3,21102,1,605,0,1105,1,563,22101,0,1,-4,21102,1,1,-1,2207,-4,-2,10,1006,10,624,21102,1,0,-1,22202,-2,-1,-2,2107,0,-3,10,1006,10,646,21201,-1,0,1,21102,646,1,0,106,0,521,21202,-2,-1,-2,22201,-4,-2,-4,109,-5,2106,0,0`
