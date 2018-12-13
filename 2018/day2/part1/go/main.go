package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	lines := readLines("input.txt")
	var count2, count3 int

	for _, line := range lines {
		distrib := lettersDistribution(line)

		if hasExactly(distrib, 2) {
			count2++
		}
		if hasExactly(distrib, 3) {
			count3++
		}
	}

	fmt.Printf("RESULT: %d\n", count2*count3)
}

func lettersDistribution(line string) map[rune]int {
	occurences := make(map[rune]int)
	for _, letter := range line {
		v, ok := occurences[letter]
		if !ok {
			v = 0
		}
		occurences[letter] = v + 1
	}
	return occurences
}

func hasExactly(distrib map[rune]int, occurences int) bool {
	for _, count := range distrib {
		if occurences == count {
			return true
		}
	}
	return false
}

func readLines(name string) []string {
	file, err := os.Open(name)
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
