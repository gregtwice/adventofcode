package pkg

import (
	"fmt"
)

const (
	UP = iota
	LEFT
	DOWN
	RIGHT
)

var DirectionMap = map[int]struct {
	x, y int
	name string
}{
	UP:    {0, 1, "up"},
	LEFT:  {1, 0, "left"},
	DOWN:  {0, -1, "down"},
	RIGHT: {-1, 0, "right"},
}

type P struct {
	X, Y             int
	CurrentDirection int
	Steps            int
}

func NewPoint() *P {
	return &P{}
}

func (p P) String() string {
	return fmt.Sprintf("[%d,%d]", p.Y, p.X)
}

func (p *P) MoveLeft(steps int) {
	p.CurrentDirection = RIGHT
	p.Move(steps)
}

func (p *P) MoveRight(steps int) {
	p.CurrentDirection = LEFT
	p.Move(steps)
}

func (p *P) MoveUp(steps int) {
	p.CurrentDirection = UP
	p.Move(steps)
}

func (p *P) MoveDown(steps int) {
	p.CurrentDirection = DOWN
	p.Move(steps)
}

func (p *P) Move(steps int) {
	p.X += DirectionMap[p.CurrentDirection].x * steps
	p.Y += DirectionMap[p.CurrentDirection].y * steps
	p.Steps += steps
}

func (p *P) TurnLeft() {
	p.CurrentDirection = (p.CurrentDirection - 1 + 4) % 4
}

func (p *P) TurnRight() {
	p.CurrentDirection = (p.CurrentDirection + 1) % 4
}

func (p *P) DistFromOrigin() int {
	return p.DistFrom(&P{X: 0, Y: 0})
}

func (p *P) DistFrom(from *P) int {
	return Abs(p.X-from.X) + Abs(p.Y-from.Y)
}

func (p *P) FindInMap(stringMap [][]string, origin P) string {
	newY := origin.Y - p.Y
	if newY >= len(stringMap) || newY < 0 {
		panic(fmt.Errorf("out or range Y:%d", p.Y))
	}

	newX := origin.X + p.X
	if newX >= len(stringMap[newY]) || newX < 0 {
		panic(fmt.Errorf("out or range X:%d", p.X))
	}

	return stringMap[newY][newX]
}

func (p *P) FourWays() []P {
	var dirs = []int{DOWN, RIGHT, LEFT, UP}

	adjs := make([]P, 4)
	for i, dir := range dirs {
		newP := P{X: p.X, Y: p.Y, CurrentDirection: dir}
		newP.Move(1)
		newP.CurrentDirection = 0
		adjs[i] = newP
	}
	return adjs
}
