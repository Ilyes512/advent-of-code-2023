package main

import (
	"fmt"

	"github.com/Ilyes512/advent-of-code-2023/advent/day07"
)

func main() {
	hands, handsWithJoker := day07.NewHands("./solutions/day07/input.txt")

	fmt.Printf("Part 1 answer: %d\n", hands.Result())
	fmt.Printf("Part 2 answer: %d\n", handsWithJoker.Result())
}
