package utils

import (
	"bufio"
	"io"
	"log"
	"os"
)

func ReadFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func ReadLines(filename string, callback func(string)) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		callback(sc.Text())
	}
	if err = sc.Err(); err != nil {
		log.Fatal(err)
	}
}
