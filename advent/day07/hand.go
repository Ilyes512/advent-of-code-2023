package day07

import (
	"bufio"
	"cmp"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type hands []*hand

func NewHands(filePath string) hands {
	file, err := os.Open(filePath)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	result := make([]*hand, 0)
	cardMapping := NewCardMapping()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Fields(line)
		if len(split) != 2 {
			log.Panicf("Invalid input: %s", line)
		}

		result = append(result, NewHand(split[1], cardMapping.GetCardValues(split[0])))
	}

	result = sortHands(result)

	return result
}

type hand struct {
	bid      int
	cards    []int
	strength int
	handType handType
}

func NewHand(bid string, cards []int) *hand {
	bidValue, err := strconv.Atoi(bid)
	if err != nil {
		log.Panic(err)
	}

	return &hand{
		bid:      bidValue,
		cards:    cards,
		strength: sumInt(cards),
		handType: determineHandType(cards),
	}
}

func sortHands(input hands) hands {
	slices.SortFunc(input, func(a, b *hand) int {
		return cmp.Compare(a.strength, b.strength)
	})

	slices.SortStableFunc(input, func(a, b *hand) int {
		result := cmp.Compare(a.handType, b.handType)

		if result == 0 {
			return slices.Compare(a.cards, b.cards)
		}

		return result
	})

	return input
}

func (h hands) ResultPart1() int {
	total := 0

	for i, hand := range h {
		total += hand.bid * (i + 1)
	}

	return total
}

func sumInt(input []int) int {
	sum := 0
	for _, i := range input {
		sum += i
	}

	return sum
}
