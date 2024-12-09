package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

type Point struct {
	x int
	y int
}

type GuardMap map[Point]rune

func (g GuardMap) Get(p *Point) rune {
	ch, ok := g[*p]
	if !ok {
		return 0
	}
	return ch
}

type PointSet map[Point]struct{}

func (ps PointSet) Add(p Point) {
	if !ps.Contains(p) {
		ps[p] = struct{}{}
	}
}

func (ps PointSet) Contains(p Point) bool {
	_, exists := ps[p]
	return exists
}

func (ps PointSet) Size() int {
	return len(ps)
}

type Patrol struct {
	pos      Point
	dir      Direction
	guardMap GuardMap
	partOne  PointSet
}

func (p *Patrol) Peek() rune {
	peek := Point{}
	peek.x = p.pos.x
	peek.y = p.pos.y
	switch p.dir {
	case North:
		peek.y--
	case East:
		peek.x++
	case South:
		peek.y++
	case West:
		peek.x--
	}
	return p.guardMap.Get(&peek)
}

func (p *Patrol) isWalking() bool {
	return p.Peek() != 0
}

func (p *Patrol) Move() {
	switch p.dir {
	case North:
		p.pos.y--
	case East:
		p.pos.x++
	case South:
		p.pos.y++
	case West:
		p.pos.x--
	}
	p.partOne.Add(p.pos)
}

func (p *Patrol) Turn() {
	switch p.dir {
	case North:
		p.dir = East
	case East:
		p.dir = South
	case South:
		p.dir = West
	case West:
		p.dir = North
	}
}

func NewPatrol(filename string) *Patrol {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	guardMap := make(GuardMap)
	startingPos := Point{}
	y := 0
	for sc.Scan() {
		for x, ch := range sc.Text() {
			guardMap[Point{x: x, y: y}] = ch
			if ch == '^' {
				startingPos.x = x
				startingPos.y = y
			}
		}
		y++
	}
	partOne := PointSet{}
	partOne.Add(startingPos)
	return &Patrol{
		pos:      startingPos,
		dir:      North,
		guardMap: guardMap,
		partOne:  partOne,
	}
}

func (p *Patrol) String() string {
	return fmt.Sprintf("Position: %+v, Direction: %+v", p.pos, p.dir)
}

func Run(filename string) {
	patrol := NewPatrol(filename)
	for patrol.isWalking() {
		switch patrol.Peek() {
		case '.', '^':
			patrol.Move()
		case '#':
			patrol.Turn()
			patrol.Move()
		}
	}
	log.Printf("Part One: %d\n", patrol.partOne.Size())
}
