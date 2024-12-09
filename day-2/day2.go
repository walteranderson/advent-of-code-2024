package day2

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Run(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	partOneSafe := 0
	partTwoSafe := 0
	for sc.Scan() {
		line := strings.Split(sc.Text(), " ")
		if isSafe(line) {
			partOneSafe++
			partTwoSafe++
		} else {
			for i := 0; i < len(line); i++ {
				subset := remove(line, i)
				if isSafe(subset) {
					partTwoSafe++
					break
				}
			}
		}
	}
	log.Printf("Safe part 1 = %d", partOneSafe)
	log.Printf("Safe part 2 = %d", partTwoSafe)
}

func isSafe(line []string) bool {
	direction := 0
	for i := 0; i < len(line); i++ {
		level := line[i]
		if i+1 >= len(line) {
			break
		}
		peek, err := strconv.Atoi(line[i+1])
		if err != nil {
			log.Fatal(err)
		}
		cur, err := strconv.Atoi(level)
		if err != nil {
			log.Fatal(err)
		}
		diff := cur - peek
		if diff == 0 {
			return false
		}

		absDiff := int(math.Abs(float64(diff)))
		if absDiff < 1 || absDiff > 3 {
			return false
		}

		if direction == 0 {
			direction = diff
		} else {
			if direction*diff < 0 {
				return false
			}
		}
	}
	return true
}

func remove(arr []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, arr[:index]...)
	return append(ret, arr[index+1:]...)
}
