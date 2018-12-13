[Link to the puzzle](https://adventofcode.com/2018/day/1)

> Starting with a frequency of zero, what is the resulting frequency after all of the changes in frequency have been applied?

The result can be obtained by calculating a simple running total:

```go
func main() {
	values := readInts(readLines("input.txt"))
	var result int
	for _, v := range values {
		result += v
	}
	fmt.Printf("RESULT: %d\n", result)
}
```
