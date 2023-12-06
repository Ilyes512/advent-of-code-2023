package day4

func NewCards() *cards {
	return &cards{
		cards:   make(map[int]*card),
		matches: make(map[int][]*card),
	}
}

type cards struct {
	processed              bool
	order                  []int
	cards                  map[int]*card
	matches                map[int][]*card
	totalPoints            int
	numberOfCardsAndCopies int
}

func (c *cards) AddCard(line string) {
	card := NewCard(line)

	c.order = append(c.order, card.CardNumber)
	c.cards[card.CardNumber] = card
}

func (c *cards) Process() {
	numberOfOrginalCards := len(c.order)
	for _, i := range c.order {
		card := c.cards[i]
		c.totalPoints += card.Points

		for match := card.CardNumber + 1; match <= card.CardNumber+card.Matches && match <= numberOfOrginalCards; match++ {
			c.matches[card.CardNumber] = append(c.matches[card.CardNumber], c.cards[match])
		}
	}

	c.numberOfCardsAndCopies = numberOfOrginalCards
	for _, card := range c.cards {
		c.numberOfCardsAndCopies += card.getNumberOfCopies(c.matches)
	}

	c.processed = true
}

func (c *cards) GetTotalPoints() int {
	if !c.processed {
		c.Process()
	}

	return c.totalPoints
}

func (c *cards) GetNumberOfCardsAndCopies() int {
	if !c.processed {
		c.Process()
	}

	return c.numberOfCardsAndCopies
}
