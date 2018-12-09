package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	lines := readLines()
	var count2, count3 int

	for _, line := range lines {
		occurences := lettersDistribution(line)

		var has2, has3 bool
		for _, count := range occurences {
			if count == 2 {
				has2 = true
			} else if count == 3 {
				has3 = true
			}
		}

		if has2 {
			count2++
		}
		if has3 {
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
