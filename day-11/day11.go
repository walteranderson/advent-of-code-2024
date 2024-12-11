package day11

import (
	"aoc/utils"
	"log"
	"strconv"
	"strings"
)

func Run(filename string) {
	contents := utils.ReadFile(filename)
	parts := strings.Split(contents, " ")
	input := make([]int, 0)
	for _, part := range parts {
		part := strings.ReplaceAll(part, "\n", "")
		n, err := strconv.Atoi(part)
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, n)
	}
	stones1 := StoneList{input: input, countMap: make(map[int]int)}
	stones1.Blink(25)
	log.Printf("Part One: %d\n", stones1.GetCount())

	stones2 := StoneList{input: input, countMap: make(map[int]int)}
	stones2.Blink(75)
	log.Printf("Part Two: %d\n", stones2.GetCount())
}

type StoneList struct {
	input    []int
	countMap map[int]int
}

func (sl *StoneList) Blink(count int) {
	for _, s := range sl.input {
		sl.countMap[s] = 1
	}

	for i := 0; i < count; i++ {
		cm := mapCopy(sl.countMap)
		for stone, count := range sl.countMap {
			if count == 0 {
				continue
			}
			if stone == 0 {
				onecount := cm[1]
				cm[1] = onecount + count
				cm[stone] = cm[stone] - count
			} else if digitCountIsEven(stone) {
				left, right := splitDigit(stone)
				cm[left] = cm[left] + count
				cm[right] = cm[right] + count
				cm[stone] = cm[stone] - count
			} else {
				newVal := stone * 2024
				cm[newVal] = cm[newVal] + count
				cm[stone] = cm[stone] - count
			}
		}
		sl.countMap = cm
	}
}

func (sl *StoneList) GetCount() int {
	sum := 0
	for _, count := range sl.countMap {
		sum += count
	}
	return sum
}

func mapCopy(input map[int]int) map[int]int {
	output := make(map[int]int, len(input))
	for i, v := range input {
		output[i] = v
	}
	return output
}

var cache = make(map[int][2]int)

func splitDigit(input int) (int, int) {
	cached, ok := cache[input]
	if ok {
		return cached[0], cached[1]
	}

	str := strconv.Itoa(input)
	mid := len(str) / 2
	left := str[:mid]
	right := str[mid:]
	leftI, err := strconv.Atoi(left)
	if err != nil {
		log.Fatal(err)
	}
	rightI, err := strconv.Atoi(right)
	if err != nil {
		log.Fatal(err)
	}
	cache[input] = [2]int{leftI, rightI}
	return leftI, rightI
}

func digitCountIsEven(input int) bool {
	return len(strconv.Itoa(input))%2 == 0
}
