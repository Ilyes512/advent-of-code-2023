package main

import (
	"fmt"

	"github.com/Ilyes512/advent-of-code-2023/advent/day05"
)

func main() {
	lowestLocation := day05.NewReader("./solutions/day05/input.txt").Process()

	fmt.Printf("Part 1 answer: %d\n", lowestLocation)
}
