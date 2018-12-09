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
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			lineI, lineJ := lines[i], lines[j]

			hasDiff, idx := hasSingleDifference(lineI, lineJ)
			if !hasDiff {
				continue
			}

			var runes []rune
			for k := 0; k < len(lineI); k++ {
				if k != idx {
					runes = append(runes, rune(lineI[k]))
				}
			}
			fmt.Printf("RESULT: %s\n", string(runes))
			os.Exit(0)
		}
	}
}

func hasSingleDifference(s1, s2 string) (bool, int) {
	var hasSingleDiff bool
	var diffIdx int
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			continue
		}

		if hasSingleDiff {
			hasSingleDiff = false
			return false, -1
		}

		hasSingleDiff = true
		diffIdx = i
	}
	return hasSingleDiff, diffIdx
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
