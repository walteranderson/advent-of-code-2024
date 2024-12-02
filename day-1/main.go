package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Provide input")
	}
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	var first []int
	var second []int
	freq := map[int]int{}
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		parts := strings.Split(line, " ")
		left, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		right, err := strconv.Atoi(parts[len(parts)-1])
		if err != nil {
			log.Fatal(err)
		}
		first = append(first, left)
		second = append(second, right)

		_, ok := freq[right]
		if ok {
			freq[right] += 1
		} else {
			freq[right] = 1
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	sort.Ints(first)
	sort.Ints(second)
	if len(first) != len(second) {
		log.Fatal("length of the two lists are not the same")
	}

	sum := 0
	for i := 0; i < len(first); i++ {
		left := first[i]
		leftFreq, ok := freq[left]
		if !ok {
			leftFreq = 0
		}
		sum += left * leftFreq
		// sum += int(math.Abs(float64(first[i] - second[i])))
	}
	log.Printf("Sum = %d", sum)
}
