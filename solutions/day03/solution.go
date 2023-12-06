package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Ilyes512/advent-of-code-2023/advent/day03"
)

func main() {
	file, err := os.Open("./solutions/day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	schematic := day03.NewSchematic()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		schematic.AddRow(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	totalPart1, totalPart2 := schematic.Calculate()

	fmt.Printf("Part 1 total: %d\n", totalPart1)
	fmt.Printf("Part 2 total: %d\n", totalPart2)
}
