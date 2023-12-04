package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Ilyes512/advent-of-code-2023/advent/day2"
)

func main() {
	file, err := os.Open("./solutions/day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0
	input := day2.CubeSet{
		Cubes: []day2.Cube{
			{Amount: 12, Color: day2.Red},
			{Amount: 13, Color: day2.Green},
			{Amount: 14, Color: day2.Blue},
		},
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		game := day2.NewGame(scanner.Text())
		if game.IsPossible(input) {
			total += game.Id
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1 total: %d\n", total)
}
