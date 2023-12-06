package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Ilyes512/advent-of-code-2023/advent/day04"
)

func main() {
	file, err := os.Open("./solutions/day04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cards := day04.NewCards()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cards.AddCard(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1 total: %d\n", cards.GetTotalPoints())
	fmt.Printf("Part 2 total: %d\n", cards.GetNumberOfCardsAndCopies())
}
