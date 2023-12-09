package main

import (
	"fmt"

	"github.com/Ilyes512/advent-of-code-2023/advent/day06"
)

func main() {
	part1 := day06.NewRecordsPart1("./solutions/day06/input.txt").Result()
	part2 := day06.NewRecordsPart2("./solutions/day06/input.txt").Result()

	fmt.Printf("Part 1 answer: %d\n", part1)
	fmt.Printf("Part 2 answer: %d\n", part2)
}
