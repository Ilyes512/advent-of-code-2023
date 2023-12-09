package main

import (
	"fmt"

	"github.com/Ilyes512/advent-of-code-2023/advent/day05"
)

func main() {
	reader := day05.NewReader("./solutions/day05/input.txt")

	fmt.Printf("Part 1 answer: %d\n", reader.ProcessPart1())
	fmt.Printf("Part 2 answer: %d\n", reader.ProcessPart2())
}
