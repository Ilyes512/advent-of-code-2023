package day05

import "slices"

type mapTable struct {
	source string
	dest   string
	maps   []*mapping
}

func (mt *mapTable) getMappingFor(source int) int {
	for _, m := range mt.maps {
		if m.inRange(source) {
			return source - m.diff()
		}
	}

	return source
}

type mapping struct {
	source int
	dest   int
	len    int
}

func (m *mapping) inRange(source int) bool {
	return source >= m.source && source < m.source+m.len
}

func (m *mapping) contains(seedRange seedRange) bool {
	return seedRange.min >= m.sourceMin() && seedRange.max <= m.sourceMax()
}

func (m *mapping) containsOrDoesNotOverlap(seedRange seedRange) bool {
	return m.contains(seedRange) || m.doesNotOverlap(seedRange)
}

func (m *mapping) hasOverlap(seedRange seedRange) bool {
	return m.sourceMin() <= seedRange.max && m.sourceMax() >= seedRange.min
}

func (m *mapping) doesNotOverlap(seedRange seedRange) bool {
	return !m.hasOverlap(seedRange)
}

func (m *mapping) diff() int {
	return m.source - m.dest
}

func (m *mapping) sourceMin() int {
	return m.source
}

func (m *mapping) sourceMax() int {
	return m.source + m.len - 1
}

type seedRanges []seedRange

func (s seedRanges) transform(mapTable *mapTable) seedRanges {
	var result seedRanges

	for _, iSeedRange := range s {
		result = append(result, iSeedRange.transform(mapTable)...)
	}

	return result
}

type seedRange struct {
	min int // inclusive
	max int // inclusive
}

func (s seedRange) transform(mapTable *mapTable) seedRanges {
	result := seedRanges{s}

LOOP:
	for i, iSeedRange := range result {
		for _, mapping := range mapTable.maps {
			if len(result) == 1 && mapping.contains(iSeedRange) {
				// there is only one seedRange and it is contained in a mapping
				goto DONE
			}

			if mapping.containsOrDoesNotOverlap(iSeedRange) {
				continue
			}

			if iSeedRange.min < mapping.sourceMin() {
				result = append(result, seedRange{
					min: iSeedRange.min,
					max: mapping.sourceMin() - 1,
				})
				result = slices.Replace(result, i, i+1, seedRange{
					min: mapping.sourceMin(),
					max: iSeedRange.max,
				})

				// restart the loop, because the new seedRange could overlap with another mapping
				goto LOOP
			}

			if iSeedRange.max > mapping.sourceMax() {
				result = append(result, seedRange{
					min: mapping.sourceMax() + 1,
					max: iSeedRange.max,
				})
				result = slices.Replace(result, i, i+1, seedRange{
					min: iSeedRange.min,
					max: mapping.sourceMax(),
				})

				// restart the loop, because the new seedRange could overlap with another mapping
				goto LOOP
			}
		}
	}

DONE:
	for i, iSeedRange := range result {
		for _, mapping := range mapTable.maps {
			if mapping.contains(iSeedRange) {
				result = slices.Replace(result, i, i+1, seedRange{
					min: iSeedRange.min - mapping.diff(),
					max: iSeedRange.max - mapping.diff(),
				})
				break
			}
		}
	}

	return result
}
