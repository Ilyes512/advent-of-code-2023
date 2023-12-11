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

type hands []hand

func NewHands(filePath string) (hands, hands) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	hands := make([]hand, 0)
	handsWithJoker := make([]hand, 0)

	cardMapping := NewCardMapping()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Fields(line)
		if len(split) != 2 {
			log.Panicf("Invalid input: %s", line)
		}

		cards := cardMapping.GetCardValues(split[0])

		hands = append(hands, NewHand(split[1], split[0], cards))
		handsWithJoker = append(handsWithJoker, NewHandWithJoker(split[1], split[0], cards))
	}

	hands = sortHands(hands)
	handsWithJoker = sortHands(handsWithJoker)

	return hands, handsWithJoker
}

type hand struct {
	cardsString   string
	bid           int
	originalCards []int
	cards         []int
	strength      int
	handType      handType
}

func NewHand(bid string, cardsString string, cards []int) hand {
	bidValue, err := strconv.Atoi(bid)
	if err != nil {
		log.Panic(err)
	}

	return hand{
		cardsString:   cardsString,
		bid:           bidValue,
		originalCards: cards,
		cards:         cards,
		strength:      calculateStrength(cards),
		handType:      determineHandType(cards),
	}
}

func NewHandWithJoker(bid string, cardsString string, cards []int) hand {
	bidValue, err := strconv.Atoi(bid)
	if err != nil {
		log.Panic(err)
	}

	cardsWithJoker := replaceValues(cards, joker_value, 0)

	return hand{
		cardsString:   cardsString,
		bid:           bidValue,
		originalCards: cards,
		cards:         cardsWithJoker,
		strength:      calculateStrength(cardsWithJoker),
		handType:      determineHandTypeWithJoker(cards),
	}
}

func sortHands(input hands) hands {
	slices.SortFunc(input, func(a, b hand) int {
		return cmp.Compare(a.strength, b.strength)
	})

	slices.SortStableFunc(input, func(a, b hand) int {
		result := cmp.Compare(a.handType, b.handType)

		if result == 0 {
			return slices.Compare(a.cards, b.cards)
		}

		return result
	})

	return input
}

func (h hands) Result() int {
	total := 0

	for i, hand := range h {
		total += hand.bid * (i + 1)
	}

	return total
}

func calculateStrength(input []int) int {
	sum := 0
	for _, v := range input {
		sum += v
	}

	return sum
}

func replaceValues(s []int, old, new int) []int {
	copy := slices.Clone(s)
	for i, v := range copy {
		if v == old {
			copy[i] = new
		}
	}

	return copy
}
