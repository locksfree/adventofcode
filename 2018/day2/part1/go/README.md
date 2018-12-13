[Link to Puzzle](https://adventofcode.com/2018/day/2)

> counting the number that have an ID containing exactly two of any letter and then separately counting those with exactly three of any letter. You can multiply those two counts together to get a rudimentary checksum

The first step is to find the number of time each letter occurs within an ID:

```go
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
```

Follows then count of the occurences as described in the puzzle, any ID will have:
 
 * no repeating letters (ignored)
 * one or more letter repeating twice (it just count once as 1 ID)
 * one or more letter repeating thrice (it just count once as 1 ID)
 * letters repeating more times (ignored)
 
for this we use a helper indicator function that returns true when at least one character has the exact requested number of repetitions:

```go
func hasExactly(distrib map[rune]int, occurences int) bool {
	for _, count := range distrib {
		if occurences == count {
			return true
		}
	}
	return false
}
```
 
with which the number of IDs with exactly 2 and exactly 3 repetition of the same letter can be found and multiplied together to produce the answer:
 
```go
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
```

[View Full Code](./main.go)
