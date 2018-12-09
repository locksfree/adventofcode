package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	rects := readRects(readLines())

	var taken []rect
	for i := 0; i < len(rects); i++ {
		for j := i + 1; j < len(rects); j++ {
			r1, r2 := rects[i], rects[j]

			r := r1.intersect(r2)
			if !r.empty() {
				taken = append(taken, r)
			}
		}
	}

	used := makeSet()
	for _, t := range taken {
		for _, c := range t.cells() {
			used.store(c)
		}
	}

	fmt.Printf("RESULT: %d\n", used.size())
}

type rect struct {
	id         string
	x, y, w, h int
}

func (a rect) empty() bool {
	return a.w == 0 && a.h == 0
}

func (a rect) intersect(b rect) rect {
	r1, r2 := vsortRect(a, b)
	x1, y1 := intMax(r1.x, r2.x), intMax(r1.y, r2.y)
	x2, y2 := intMin(r1.x+r1.w, r2.x+r2.w), intMin(r1.y+r1.h, r2.y+r2.h)
	if x2 <= x1 || y2 <= y1 {
		return rect{}
	}
	return rect{id: "", x: x1, y: y1, w: x2 - x1, h: y2 - y1}
}

func intMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func intMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func vsortRect(a, b rect) (rect, rect) {
	if a.y > b.y {
		return b, a
	}
	if a.y == b.y {
		if a.x >= b.x {
			return b, a
		}
	}
	return a, b
}

type cell struct {
	x, y int
}

func (a rect) cells() []cell {
	var result []cell
	for i := a.x; i < a.x+a.w; i++ {
		for j := a.y; j < a.y+a.h; j++ {
			result = append(result, cell{i, j})
		}
	}
	return result
}

func readRects(lines []string) []rect {
	var rects []rect
	for _, line := range lines {
		var id string
		var x, y, w, h int

		_, err := fmt.Sscanf(line, "#%s @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
		if err != nil {
			log.Fatalf("Error scanning the input: %v", err)
		}
		rects = append(rects, rect{id: id, x: x, y: y, w: w, h: h})
	}
	return rects
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

func makeSet() *set {
	return &set{bag: make(map[cell]struct{})}
}

type set struct {
	bag map[cell]struct{}
}

func (s *set) has(object cell) bool {
	_, ok := s.bag[object]
	return ok
}

func (s *set) store(object cell) {
	s.bag[object] = struct{}{}
}

func (s *set) size() int {
	return len(s.bag)
}
