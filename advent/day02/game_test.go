package day02

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewGame(t *testing.T) {
	tests := []struct {
		input string
		want  Game
	}{
		{
			input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want: Game{
				Id: 1,
				Sets: []CubeSet{
					{
						Cubes: []Cube{
							{Amount: 3, Color: Blue},
							{Amount: 4, Color: Red},
						},
					},
					{
						Cubes: []Cube{
							{Amount: 1, Color: Red},
							{Amount: 2, Color: Green},
							{Amount: 6, Color: Blue},
						},
					},
					{
						Cubes: []Cube{
							{Amount: 2, Color: Green},
						},
					},
				},
			},
		},
		{
			input: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			want: Game{
				Id: 2,
				Sets: []CubeSet{
					{
						Cubes: []Cube{
							{Amount: 1, Color: Blue},
							{Amount: 2, Color: Green},
						},
					},
					{
						Cubes: []Cube{
							{Amount: 3, Color: Green},
							{Amount: 4, Color: Blue},
							{Amount: 1, Color: Red},
						},
					},
					{
						Cubes: []Cube{
							{Amount: 1, Color: Green},
							{Amount: 1, Color: Blue},
						},
					},
				},
			},
		},
	}

	for i, test := range tests {
		test := test
		t.Run(fmt.Sprintf("case %d", i+1), func(tt *testing.T) {
			tt.Parallel()

			got := NewGame(test.input)

			if diff := cmp.Diff(test.want, got); diff != "" {
				tt.Errorf("NewGame(%q) mismatch (-want +got):\n%s", test.input, diff)
			}
		})
	}
}

func TestGameIsPossible(t *testing.T) {
	isPossibleInput := CubeSet{
		Cubes: []Cube{
			{Amount: 12, Color: Red},
			{Amount: 13, Color: Green},
			{Amount: 14, Color: Blue},
		},
	}
	tests := []struct {
		gameInput string
		want      bool
	}{
		{
			gameInput: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want:      true,
		},
		{
			gameInput: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			want:      true,
		},
		{
			gameInput: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			want:      false,
		},
		{
			gameInput: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			want:      false,
		},
		{
			gameInput: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			want:      true,
		},
	}

	for i, test := range tests {
		test := test

		t.Run(fmt.Sprintf("case %d", i+1), func(tt *testing.T) {
			tt.Parallel()

			game := NewGame(test.gameInput)

			got := game.IsPossible(isPossibleInput)

			if diff := cmp.Diff(test.want, got); diff != "" {
				tt.Errorf("NewGame(%q).IsPossible(%q) mismatch (-want +got):\n%s", test.gameInput, isPossibleInput, diff)
			}
		})
	}
}

func TestGameGetRequiredCubes(t *testing.T) {
	tests := []struct {
		gameInput string
		want      CubeSet
	}{
		{
			gameInput: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want: CubeSet{
				Cubes: []Cube{
					{Amount: 4, Color: Red},
					{Amount: 2, Color: Green},
					{Amount: 6, Color: Blue},
				},
			},
		},
		{
			gameInput: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			want: CubeSet{
				Cubes: []Cube{
					{Amount: 1, Color: Red},
					{Amount: 3, Color: Green},
					{Amount: 4, Color: Blue},
				},
			},
		},
		{
			gameInput: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			want: CubeSet{
				Cubes: []Cube{
					{Amount: 20, Color: Red},
					{Amount: 13, Color: Green},
					{Amount: 6, Color: Blue},
				},
			},
		},
		{
			gameInput: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			want: CubeSet{
				Cubes: []Cube{
					{Amount: 14, Color: Red},
					{Amount: 3, Color: Green},
					{Amount: 15, Color: Blue},
				},
			},
		},
		{
			gameInput: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			want: CubeSet{
				Cubes: []Cube{
					{Amount: 6, Color: Red},
					{Amount: 3, Color: Green},
					{Amount: 2, Color: Blue},
				},
			},
		},
	}

	for i, test := range tests {
		test := test

		t.Run(fmt.Sprintf("case %d", i+1), func(tt *testing.T) {
			tt.Parallel()

			game := NewGame(test.gameInput)

			got := game.GetRequiredCubes()

			if diff := cmp.Diff(test.want, *got); diff != "" {
				tt.Errorf("NewGame(%q).GetRequiredCubes() mismatch (-want +got):\n%s", test.gameInput, diff)
			}
		})
	}
}
