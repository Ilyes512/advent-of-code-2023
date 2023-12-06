package day4

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewCard(t *testing.T) {
	tests := []struct {
		input string
		want  *card
	}{
		{
			input: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			want: &card{
				CardNumber:     1,
				WinningNumbers: []int{41, 48, 83, 86, 17},
				Numbers:        []int{83, 86, 6, 31, 17, 9, 48, 53},
				Matches:        4,
				Points:         8,
			},
		},
		{
			input: "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
			want: &card{
				CardNumber:     2,
				WinningNumbers: []int{13, 32, 20, 16, 61},
				Numbers:        []int{61, 30, 68, 82, 17, 32, 24, 19},
				Matches:        2,
				Points:         2,
			},
		},
	}

	for i, test := range tests {
		test := test

		t.Run(fmt.Sprintf("case %d", i+1), func(tt *testing.T) {
			tt.Parallel()

			got := NewCard(test.input)

			if diff := cmp.Diff(test.want, got); diff != "" {
				tt.Errorf("NewCard(%s) mismatch (-want +got):\n%s", test.input, diff)
			}
		})
	}
}
