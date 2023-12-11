package day07

type handType int

const (
	HighCard handType = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func determineHandType(input []int) handType {
	grouped := groupedCards(input)

	switch len(grouped) {
	case 1:
		return FiveOfAKind
	case 2:
		if mapContainsValue(grouped, 4) {
			return FourOfAKind
		} else {
			return FullHouse
		}
	case 3:
		if mapContainsValue(grouped, 3) {
			return ThreeOfAKind
		} else {
			return TwoPairs
		}
	case 4:
		return OnePair
	default:
		return HighCard
	}
}

func determineHandTypeWithJoker(input []int) handType {
	grouped := groupedCards(input)

	joker, containsJoker := grouped[joker_value]
	if !containsJoker {
		return determineHandType(input)
	}
	delete(grouped, joker_value)

	switch joker {
	case 5, 4:
		return FiveOfAKind
	case 3:
		if len(grouped) == 1 {
			return FiveOfAKind
		} else {
			return FourOfAKind
		}
	case 2:
		switch len(grouped) {
		case 1:
			return FiveOfAKind
		case 2:
			return FourOfAKind
		default:
			return ThreeOfAKind
		}
	default:
		switch len(grouped) {
		case 1:
			return FiveOfAKind
		case 2:
			if mapContainsValue(grouped, 3) {
				return FourOfAKind
			} else {
				return FullHouse
			}
		case 3:
			return ThreeOfAKind
		default:
			return OnePair
		}
	}
}

func mapContainsValue(m map[int]int, value int) bool {
	for _, v := range m {
		if v == value {
			return true
		}
	}

	return false
}

func groupedCards(input []int) map[int]int {
	grouped := make(map[int]int)

	for _, card := range input {
		c, found := grouped[card]
		if found {
			grouped[card] = c + 1
		} else {
			grouped[card] = 1
		}
	}

	return grouped
}
