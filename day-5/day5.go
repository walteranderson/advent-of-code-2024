package day5

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Rules map[int][]int

func (rules Rules) add(key, val int) {
	if _, exists := rules[key]; exists {
		rules[key] = append(rules[key], val)
	} else {
		rules[key] = []int{val}
	}
}

func (rules Rules) isValid(n int, has []int) (bool, int) {
	spots, ok := rules[n]
	if !ok {
		return true, -1
	}
	for _, spot := range spots {
		for _, h := range has {
			if spot == h {
				return false, h
			}
		}
	}
	return true, -1
}

type Instructions []int

func (ins Instructions) getMiddle() int {
	if len(ins) <= 0 {
		panic("you passed an empty slice to getMiddle")
	}
	if len(ins)%2 == 0 {
		panic("you passed an even numbered slice to getMiddle. is this valid?")
	}
	idx := len(ins) / 2
	return ins[idx]
}

func (ins Instructions) fixAndGetMiddle(rules Rules) int {
	sorted := make(Instructions, len(ins))
	copy(sorted, ins)

	i := 1
	for i < len(sorted) {
		cur := sorted[i]
		valid, invalidVal := rules.isValid(cur, sorted.getPastValues(i))
		if valid {
			i++
		} else {
			invalidIdx := sorted.indexOf(invalidVal)
			sorted[invalidIdx] = cur
			sorted[i] = invalidVal
			i = 1
		}
	}

	return sorted.getMiddle()
}

func (ins Instructions) getPastValues(idx int) []int {
	return ins[:idx]
}

func (ins Instructions) indexOf(val int) int {
	for i, v := range ins {
		if v == val {
			return i
		}
	}
	return -1
}

func Run(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	processRules := true
	rules := make(Rules, 0)
	instructions := make([]Instructions, 0)
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			processRules = false
			continue
		}

		if processRules {
			parts := strings.Split(line, "|")
			a, _ := strconv.Atoi(parts[0])
			b, _ := strconv.Atoi(parts[1])
			rules.add(a, b)
		} else {
			parts := strings.Split(line, ",")
			ins := make(Instructions, 0)
			for _, p := range parts {
				intP, _ := strconv.Atoi(p)
				ins = append(ins, intP)
			}
			instructions = append(instructions, ins)
		}
	}

	partOne := 0
	partTwo := 0
	for _, ins := range instructions {
		fixOrder := false
		for idx, cur := range ins {
			if idx == 0 {
				continue
			}
			valid, _ := rules.isValid(cur, ins.getPastValues(idx))
			if !valid {
				fixOrder = true
				break
			}
		}
		if fixOrder {
			partTwo += ins.fixAndGetMiddle(rules)
		} else {
			partOne += ins.getMiddle()
		}
	}
	log.Printf("Part One: %d\n", partOne)
	log.Printf("Part Two: %d\n", partTwo)
}
