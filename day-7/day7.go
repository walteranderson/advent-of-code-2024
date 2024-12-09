package day7

import (
	"aoc/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var operatorsPartOne = []rune{'+', '*'}
var operatorsPartTwo = []rune{'+', '*', '|'}

func Run(filename string) {
	calibrations := make([]calibration, 0)
	utils.ReadLines(filename, func(line string) {
		calibrations = append(calibrations, newCalibration(line))
	})

	partOne := 0
	partTwo := 0
	for _, c := range calibrations {
		if c.isValid(operatorsPartOne) {
			partOne += c.Result
		}
		if c.isValid(operatorsPartTwo) {
			partTwo += c.Result
		}
	}
	log.Printf("Part One: %d\n", partOne)
	log.Printf("Part Two: %d\n", partTwo)
}

type calibration struct {
	Result   int
	Operands []int
}

func newCalibration(line string) calibration {
	parts := strings.Split(line, ":")
	result, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal(err)
	}
	ops := make([]int, 0)
	for _, op := range strings.Split(strings.Trim(parts[1], " "), " ") {
		opi, err := strconv.Atoi(op)
		if err != nil {
			log.Fatal(err)
		}
		ops = append(ops, opi)
	}
	return calibration{
		Result:   result,
		Operands: ops,
	}
}

func (c *calibration) isValid(ops []rune) bool {
	possibles := c.generateOperators(ops, len(c.Operands)-1)
	for _, operators := range possibles {
		res := 0
		for i := 0; i < len(c.Operands); i++ {
			if i+1 >= len(c.Operands) {
				break
			}
			if i == 0 {
				res = doMath(c.Operands[0], c.Operands[1], operators[0])
				continue
			}
			next := c.Operands[i+1]
			temp := doMath(res, next, operators[i])
			res = temp
		}
		if c.Result == res {
			return true
		}
	}
	return false
}

func (c *calibration) generateOperators(ops []rune, length int) [][]rune {
	if length == 0 {
		return [][]rune{{}}
	}
	sub := c.generateOperators(ops, length-1)
	perms := make([][]rune, 0)
	for _, s := range sub {
		for _, op := range ops {
			newP := append([]rune{op}, s...)
			perms = append(perms, newP)
		}
	}
	return perms
}

func doMath(left, right int, operator rune) int {
	switch operator {
	case '*':
		return left * right
	case '+':
		return left + right
	case '|':
		return concatInt(left, right)
	}
	panic(fmt.Sprintf("Undefined operator: %s", string(operator)))
}

func concatInt(left, right int) int {
	l := strconv.Itoa(left)
	r := strconv.Itoa(right)
	concatted := l + r
	converted, err := strconv.Atoi(concatted)
	if err != nil {
		log.Fatal(err)
	}
	return converted
}
