package main

import (
	day1 "aoc/day-1"
	day2 "aoc/day-2"
	day3 "aoc/day-3"
	day4 "aoc/day-4"
	day5 "aoc/day-5"
	day6 "aoc/day-6"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: aoc <day> <filename>")
	}
	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	filename, err := getFilename(day, os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	run(day, filename)
}

func run(day int, filename string) {
	switch day {
	case 1:
		day1.Run(filename)
	case 2:
		day2.Run(filename)
	case 3:
		day3.Run(filename)
	case 4:
		day4.Run(filename)
	case 5:
		day5.Run(filename)
	case 6:
		day6.Run(filename)
	}
}

func getFilename(day int, filename string) (string, error) {
	path := fmt.Sprintf("day-%d/%s.txt", day, filename)
	return filepath.Abs(path)
}
