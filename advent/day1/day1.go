package day1

import (
	"fmt"
	"strconv"
	"strings"
)

var asciiReplaceOrder = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

type Document struct {
	lines []string
}

func NewDocument(input string) *Document {
	lines := strings.Split(input, "\n")

	document := Document{lines: lines}

	return document.trimSpaces()
}

func (d *Document) trimSpaces() *Document {
	for i, line := range d.lines {
		d.lines[i] = strings.TrimSpace(line)
	}
	return d
}

func (d *Document) Translate() *Document {
	for i, line := range d.lines {
		d.lines[i] = TranslateAsciiNumbers(line)
	}

	return d
}

func (d *Document) Calculate() (int, error) {
	total := 0

	for _, line := range d.lines {
		if line == "" {
			continue
		}

		result, err := GetNumberFromString(line)
		if err != nil {
			return 0, err
		}
		total += result
	}

	return total, nil
}

func GetNumberFromString(input string) (int, error) {
	start, end := 0, len(input)-1
	first, last := "", ""

	for end >= 0 {
		if first == "" && input[start] >= '0' && input[start] <= '9' {
			first = string(input[start])
		}

		if last == "" && input[end] >= '0' && input[end] <= '9' {
			last = string(input[end])
		}

		if first != "" && last != "" {
			break
		}

		start++
		end--
	}

	if first == "" || last == "" {
		return 0, fmt.Errorf("no first and/or last digits found in input string: %q", input)
	}

	result, err := strconv.Atoi(first + last)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func TranslateAsciiNumbers(input string) string {
	length := len(input)
	result := ""

	for i := 0; i < length; i++ {
		if input[i] >= '0' && input[i] <= '9' {
			result += string(input[i])
			continue
		}

		if asciiNumber, ok := GetAsciiNumberAtStart(input[i:]); ok {
			result += asciiNumber
		}
	}

	return result
}

func GetAsciiNumberAtStart(input string) (string, bool) {
	for i := range asciiReplaceOrder {
		index := strings.Index(input, asciiReplaceOrder[i])

		if index == 0 {
			return fmt.Sprint(i), true
		}
	}

	return "", false
}
