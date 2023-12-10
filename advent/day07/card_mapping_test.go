package day07

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewCardMapping(t *testing.T) {
	got := NewCardMapping()
	want := cardMapping{
		"A": 13,
		"K": 12,
		"Q": 11,
		"J": 10,
		"T": 9,
		"9": 8,
		"8": 7,
		"7": 6,
		"6": 5,
		"5": 4,
		"4": 3,
		"3": 2,
		"2": 1,
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("NewCardMapping() mismatch (-want +got):\n%s", diff)
	}
}

func TestGetCardValue(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"A", 13},
		{"K", 12},
		{"Q", 11},
		{"J", 10},
		{"T", 9},
		{"9", 8},
	}

	for i, test := range tests {
		test := test

		t.Run(fmt.Sprintf("case %d", i+1), func(tt *testing.T) {
			tt.Parallel()

			got := NewCardMapping().GetCardValue(test.input)

			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("cardMapping.GetCardValue(%s) mismatch (-want +got):\n%s", test.input, diff)
			}
		})
	}
}

func TestGetCardValues(t *testing.T) {
	tests := []struct {
		input string
		want  []int
	}{
		{"32T3K", []int{2, 1, 9, 2, 12}},
		{"T55J5", []int{9, 4, 4, 10, 4}},
		{"KK677", []int{12, 12, 5, 6, 6}},
		{"KTJJT", []int{12, 9, 10, 10, 9}},
		{"QQQJA", []int{11, 11, 11, 10, 13}},
	}

	for i, test := range tests {
		test := test

		t.Run(fmt.Sprintf("case %d", i+1), func(tt *testing.T) {
			tt.Parallel()

			got := NewCardMapping().GetCardValues(test.input)

			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("cardMapping.GetCardValues(%s) mismatch (-want +got):\n%s", test.input, diff)
			}
		})
	}
}
