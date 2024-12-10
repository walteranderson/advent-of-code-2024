package day10

import (
	"aoc/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Run(filename string) {
	lines := make([]string, 0)
	utils.ReadLines(filename, func(line string) {
		lines = append(lines, line)
	})
	grid := NewGrid(lines)
	topo := NewTopoMap(grid)
	topo.Start()
	fmt.Printf("Part One: %d\n", topo.partOneScore)
	fmt.Printf("Part Two: %d\n", topo.partTwoScore)
}

type Position struct {
	x int
	y int
}

type Grid struct {
	state map[Position]int
	size  Position
}

func NewGrid(lines []string) Grid {
	grid := Grid{
		state: make(map[Position]int),
		size:  Position{x: 0, y: 0},
	}
	maxX := 0
	maxY := len(lines)
	for y, line := range lines {
		if len(line) > maxX {
			maxX = len(line)
		}
		for x, ch := range line {
			n, err := strconv.Atoi(string(ch))
			if err != nil {
				log.Fatal(err)
			}
			pos := Position{x: x, y: y}
			grid.state[pos] = n
		}
	}
	grid.size.x = maxX
	grid.size.y = maxY
	return grid
}

func (g *Grid) Get(pos Position) int {
	val, ok := g.state[pos]
	if !ok {
		return -1
	}
	return val
}

func (g *Grid) FindNext(pos Position, searchTerm int) []Position {
	search := []Position{
		{x: pos.x, y: pos.y - 1},
		{x: pos.x + 1, y: pos.y},
		{x: pos.x, y: pos.y + 1},
		{x: pos.x - 1, y: pos.y},
	}
	valid := make([]Position, 0)
	for _, p := range search {
		if g.Get(p) == searchTerm {
			valid = append(valid, p)
		}
	}
	return valid
}

type Trail struct {
	start           Position
	partOnePath     utils.Set[Position]
	partOneEndsSeen utils.Set[Position]
	partOneScore    int
	partTwoScore    int
}

func NewTrail(start Position) *Trail {
	return &Trail{
		start:           start,
		partOnePath:     utils.NewSet[Position](),
		partOneEndsSeen: utils.NewSet[Position](),
		partOneScore:    0,
	}
}

type TopoMap struct {
	grid         Grid
	partOneScore int
	partTwoScore int
}

func NewTopoMap(grid Grid) *TopoMap {
	return &TopoMap{
		grid:         grid,
		partOneScore: 0,
	}
}

func (t *TopoMap) Start() {
	for pos, val := range t.grid.state {
		if val == 0 {
			trail := NewTrail(pos)
			t.Walk(trail, trail.start, 1, []Position{trail.start})
			if trail.partOneScore > 0 {
				t.partOneScore += trail.partOneScore
			}
			if trail.partTwoScore > 0 {
				t.partTwoScore += trail.partTwoScore
			}
		}
	}
}

func (t *TopoMap) Walk(trail *Trail, current Position, search int, history []Position) {
	if search > 9 {
		if !trail.partOneEndsSeen.Contains(current) {
			trail.partOneEndsSeen.Add(current)
			trail.partOneScore++
			for _, h := range history {
				trail.partOnePath.Add(h)
			}
		}
		trail.partTwoScore++
		return
	}

	next := t.grid.FindNext(current, search)
	if len(next) == 0 {
		return
	}

	search++
	for _, pos := range next {
		history = append(history, pos)
		t.Walk(trail, pos, search, history)
	}
}

func (t *TopoMap) Draw() string {
	var sb strings.Builder
	for i := 0; i < t.grid.size.y; i++ {
		for j := 0; j < t.grid.size.x; j++ {
			sb.WriteString(fmt.Sprintf("%d", t.grid.state[Position{x: j, y: i}]))
		}
		if i < t.grid.size.y-1 {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}

func (t *TopoMap) DrawTrail(trail *Trail) string {
	var sb strings.Builder
	for i := 0; i < t.grid.size.y; i++ {
		for j := 0; j < t.grid.size.x; j++ {
			pos := Position{x: j, y: i}
			el := t.grid.Get(pos)
			if el == -1 {
				continue
			}
			if trail.partOnePath.Contains(pos) {
				sb.WriteString(fmt.Sprintf("%d", el))
			} else {
				sb.WriteString(".")
			}
		}
		if i < t.grid.size.y-1 {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}

// func (t *TopoMap) Traverse() {
// 	for pos, val := range t.grid.state {
// 		if val != 0 {
// 			continue
// 		}
//
// 		// trail := t.Walk(pos)
// 		// log.Printf("Trail: %+v\n", trail)
// 		panic("TODO")
//
// 		// count := t.follow(pos, 1, 0)
// 		// if count > 0 {
// 		// 	log.Printf("Count for %+v = %d\n", pos, count)
// 		// 	t.trailheadScores = append(t.trailheadScores, count)
// 		// }
// 	}
// }

// func (t *TopoMap) follow(pos Position, search int, count int) int {
// 	if search == 9 {
// 		return count + 1
// 	}
// 	positions := t.grid.FindNext(pos, search)
// 	if len(positions) == 0 {
// 		return count
// 	}
// 	search++
// 	for _, p := range positions {
// 		count = t.follow(p, search, count)
// 	}
// 	return count
// }

// func (t *TopoMap) Score() int {
// 	sum := 0
// 	// for _, score := range t.trailheadScores {
// 	// 	sum += score
// 	// }
// 	return sum
// }
