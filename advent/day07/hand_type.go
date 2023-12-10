package day07

type handType int

func (h handType) String() string {
	switch h {
	case FiveOfAKind:
		return "Five of a kind"
	case FourOfAKind:
		return "Four of a kind"
	case FullHouse:
		return "Full house"
	case ThreeOfAKind:
		return "Three of a kind"
	case TwoPairs:
		return "Two pairs"
	case OnePair:
		return "One pair"
	case HighCard:
		return "High card"
	default:
		return "Unknown"
	}
}

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
	grouped := make(map[int]int)

	for _, card := range input {
		c, found := grouped[card]
		if found {
			grouped[card] = c + 1
		} else {
			grouped[card] = 1
		}
	}

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

func mapContainsValue(m map[int]int, value int) bool {
	for _, v := range m {
		if v == value {
			return true
		}
	}

	return false
}
