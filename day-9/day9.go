package day9 // TV

import (
	"aoc/utils"
	"bytes"
	"fmt"
	"log"
	"strconv"
)

func Run(filename string) {
	input := utils.ReadFile(filename)
	partOne := NewDiskMap(input)
	partOne.DefragPartOne()
	fmt.Printf("Part One: \n%s\n", partOne)
	partTwo := NewDiskMap(input)
	partTwo.DefragPartTwo()
	fmt.Printf("Part Two: \n%s\n", partTwo)
}

type File struct {
	ID   int // empty is -1
	Size int
}

func (f *File) IsEmpty() bool {
	return f.ID == -1
}

type DiskMap struct {
	input  string
	Files  []File
	Memory []int
}

func NewDiskMap(input string) *DiskMap {
	d := &DiskMap{
		input:  input,
		Files:  nil,
		Memory: nil,
	}

	d.createFilesFromInput()
	d.readIntoMemory()

	return d
}

func (d *DiskMap) DefragPartOne() {
	i := 0
	lasti := len(d.Memory) - 1
	for i < len(d.Memory) {
		if d.Memory[i] == -1 {
			for d.Memory[lasti] == -1 {
				lasti--
			}
			last := d.Memory[lasti]
			d.Memory[i] = last
			d.Memory[lasti] = -1
			lasti--
			i++
		} else {
			i++
		}
		if i == lasti {
			break
		}
	}
}

func (d *DiskMap) DefragPartTwo() {
	for fi := len(d.Files) - 1; fi >= 0; fi-- {
		file := d.Files[fi]
		if file.ID == -1 {
			continue
		}
		idx := d.findSpace(file)
		if idx == -1 {
			continue
		}
		d.moveFile(file, idx)
	}
}

func (d *DiskMap) findSpace(file File) int {
	for memi, mem := range d.Memory {
		if mem == file.ID {
			return -1
		}
		if mem != -1 {
			continue
		}
		toobig := false
		for i := 1; i < file.Size; i++ {
			if d.Memory[memi+i] != -1 {
				toobig = true
				break
			}
		}
		if toobig {
			continue
		}
		return memi
	}
	return -1
}

func (d *DiskMap) moveFile(file File, idx int) {
	mi := idx + file.Size
	for i := idx; i < mi; i++ {
		d.Memory[i] = file.ID
	}

	for mi < len(d.Memory) {
		if d.Memory[mi] == file.ID {
			d.Memory[mi] = -1
		}
		mi++
	}
}

func (d *DiskMap) createFilesFromInput() {
	isFile := false
	id := 0
	for _, ch := range d.input {
		if ch == '\n' {
			break
		}

		isFile = !isFile
		size, err := strconv.Atoi(string(ch))
		if err != nil {
			log.Fatal(err)
		}

		f := File{Size: size}
		if isFile {
			f.ID = id
			id++
		} else {
			f.ID = -1
		}
		d.Files = append(d.Files, f)
	}
}

func (d *DiskMap) readIntoMemory() {
	for _, file := range d.Files {
		for i := 0; i < file.Size; i++ {
			d.Memory = append(d.Memory, file.ID)
		}
	}
}

func (d *DiskMap) MemoryString() string {
	var out bytes.Buffer
	for _, m := range d.Memory {
		if m == -1 {
			out.WriteString(".")
		} else {
			out.WriteString(fmt.Sprintf("%d", m))
		}
	}
	return out.String()
}

func (d *DiskMap) CalculateChecksum() int {
	sum := 0
	for i, mem := range d.Memory {
		if mem == -1 {
			continue
		}
		sum += i * mem
	}
	return sum
}

func (d *DiskMap) String() string {
	var out bytes.Buffer
	out.WriteString("  Input    -> ")
	if len(d.input) > 100 {
		out.WriteString("REDACTED - TOO BIG")
	} else {
		out.WriteString(d.input)
	}
	out.WriteString("\n  Files    -> ")
	if len(d.Files) > 100 {
		out.WriteString("REDACTED - TOO BIG")
	} else {
		for _, f := range d.Files {
			for i := 0; i < f.Size; i++ {
				if f.ID == -1 {
					out.WriteString(".")
				} else {
					out.WriteString(fmt.Sprintf("%d", f.ID))
				}
			}
		}
	}

	out.WriteString("\n  Memory   -> ")
	if len(d.Memory) > 1000 {
		out.WriteString("REDACTED - TOO BIG")
	} else {
		out.WriteString(d.MemoryString())
	}

	out.WriteString(fmt.Sprintf("\n  Checksum -> %d\n", d.CalculateChecksum()))

	return out.String()
}
