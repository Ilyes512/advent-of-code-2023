package day3

type cell struct {
	value rune
	coord coord
}

func (c *cell) isGear() bool {
	return c.value == '*'
}

func (c *cell) isPeriod() bool {
	return c.value == '.'
}

func (c *cell) isNumber() bool {
	return c.value >= '0' && c.value <= '9'
}

func (c *cell) isSymbol() bool {
	return !c.isPeriod() && !c.isNumber()
}

func (c *cell) getSurroundingCoords() []coord {
	X := c.coord.X
	Y := c.coord.Y

	coords := []coord{
		{X: X - 1, Y: Y - 1},
		{X: X, Y: Y - 1},
		{X: X + 1, Y: Y - 1},

		{X: X - 1, Y: Y},
		{X: X + 1, Y: Y},

		{X: X - 1, Y: Y + 1},
		{X: X, Y: Y + 1},
		{X: X + 1, Y: Y + 1},
	}

	result := make([]coord, 0, len(coords))
	for _, coord := range coords {
		if coord.isValid() {
			result = append(result, coord)
		}
	}

	return result
}
