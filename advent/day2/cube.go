package day2

import (
	"fmt"
	"log"
)

type Color string

const (
	Green Color = "green"
	Red   Color = "red"
	Blue  Color = "blue"
)

type Cube struct {
	Amount int
	Color  Color
}

func (c *Cube) Update(input string) {
	var amount int
	var color string
	_, err := fmt.Sscanf(input, "%d %s", &amount, &color)
	if err != nil {
		log.Panic(err)
	}

	c.Amount = amount
	c.Color = Color(color)
}

func (c *Cube) isColor(color Color) bool {
	return c.Color == color
}
