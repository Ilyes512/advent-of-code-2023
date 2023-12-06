package day02

import (
	"fmt"
	"log"
	"strings"
)

type Game struct {
	Id   int
	Sets []CubeSet
}

func (g *Game) IsPossible(set CubeSet) bool {
	for i := range g.Sets {
		isPossible := g.Sets[i].IsPossible(set)
		if !isPossible {
			return false
		}
	}

	return true
}

func (g *Game) GetRequiredCubes() *CubeSet {
	cubeSet := CubeSet{
		Cubes: []Cube{
			{Amount: 0, Color: Red},
			{Amount: 0, Color: Green},
			{Amount: 0, Color: Blue},
		},
	}

	for i := range g.Sets {
		for j := range g.Sets[i].Cubes {
			cube := g.Sets[i].Cubes[j]

			minCube, ok := cubeSet.getByColor(cube.Color)
			if !ok {
				log.Panic("Color not found in cube set")
			}

			if minCube.Amount == 0 || minCube.Amount < cube.Amount {
				minCube.Amount = cube.Amount
			}
		}
	}

	return &cubeSet

}

func NewGame(input string) Game {
	result := strings.SplitN(input, ":", 2)
	if len(result) < 1 {
		log.Panic("Invalid input")
	}
	gameId := getGameId(result[0])

	cubeSets := strings.Split(result[1], ";")
	sets := make([]CubeSet, len(cubeSets))

	for i, cubesString := range cubeSets {
		sets[i] = *getCubeSet(cubesString)
	}

	return Game{
		Id:   gameId,
		Sets: sets,
	}
}

func getGameId(input string) int {
	var gameId int
	_, err := fmt.Sscanf(input, "Game %d", &gameId)
	if err != nil {
		log.Panic(err)
	}

	return gameId
}

func getCubeSet(cubesString string) *CubeSet {
	cubeStrings := strings.Split(cubesString, ",")
	cubes := make([]Cube, len(cubeStrings))

	for i := range cubes {
		cubes[i].Update(cubeStrings[i])
	}

	return &CubeSet{
		Cubes: cubes,
	}
}
