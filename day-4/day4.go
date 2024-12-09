package day4

import (
	"bufio"
	"log"
	"os"
)

type Direction int

const (
	NW Direction = iota
	N
	NE
	W
	E
	SW
	S
	SE
)

type Point struct {
	x int
	y int
}

func (p *Point) Next(dir Direction) Point {
	switch dir {
	case NW:
		return Point{x: p.x - 1, y: p.y - 1}
	case N:
		return Point{x: p.x, y: p.y - 1}
	case NE:
		return Point{x: p.x + 1, y: p.y - 1}
	case W:
		return Point{x: p.x - 1, y: p.y}
	case E:
		return Point{x: p.x + 1, y: p.y}
	case SW:
		return Point{x: p.x - 1, y: p.y + 1}
	case S:
		return Point{x: p.x, y: p.y + 1}
	case SE:
		return Point{x: p.x + 1, y: p.y + 1}
	}
	panic("Tried to call point.Next() with an invalid direction")
}

func (p *Point) All() []Neighbor {
	dirs := []Direction{NW, N, NE, W, E, SW, S, SE}
	neighbors := make([]Neighbor, 8)
	for i, dir := range dirs {
		n := Neighbor{
			dir:   dir,
			point: p.Next(dir),
		}
		neighbors[i] = n
	}
	return neighbors
}

type Data map[Point]string

type Neighbor struct {
	dir   Direction
	point Point
}

func (d Data) Get(p Point) string {
	v, ok := d[p]
	if !ok {
		return ""
	}
	return v
}

func Run(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	data := make(Data)
	y := 0
	for sc.Scan() {
		for x, ch := range sc.Text() {
			data[Point{x: x, y: y}] = string(ch)
		}
		y++
	}

	firstCount := partOne(&data)
	secondCount := partTwo(&data)
	log.Printf("Part One = %d\n", firstCount)
	log.Printf("Part Two = %d\n", secondCount)
}

func partOne(data *Data) int {
	count := 0
	found := make([]Point, 0)
	for p, v := range *data {
		if v == "X" {
			neighbors := p.All()
			for _, n := range neighbors {
				if data.Get(n.point) == "M" {
					np := n.point.Next(n.dir)
					if data.Get(np) == "A" {
						nnp := np.Next(n.dir)
						if data.Get(nnp) == "S" {
							found = append(found, p, n.point, np, nnp)
							count++
						}
					}
				}
			}
		}
	}
	return count
}

func partTwo(data *Data) int {
	count := 0
	for p, v := range *data {
		if v == "A" {
			nw := data.Get(p.Next(NW))
			ne := data.Get(p.Next(NE))
			sw := data.Get(p.Next(SW))
			se := data.Get(p.Next(SE))
			if (nw == "M" && se == "S") || (nw == "S" && se == "M") {
				if (ne == "M" && sw == "S") || (ne == "S" && sw == "M") {
					count++
				}
			}
		}
	}
	return count
}
