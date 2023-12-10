package day07

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDetermineHandType(t *testing.T) {
	tests := []struct {
		input []int
		want  handType
	}{
		{
			// QQQQQ
			input: []int{11, 11, 11, 11, 11},
			want:  FiveOfAKind,
		},
		{
			// 99992
			input: []int{8, 8, 8, 8, 1},
			want:  FourOfAKind,
		},
		{
			// 33322
			input: []int{2, 2, 2, 1, 1},
			want:  FullHouse,
		},
		{
			// T55J5
			input: []int{9, 4, 4, 10, 4},
			want:  ThreeOfAKind,
		},
		{
			// KK677
			input: []int{12, 12, 5, 6, 6},
			want:  TwoPairs,
		},
		{
			// 32T3K
			input: []int{2, 1, 9, 2, 12},
			want:  OnePair,
		},
		{
			// 9TJQK
			input: []int{9, 10, 11, 12, 13},
			want:  HighCard,
		},
	}

	for i, test := range tests {
		test := test

		t.Run(fmt.Sprintf("case %d", i+1), func(tt *testing.T) {
			tt.Parallel()

			got := determineHandType(test.input)

			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("determineHandType() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
