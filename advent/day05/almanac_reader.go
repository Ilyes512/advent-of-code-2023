package day05

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type Reader struct {
	loaded    bool
	filePath  string
	seeds     []int
	seedPairs [][]int
	mappings  []*mapping
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

func (r *Reader) ProcessPart2() int {
	r.read()

	consumer := 20
	producerChan := make(chan int, 1_000_000)
	lowestChan := make(chan int, consumer)

	go func(ch chan<- int) {
		for _, seedPair := range r.seedPairs {
			for j := seedPair[0]; j < seedPair[1]; j++ {
				ch <- j
			}
		}

		defer close(producerChan)
	}(producerChan)

	wg := sync.WaitGroup{}

	for i := 0; i < consumer; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			first := true
			min := 0
			for seed := range producerChan {
				location := r.getLocation(seed)

				if first == true {
					min = location
					first = false
				}

				if location < min {
					min = location
				}
			}
			if first == false {
				lowestChan <- min
			}
		}()
	}
	wg.Wait()
	close(lowestChan)

	result := 0
	for i := 0; i < consumer; i++ {
		min, ok := <-lowestChan
		if !ok {
			break
		}

		if i == 0 {
			result = min
		}

		if min < result {
			result = min
		}
	}

	return result
}

func (r *Reader) getLocation(seed int) int {
	output := seed
	for _, m := range r.mappings {
		output = m.getMappingFor(output)
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
			r.processMapping(data)
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
		r.seedPairs = append(r.seedPairs, []int{r.seeds[i], r.seeds[i] + r.seeds[i+1]})
	}

	// r.seedPairs = removeOverlappingRanges(r.seedPairs)
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

func removeOverlappingRanges(ranges [][]int) [][]int {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

LOOP:
	end := ranges[0][1]
	for i := 1; i <= len(ranges)-1; i++ {
		if ranges[i][0] < end && ranges[i][1] < end {
			// removeOverlappingRange
			ranges = append(ranges[:i], ranges[i+1:]...)
			goto LOOP
		}

		if ranges[i][0] < end && ranges[i][1] > end {
			// left is inside previous range but end is outside
			ranges[i-1][1] = ranges[i][1]
			ranges = append(ranges[:i], ranges[i+1:]...)
			goto LOOP
		}
	}

	return ranges
}
