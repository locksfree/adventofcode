[Link to Puzzle](https://adventofcode.com/2018/day/1)

> To calibrate the device, you need to find the first frequency it reaches twice.

This works the same way, except that we need to keep every partial sum to know when any of them repeats, and we need to keep summing until we find a repeat value. For fast retrieval, a hash map/set is best here.

```go
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

```

[Full Code](./main.go)
