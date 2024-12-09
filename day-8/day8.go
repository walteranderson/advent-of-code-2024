package day8

import (
	"aoc/utils"
	"bytes"
	"fmt"
	"strings"
)

func Run(filename string) {
	city := newCityMap()
	y := 0
	utils.ReadLines(filename, func(line string) {
		if city.size.x == 0 {
			city.size.x = len(line)
		}

		for x, ch := range line {
			p := point{x: x, y: y}
			city.graph[p] = ch
			if ch != '.' {
				city.antennas[ch] = append(city.antennas[ch], p)
			}
		}
		y++
	})
	city.size.y = y

	for _, points := range city.antennas {

		for i, p := range points {
			if len(points) > 1 {
				city.antinodesPart2.Add(p)
			}
			others := utils.RemoveElementFromSlice(points, i)
			for _, tocmp := range others {
				antinode := p.getAntinode(tocmp)
				if city.isOnGraph(antinode) {
					city.antinodesPart1.Add(antinode)
					city.antinodesPart2.Add(antinode)
				}

				cur := antinode
				prev := p
				cont := cur.getAntinode(prev)
				for city.isOnGraph(cont) {
					city.antinodesPart2.Add(cont)
					prev = cur
					cur = cont
					cont = cur.getAntinode(prev)
				}
			}
		}
	}
	fmt.Println(city)
}

type point struct {
	x int
	y int
}

func (p *point) getAntinode(t point) point {
	xdiff := p.x - t.x
	ydiff := p.y - t.y
	return point{
		x: p.x + xdiff,
		y: p.y + ydiff,
	}
}

type cityMap struct {
	size           point
	graph          map[point]rune
	antennas       map[rune][]point
	antinodesPart1 utils.Set[point]
	antinodesPart2 utils.Set[point]
}

func newCityMap() *cityMap {
	return &cityMap{
		size:           point{x: 0, y: 0},
		graph:          make(map[point]rune),
		antennas:       make(map[rune][]point),
		antinodesPart1: utils.NewSet[point](),
		antinodesPart2: utils.NewSet[point](),
	}
}

func (cm *cityMap) isOnGraph(p point) bool {
	return p.x >= 0 &&
		p.x < cm.size.x &&
		p.y >= 0 &&
		p.y < cm.size.y
}

func (cm *cityMap) drawGraph() string {
	grid := make([][]rune, cm.size.y)
	for i := range grid {
		grid[i] = make([]rune, cm.size.x)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	for p, r := range cm.graph {
		char := r
		if r == '.' && cm.antinodesPart2.Contains(p) {
			char = '#'
		}
		grid[p.y][p.x] = char
	}
	var sb strings.Builder
	for _, row := range grid {
		sb.WriteString(fmt.Sprintf("%s\n", string(row)))
	}

	return sb.String()
}

func (cm *cityMap) String() string {
	var out bytes.Buffer
	out.WriteString(cm.drawGraph())
	// out.WriteString("City Map:\n")
	// out.WriteString(fmt.Sprintf("  size: %dx%d\n", cm.size.x, cm.size.y))
	// out.WriteString("Antennas:\n")
	// for ch, points := range cm.antennas {
	// 	out.WriteString(fmt.Sprintf("  %s -> %+v\n", string(ch), points))
	// }
	out.WriteString("Antinodes Count: \n")
	out.WriteString(fmt.Sprintf("  Part One: %d\n", cm.antinodesPart1.Size()))
	out.WriteString(fmt.Sprintf("  Part Two: %d\n", cm.antinodesPart2.Size()))
	// out.WriteString(fmt.Sprintf("  nodes: %s\n", cm.antinodesPart1.String()))
	return out.String()
}
