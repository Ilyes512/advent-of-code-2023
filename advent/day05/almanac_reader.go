package day05

import (
	"bufio"
	"cmp"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Reader struct {
	loaded     bool
	filePath   string
	seeds      []int
	seedRanges seedRanges
	mapTables  []mapTable
}

func NewReader(filePath string) *Reader {
	return &Reader{filePath: filePath}
}

func (r *Reader) ProcessPart1() int {
	r.read()

	var results []int

	for _, seed := range r.seeds {
		results = append(results, r.getLocation(seed))
	}

	return slices.Min(results)
}

func (r *Reader) ProcessPart2() int {
	r.read()

	results := slices.Clone(r.seedRanges)
	for _, mapTable := range r.mapTables {
		results = results.transform(&mapTable)
	}

	return slices.MinFunc(results, func(a, b seedRange) int {
		return cmp.Compare(a.min, b.min)
	}).min
}

func (r *Reader) getLocation(seed int) int {
	output := seed
	for _, mapTable := range r.mapTables {
		output = mapTable.getMappingFor(output)
	}

	return output
}

func (r *Reader) read() {
	if r.loaded {
		return
	}

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
			r.processMapTables(data)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	r.loaded = true
}

func (r *Reader) processSeeds(input string) {
	seeds, found := strings.CutPrefix(input, "seeds:")
	if !found {
		log.Fatalf("Invalid seed input: %s", input)
	}

	r.seeds = getSliceOfInts(seeds)

	if len(r.seeds)%2 != 0 {
		log.Fatalf("Invalid seed input (number of seeds should be divisible by 2): %s", input)
	}

	for i := 0; i < len(r.seeds); i += 2 {
		r.seedRanges = append(r.seedRanges, seedRange{min: r.seeds[i], max: r.seeds[i] + r.seeds[i+1]})
	}
}

func (r *Reader) processMapTables(input string) {
	name, data, found := strings.Cut(input, ":")
	if !found {
		log.Fatalf("Invalid mapping input: %s", input)
	}

	name = strings.TrimSuffix(name, "map")
	source, dest, found := strings.Cut(name, "-to-")
	if !found {
		log.Fatalf("Invalid mapping input: %s", name)
	}

	mapTable := mapTable{
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

		mapTable.maps = append(mapTable.maps, &mapping{
			source: values[1],
			dest:   values[0],
			len:    values[2],
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	r.mapTables = append(r.mapTables, mapTable)
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
