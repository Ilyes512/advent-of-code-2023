package day03

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCellGetSurroundingCoords(t *testing.T) {
	tests := []struct {
		input cell
		want  []coord
	}{
		{
			input: cell{coord: coord{X: 0, Y: 0}},
			want: []coord{
				{X: 1, Y: 0},
				{X: 0, Y: 1},
				{X: 1, Y: 1},
			},
		},
		{
			input: cell{coord: coord{X: 2, Y: 4}},
			want: []coord{
				{X: 1, Y: 3},
				{X: 2, Y: 3},
				{X: 3, Y: 3},

				{X: 1, Y: 4},
				{X: 3, Y: 4},

				{X: 1, Y: 5},
				{X: 2, Y: 5},
				{X: 3, Y: 5},
			},
		},
		{
			input: cell{coord: coord{X: 4, Y: 0}},
			want: []coord{
				{X: 3, Y: 0},
				{X: 5, Y: 0},

				{X: 3, Y: 1},
				{X: 4, Y: 1},
				{X: 5, Y: 1},
			},
		},
	}

	for i, test := range tests {
		test := test

		t.Run(fmt.Sprintf("case %d", i+1), func(tt *testing.T) {
			tt.Parallel()

			got := test.input.getSurroundingCoords()
			sortCoords := func(c1, c2 coord) bool {
				if c1.X == c2.X {
					return c1.Y < c2.Y
				}

				return c1.X < c2.X
			}

			if diff := cmp.Diff(test.want, got, cmpopts.SortSlices(sortCoords)); diff != "" {
				tt.Errorf("Cell.GetSurroundingCoords() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
