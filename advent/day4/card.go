package day4

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func NewCard(input string) *card {
	cardPart, numbersPart, ok := strings.Cut(input, ":")
	if !ok {
		log.Panic("Invalid input")
	}

	winningNumbersString, numbersString, ok := strings.Cut(numbersPart, "|")
	if !ok {
		log.Panic("Invalid input")
	}

	card := &card{
		CardNumber:     getCardId(cardPart),
		WinningNumbers: getSliceOfInts(winningNumbersString),
		Numbers:        getSliceOfInts(numbersString),
	}

	card.process()

	return card
}

type card struct {
	CardNumber     int
	WinningNumbers []int
	Numbers        []int
	Matches        int
	Points         int
}

func (c *card) getNumberOfCopies(matches map[int][]*card) int {
	current := matches[c.CardNumber]

	total := len(current)
	for _, card := range current {
		total += card.getNumberOfCopies(matches)
	}

	return total
}

func (c *card) process() {
	matches := 0
	for _, number := range c.WinningNumbers {
		if contains(c.Numbers, number) {
			matches++
		}
	}

	c.Matches = matches

	if matches == 0 {
		c.Points = 0
	} else {
		c.Points = int(math.Pow(2, float64(matches-1)))
	}
}

func getCardId(input string) int {
	var cardId int
	_, err := fmt.Sscanf(input, "Card %d", &cardId)
	if err != nil {
		log.Panic(err)
	}

	return cardId
}

func getSliceOfInts(input string) []int {
	words := strings.Fields(input)

	ints := make([]int, len(words))
	for i, word := range words {
		int, err := strconv.Atoi(word)
		if err != nil {
			log.Panic(err)
		}

		ints[i] = int
	}

	return ints
}

func contains(slice []int, number int) bool {
	for _, n := range slice {
		if n == number {
			return true
		}
	}

	return false
}
