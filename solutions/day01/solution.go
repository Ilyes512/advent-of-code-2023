package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Ilyes512/advent-of-code-2023/advent/day01"
)

func main() {
	contents, err := os.ReadFile("./solutions/day01/input.txt")
	if err != nil {
		log.Panic(err)
	}

	document := day01.NewDocument(string(contents))

	answerPart1, err := document.Calculate()
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Answer Part 1 is: %d\n", answerPart1)

	answerPart2, err := document.Translate().Calculate()
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Answer Part 2 is: %d\n", answerPart2)
}
