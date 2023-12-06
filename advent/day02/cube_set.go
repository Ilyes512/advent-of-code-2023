package day02

type CubeSet struct {
	Cubes []Cube
}

func (cs *CubeSet) IsPossible(set CubeSet) bool {
	for i := range set.Cubes {
		cube, ok := cs.getByColor(set.Cubes[i].Color)
		if !ok {
			continue
		}

		if cube.Amount > set.Cubes[i].Amount {
			return false
		}
	}

	return true
}

func (cs *CubeSet) getByColor(color Color) (*Cube, bool) {
	for i := range cs.Cubes {
		if cs.Cubes[i].isColor(color) {
			return &cs.Cubes[i], true
		}
	}

	return nil, false
}

func (cs *CubeSet) GetPower() int {
	power := 0
	for i := range cs.Cubes {
		if power == 0 {
			power = cs.Cubes[i].Amount
			continue
		}

		power *= cs.Cubes[i].Amount
	}

	return power
}
