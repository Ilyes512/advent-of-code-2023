package day05

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Reader struct {
	filePath string
	seeds    []int
	mappings []*mapping
}

func NewReader(filePath string) *Reader {
	return &Reader{filePath: filePath}
}

func (r *Reader) Process() int {
	r.read()

	lowestLocation := 0
	for i, seed := range r.seeds {
		location := r.getLocation(seed)

		if i == 0 {
			lowestLocation = location
		}

		if location < lowestLocation {
			lowestLocation = location
		}
	}

	return lowestLocation
}

func (r *Reader) getLocation(seed int) int {
	output := seed
	for _, m := range r.mappings {
		output = m.getMappingFor(output)
	}

	return output
}

func (r *Reader) read() {
	file, err := os.Open(r.filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		if i := strings.Index(string(data), "\n\n"); i >= 0 {
			return i + 2, data[0:i], nil
		}

		if atEOF {
			return len(data), data, nil
		}

		return 0, nil, nil
	}

	scanner.Split(split)
	for scanner.Scan() {
		data := scanner.Text()
		if strings.HasPrefix(data, "seeds") {
			r.processSeeds(data)
		} else {
			r.processMapping(data)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (r *Reader) processSeeds(input string) {
	seeds, found := strings.CutPrefix(input, "seeds:")
	if !found {
		log.Fatalf("Invalid seed input: %s", input)
	}

	r.seeds = getSliceOfInts(seeds)
}

func (r *Reader) processMapping(input string) {
	name, data, found := strings.Cut(input, ":")
	if !found {
		log.Fatalf("Invalid mapping input: %s", input)
	}

	name = strings.TrimSuffix(name, "map")
	source, dest, found := strings.Cut(name, "-to-")
	if !found {
		log.Fatalf("Invalid mapping input: %s", name)
	}

	mapping := mapping{
		source: source,
		dest:   dest,
	}

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}

		values := getSliceOfInts(text)
		if len(values) != 3 {
			log.Fatalf("Invalid number of values for mapping: %s\n", scanner.Text())
		}

		mapping.ranges = append(mapping.ranges, &mapRange{
			source: values[1],
			dest:   values[0],
			len:    values[2],
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	r.mappings = append(r.mappings, &mapping)
}

func getSliceOfInts(input string) []int {
	words := strings.Fields(input)

	ints := make([]int, len(words))
	for i, word := range words {
		int, err := strconv.Atoi(word)
		if err != nil {
			log.Panic(err)
		}

		ints[i] = int
	}

	return ints
}
