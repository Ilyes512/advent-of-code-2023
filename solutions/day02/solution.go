package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Ilyes512/advent-of-code-2023/advent/day02"
)

func main() {
	file, err := os.Open("./solutions/day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	totalPart1 := 0
	totalPart2 := 0
	input := day02.CubeSet{
		Cubes: []day02.Cube{
			{Amount: 12, Color: day02.Red},
			{Amount: 13, Color: day02.Green},
			{Amount: 14, Color: day02.Blue},
		},
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		game := day02.NewGame(scanner.Text())
		if game.IsPossible(input) {
			totalPart1 += game.Id
		}

		totalPart2 += game.GetRequiredCubes().GetPower()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1 total: %d\n", totalPart1)
	fmt.Printf("Part 2 total: %d\n", totalPart2)
}
