package day07

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDetermineHandType(t *testing.T) {
	tests := []struct {
		label string
		input []int
		want  handType
	}{
		{
			label: "QQQQQ",
			input: []int{11, 11, 11, 11, 11},
			want:  FiveOfAKind,
		},
		{
			label: "99992",
			input: []int{8, 8, 8, 8, 1},
			want:  FourOfAKind,
		},
		{
			label: "33322",
			input: []int{2, 2, 2, 1, 1},
			want:  FullHouse,
		},
		{
			label: "T55J5",
			input: []int{9, 4, 4, 10, 4},
			want:  ThreeOfAKind,
		},
		{
			label: "KK677",
			input: []int{12, 12, 5, 6, 6},
			want:  TwoPairs,
		},
		{
			label: "32T3K",
			input: []int{2, 1, 9, 2, 12},
			want:  OnePair,
		},
		{
			label: "9TJQK",
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
				t.Errorf("determineHandType(%v) mismatch (-want +got):\n%s", test.label, diff)
			}
		})
	}
}

func TestDetermineHandTypeWithJoker(t *testing.T) {
	tests := []struct {
		label string
		input []int
		want  handType
	}{
		{
			label: "QQQQQ",
			input: []int{11, 11, 11, 11, 11},
			want:  FiveOfAKind,
		},
		{
			label: "99992",
			input: []int{8, 8, 8, 8, 1},
			want:  FourOfAKind,
		},
		{
			label: "33322",
			input: []int{2, 2, 2, 1, 1},
			want:  FullHouse,
		},
		{
			label: "T55J5",
			input: []int{9, 4, 4, 10, 4},
			want:  FourOfAKind,
		},
		{
			label: "KK677",
			input: []int{12, 12, 5, 6, 6},
			want:  TwoPairs,
		},
		{
			label: "32T3K",
			input: []int{2, 1, 9, 2, 12},
			want:  OnePair,
		},
		{
			label: "9TJQK",
			input: []int{9, 10, 11, 12, 13},
			want:  OnePair,
		},

		// 5 jokers
		{
			label: "JJJJJ",
			input: []int{10, 10, 10, 10, 10},
			want:  FiveOfAKind,
		},
		// 4 jokers
		{
			label: "JJJJ5",
			input: []int{10, 10, 10, 10, 4},
			want:  FiveOfAKind,
		},
		// 3 jokers
		{
			label: "JJJ66",
			input: []int{10, 10, 10, 5, 5},
			want:  FiveOfAKind,
		},
		{
			label: "JJJ65",
			input: []int{10, 10, 10, 5, 4},
			want:  FourOfAKind,
		},
		// 2 jokers
		{
			label: "JJ666",
			input: []int{10, 10, 5, 5, 5},
			want:  FiveOfAKind,
		},
		{
			label: "JJ665",
			input: []int{10, 10, 5, 5, 4},
			want:  FourOfAKind,
		},
		{
			label: "JJ654",
			input: []int{10, 10, 5, 4, 3},
			want:  ThreeOfAKind,
		},
		// 1 joker
		{
			label: "J6666",
			input: []int{10, 5, 5, 5, 5},
			want:  FiveOfAKind,
		},
		{
			label: "J6665",
			input: []int{10, 5, 5, 5, 4},
			want:  FourOfAKind,
		},
		{
			label: "J6654",
			input: []int{10, 5, 5, 4, 3},
			want:  ThreeOfAKind,
		},
		{
			label: "J6543",
			input: []int{10, 5, 4, 3, 2},
			want:  OnePair,
		},
		{
			label: "J6655",
			input: []int{10, 5, 5, 4, 4},
			want:  FullHouse,
		},
	}

	for i, test := range tests {
		test := test

		t.Run(fmt.Sprintf("case %d", i+1), func(tt *testing.T) {
			tt.Parallel()

			got := determineHandTypeWithJoker(test.input)

			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("determineHandTypeWithJoker(%s) mismatch (-want +got):\n%s", test.label, diff)
			}
		})
	}
}
