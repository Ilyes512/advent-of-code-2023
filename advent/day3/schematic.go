package day3

import (
	"fmt"
	"log"
	"strconv"
)

type gear map[coord]*partNumber

type partNumber struct {
	startCoord coord
	value      int
	counts     bool
}

type coord struct {
	X int
	Y int
}

func (c *coord) isValid() bool {
	return c.X >= 0 && c.Y >= 0
}

type schematic struct {
	symbols    []*cell
	gears      []*cell
	numbers    map[coord]*partNumber
	allNumbers []*partNumber
	rowNumber  int
}

func NewSchematic() *schematic {
	return &schematic{
		numbers: make(map[coord]*partNumber),
	}
}

func (s *schematic) AddRow(input string) {
	var pn *partNumber
	for i, c := range input {
		coord := coord{X: i, Y: s.rowNumber}
		cell := &cell{value: c, coord: coord}

		if !cell.isNumber() {
			pn = nil
		}

		if cell.isPeriod() {
			continue
		}

		if cell.isSymbol() {
			s.symbols = append(s.symbols, cell)
		}

		if cell.isNumber() {
			if pn == nil {
				pn = &partNumber{
					startCoord: coord,
					value:      int(c) - '0',
				}
				s.allNumbers = append(s.allNumbers, pn)
			} else {
				pn.value = appendNumber(pn.value, c)
			}

			s.numbers[coord] = pn
		}
	}
	s.rowNumber++
}

func (s *schematic) Calculate() (int, int) {
	var gears []gear

	for _, symbol := range s.symbols {
		g := make(gear)

		for _, coord := range symbol.getSurroundingCoords() {
			if number, ok := s.numbers[coord]; ok {
				number.counts = true

				if symbol.isGear() {
					g[number.startCoord] = number
				}
			}
		}

		if symbol.isGear() && len(g) == 2 {
			gears = append(gears, g)
		}
	}

	total := 0
	for _, number := range s.allNumbers {
		if number.counts {
			total += number.value
		}
	}

	gearTotals := 0
	for _, gear := range gears {
		gearTotal := 0
		for _, number := range gear {
			if gearTotal == 0 {
				gearTotal = number.value
				continue
			}
			gearTotal *= number.value
		}

		gearTotals += gearTotal
	}

	return total, gearTotals
}

func appendNumber(number int, c rune) int {
	value, err := strconv.Atoi(fmt.Sprintf("%d%c", number, c))
	if err != nil {
		log.Panic(err)
	}

	return value
}
