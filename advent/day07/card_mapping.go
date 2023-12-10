package day07

import "log"

type cardMapping map[string]int

func NewCardMapping() cardMapping {
	return cardMapping{
		"A": 13,
		"K": 12,
		"Q": 11,
		"J": 10,
		"T": 9,
		"9": 8,
		"8": 7,
		"7": 6,
		"6": 5,
		"5": 4,
		"4": 3,
		"3": 2,
		"2": 1,
	}
}

func (cm cardMapping) GetCardValue(card string) int {
	value, found := cm[card]
	if !found {
		log.Panicf("Card %s not found in mapping", card)
	}

	return value
}

func (cm cardMapping) GetCardValues(cards string) []int {
	values := make([]int, len(cards))
	for i, card := range cards {
		values[i] = cm.GetCardValue(string(card))
	}

	return values
}
