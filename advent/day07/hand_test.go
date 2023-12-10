package day07

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var compareOptions = []cmp.Option{
	cmp.AllowUnexported(hand{}),
}

func TestNewHands(t *testing.T) {
	got := NewHands("test_input.txt")

	want := hands{
		// 32T3K
		&hand{
			bid:      765,
			cards:    []int{2, 1, 9, 2, 12},
			strength: 26,
			handType: OnePair,
		},
		// KTJJT
		&hand{
			bid:      220,
			cards:    []int{12, 9, 10, 10, 9},
			strength: 50,
			handType: TwoPairs,
		},
		// KK677
		&hand{
			bid:      28,
			cards:    []int{12, 12, 5, 6, 6},
			strength: 41,
			handType: TwoPairs,
		},
		// T55J5
		&hand{
			bid:      684,
			cards:    []int{9, 4, 4, 10, 4},
			strength: 31,
			handType: ThreeOfAKind,
		},
		// QQQJA
		&hand{
			bid:      483,
			cards:    []int{11, 11, 11, 10, 13},
			strength: 56,
			handType: ThreeOfAKind,
		},
	}

	if diff := cmp.Diff(want, got, compareOptions...); diff != "" {
		t.Errorf("TestNewHands() mismatch (-want +got):\n%s", diff)
	}
}

func TestHandsResultPart1(t *testing.T) {
	got := NewHands("test_input.txt").ResultPart1()

	want := 6440

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("hands.ResultPart1() mismatch (-want +got):\n%s", diff)
	}
}
