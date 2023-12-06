package day02

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCubeSetIsPossible(t *testing.T) {
	availableCubes := CubeSet{
		Cubes: []Cube{
			{Amount: 5, Color: Red},
			{Amount: 3, Color: Green},
			{Amount: 4, Color: Blue},
		},
	}
	tests := []struct {
		cubeSet CubeSet
		want    bool
	}{
		{
			cubeSet: CubeSet{
				Cubes: []Cube{
					{Amount: 3, Color: Red},
				},
			},
			want: true,
		},
		{
			cubeSet: CubeSet{
				Cubes: []Cube{
					{Amount: 5, Color: Red},
				},
			},
			want: true,
		},
		{
			cubeSet: CubeSet{
				Cubes: []Cube{
					{Amount: 4, Color: Red},
					{Amount: 4, Color: Blue},
				},
			},
			want: true,
		},
		{
			cubeSet: CubeSet{
				Cubes: []Cube{
					{Amount: 3, Color: Red},
					{Amount: 5, Color: Blue},
				},
			},
			want: false,
		},
		{
			cubeSet: CubeSet{
				Cubes: []Cube{
					{Amount: 5, Color: Red},
					{Amount: 3, Color: Green},
					{Amount: 4, Color: Blue},
				},
			},
			want: true,
		},
	}

	for i, test := range tests {
		test := test

		t.Run(fmt.Sprintf("case %d", i+1), func(tt *testing.T) {
			tt.Parallel()

			got := test.cubeSet.IsPossible(availableCubes)

			if diff := cmp.Diff(test.want, got); diff != "" {
				tt.Errorf("With available cubes %+v calling IsPossible(%q) mismatch (-want +got):\n%s", availableCubes, test.cubeSet, diff)
			}
		})
	}
}

func TestCubeSetGetPower(t *testing.T) {
	tests := []struct {
		cubeSet CubeSet
		want    int
	}{
		{
			cubeSet: CubeSet{
				Cubes: []Cube{
					{Amount: 3, Color: Red},
				},
			},
			want: 3,
		},
		{
			cubeSet: CubeSet{
				Cubes: []Cube{
					{Amount: 5, Color: Red},
				},
			},
			want: 5,
		},
		{
			cubeSet: CubeSet{
				Cubes: []Cube{
					{Amount: 4, Color: Red},
					{Amount: 4, Color: Blue},
				},
			},
			want: 16,
		},
		{
			cubeSet: CubeSet{
				Cubes: []Cube{
					{Amount: 3, Color: Red},
					{Amount: 5, Color: Blue},
				},
			},
			want: 15,
		},
		{
			cubeSet: CubeSet{
				Cubes: []Cube{
					{Amount: 5, Color: Red},
					{Amount: 3, Color: Green},
					{Amount: 4, Color: Blue},
				},
			},
			want: 60,
		},
	}

	for i, test := range tests {
		test := test

		t.Run(fmt.Sprintf("case %d", i+1), func(tt *testing.T) {
			tt.Parallel()

			got := test.cubeSet.GetPower()

			if diff := cmp.Diff(test.want, got); diff != "" {
				tt.Errorf("CubeSet{%+v}.GetPower() mismatch (-want +got):\n%s", test.cubeSet, diff)
			}
		})
	}
}
