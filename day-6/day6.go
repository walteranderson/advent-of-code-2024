package day6

import (
	"aoc/utils"
	"bytes"
	"errors"
	"fmt"
	"log"
	"strings"
)

func Run(filename string) {
	guardMap := make(GuardMap)
	startingPos := Position{}
	y := 0
	utils.ReadLines(filename, func(line string) {
		for x, ch := range line {
			guardMap[Position{x: x, y: y}] = ch
			if ch == '^' {
				startingPos.x = x
				startingPos.y = y
			}
		}
		y++
	})
	sum := utils.NewSet[Position]()
	sum.Add(startingPos)
	patrol := NewPatrol(startingPos, guardMap, North)
	patrol.Start(func(pos Position) {
		sum.Add(pos)
	})
	fmt.Printf("Part One: %d\n\n", sum.Size())

	sim := NewSimulator(patrol)
	sim.Simulate()
	fmt.Println(sim)
}

var LoopError = errors.New("LOOP")

type MapSimulation struct {
	guardMap GuardMap
	obstacle Position
}

type Simulator struct {
	origin    *Patrol
	maps      []MapSimulation
	loopCount int
}

func NewSimulator(origin *Patrol) *Simulator {
	return &Simulator{
		origin:    origin,
		maps:      nil,
		loopCount: 0,
	}
}

func (p *Simulator) String() string {
	var out bytes.Buffer
	out.WriteString("===SIMULATIONS===\n")
	out.WriteString(fmt.Sprintf("maps=%d\n", len(p.maps)))
	out.WriteString(fmt.Sprintf("Loop Count=%d\n", p.loopCount))
	return out.String()
}

func (s *Simulator) Simulate() {
	log.Println("Starting simulation")
	s.getTestMaps()
	log.Printf("Simulating %d Maps\n", len(s.maps))

	mset := utils.NewSet[Position]()
	for _, m := range s.maps {
		if mset.Contains(m.obstacle) {
			continue
		}
		mset.Add(m.obstacle)
		patrol := NewPatrol(s.origin.startingPos, m.guardMap, s.origin.startingDir)
		s.runPatrol(patrol, m.obstacle)
	}
}

func (s *Simulator) runPatrol(patrol *Patrol, obstacle Position) {
	err := patrol.Start(func(p Position) {})
	if errors.Is(err, LoopError) {
		fmt.Printf("Loop!: x=%d,y=%v\n", obstacle.x, obstacle.y)
		fmt.Println(patrol.Draw())
		s.loopCount++
	}
}

func (s *Simulator) getTestMaps() {
	for _, h := range s.origin.history {
		totest := []Position{
			{x: h.pos.x, y: h.pos.y - 1},
			{x: h.pos.x + 1, y: h.pos.y},
			{x: h.pos.x, y: h.pos.y + 1},
			{x: h.pos.x - 1, y: h.pos.y},
		}
		for _, pos := range totest {
			if s.origin.guardMap.Get(&pos) == '#' {
				continue
			}
			guardMap := make(GuardMap)
			for k, v := range s.origin.guardMap {
				guardMap[k] = v
			}
			guardMap[pos] = '0'

			m := MapSimulation{guardMap: guardMap, obstacle: pos}
			s.maps = append(s.maps, m)
		}
	}
}

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

type Position struct {
	x int
	y int
}

type GuardMap map[Position]rune

func (g GuardMap) Get(p *Position) rune {
	ch, ok := g[*p]
	if !ok {
		return 0
	}
	return ch
}

type PatrolHistory struct {
	pos          Position
	dir          Direction
	intersecting bool
}

type Patrol struct {
	pos         Position
	dir         Direction
	history     []PatrolHistory
	hSet        utils.Set[PatrolHistory]
	guardMap    GuardMap
	startingPos Position
	startingDir Direction
}

func NewPatrol(pos Position, guardMap GuardMap, dir Direction) *Patrol {
	return &Patrol{
		pos:         pos,
		dir:         dir,
		history:     nil,
		hSet:        utils.NewSet[PatrolHistory](),
		guardMap:    guardMap,
		startingPos: pos,
		startingDir: dir,
	}
}

func (p *Patrol) Start(onMove func(p Position)) error {
	for p.isWalking() {
		if p.isLoop() {
			return LoopError
		}
		switch p.Peek() {
		case '.', '^':
			p.Move()
			onMove(p.pos)
		case '#', '0':
			p.Turn()
		}
	}
	return nil
}

func (p *Patrol) isLoop() bool {
	ht := PatrolHistory{pos: p.pos, dir: p.dir, intersecting: true}
	hf := PatrolHistory{pos: p.pos, dir: p.dir, intersecting: false}
	if p.hSet.Contains(ht) || p.hSet.Contains(hf) {
		return true
	}
	return false
}

func (p *Patrol) Peek() rune {
	peek := Position{}
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
	hist := PatrolHistory{pos: p.pos, dir: p.dir, intersecting: false}
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
	for _, h := range p.history {
		if hist.pos == h.pos {
			hist.intersecting = true
		}
	}
	p.history = append(p.history, hist)
	p.hSet.Add(hist)
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
	p.history = append(p.history, PatrolHistory{pos: p.pos, dir: p.dir, intersecting: false})
}

func (p *Patrol) Draw() string {
	maxX, maxY := 0, 0
	for pos := range p.guardMap {
		if pos.x > maxX {
			maxX = pos.x
		}
		if pos.y > maxY {
			maxY = pos.y
		}
	}
	grid := make([][]rune, maxY+1)
	for i := range grid {
		grid[i] = make([]rune, maxX+1)
	}
	for pos, ch := range p.guardMap {
		grid[pos.y][pos.x] = ch
	}
	for _, h := range p.history {
		var ch rune
		switch h.dir {
		case North, South:
			ch = '|'
		case East, West:
			ch = '-'
		}
		if h.intersecting {
			ch = '+'
		}
		grid[h.pos.y][h.pos.x] = ch
	}

	grid[p.startingPos.y][p.startingPos.x] = '*'

	var ch rune
	switch p.dir {
	case North:
		ch = '^'
	case South:
		ch = 'v'
	case East:
		ch = '>'
	case West:
		ch = '<'
	}
	grid[p.pos.y][p.pos.x] = ch

	var sb strings.Builder
	for _, row := range grid {
		sb.WriteString(string(row))
		sb.WriteString("\n")
	}
	return sb.String()
}

func (p *Patrol) String() string {
	var out bytes.Buffer
	out.WriteString("===PATROL===\n")
	out.WriteString("Map: ")
	if len(p.guardMap) > 1000 {
		out.WriteString("REDACTED\n")
	} else {
		out.WriteString("\n")
		out.WriteString(p.Draw())
	}
	return out.String()
}
