package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readInts(readLines())

	results := makeSet()
	results.store(0)

	var runTot int
	for {
		for _, v := range lines {
			runTot += v

			if results.has(runTot) {
				fmt.Printf("RESULT: %d\n", runTot)
				os.Exit(0)
			}

			results.store(runTot)
		}
	}
}

func makeSet() *set {
	return &set{bag: make(map[int]struct{})}
}

type set struct {
	bag map[int]struct{}
}

func (s *set) has(object int) bool {
	_, ok := s.bag[object]
	return ok
}

func (s *set) store(object int) {
	s.bag[object] = struct{}{}
}

func readLines() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not open input file")
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Could not read content of input file")
	}

	return strings.Split(string(content), "\r\n")
}

func readInts(lines []string) []int {
	var results []int
	for _, line := range lines {
		v, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Could not convert %s to an int", line)
		}
		results = append(results, v)
	}
	return results
}
