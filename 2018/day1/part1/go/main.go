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
	values := readInts(readLines("input.txt"))
	var result int
	for _, v := range values {
		result += v
	}
	fmt.Printf("RESULT: %d\n", result)
}

func readInts(entries []string) []int {
	var res []int
	for _, e := range entries {
		e = strings.Replace(e, "\n", "", -1)
		i, err := strconv.Atoi(e)
		if err != nil {
			log.Fatalf("Could not convert %s to an int: %v", e, err)
		}
		res = append(res, i)
	}

	return res
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
