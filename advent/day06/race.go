package day06

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type raceRecord struct {
	time     int
	distance int
}

func (r raceRecord) getDistanceForSpeed(speed int) int {
	return speed * (r.time - speed)
}

func (r raceRecord) GetNumberOfRecordBreakingRaces() int {
	races := make([]int, r.time+1)
	for i := range races {
		races[i] = i
	}

	ceiling := int(uint(r.time)) >> 1
	result := sort.Search(ceiling, func(i int) bool {
		return r.getDistanceForSpeed(i) > r.distance
	})

	recordBreakers := (ceiling + 1 - result) * 2

	if r.time%2 == 0 {
		return recordBreakers - 1
	}

	return recordBreakers
}

type records []raceRecord

func (r records) Result() int {
	result := 0

	for i := range r {
		if i == 0 {
			result = r[i].GetNumberOfRecordBreakingRaces()
			continue
		}

		result *= r[i].GetNumberOfRecordBreakingRaces()
	}

	return result
}

func NewRecordsPart1(filePath string) records {
	contents, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(contents), "\n")

	timeSlice := getSliceOfInts(removePrefix(lines[0]))
	distanceSlice := getSliceOfInts(removePrefix(lines[1]))

	records := make(records, len(timeSlice))

	for i := range records {
		records[i] = raceRecord{
			time:     timeSlice[i],
			distance: distanceSlice[i],
		}
	}

	return records
}

func NewRecordsPart2(filePath string) records {
	contents, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(contents), "\n")

	timeString := strings.ReplaceAll(removePrefix(lines[0]), " ", "")
	time, err := strconv.Atoi(timeString)
	if err != nil {
		log.Panic(err)
	}

	distanceString := strings.ReplaceAll(removePrefix(lines[1]), " ", "")
	distance, err := strconv.Atoi(distanceString)
	if err != nil {
		log.Panic(err)
	}

	records := make(records, 1)

	for i := range records {
		records[i] = raceRecord{
			time:     time,
			distance: distance,
		}
	}

	return records
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

func removePrefix(input string) string {
	_, after, found := strings.Cut(input, ":")
	if !found {
		log.Panic("Could not find ':' in input")
	}

	return after
}
