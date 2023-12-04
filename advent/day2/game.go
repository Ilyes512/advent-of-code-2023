package day2

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
