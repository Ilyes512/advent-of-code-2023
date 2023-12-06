package day05

type mapping struct {
	source string
	dest   string
	ranges []*mapRange
}

func (m *mapping) getMappingFor(source int) int {
	for _, r := range m.ranges {
		if r.inRange(source) {
			return source - r.diff()
		}
	}

	return source
}

type mapRange struct {
	source int
	dest   int
	len    int
}

func (r *mapRange) inRange(source int) bool {
	return source >= r.source && source <= r.source+r.len
}

func (r *mapRange) diff() int {
	return r.source - r.dest
}
