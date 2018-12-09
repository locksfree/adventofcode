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
	lines := readLines()
	var result int
	for _, line := range lines {
		v, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Could not convert %s to an int", line)
		}
		result += v
	}
	fmt.Printf("RESULT: %d\n", result)
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
