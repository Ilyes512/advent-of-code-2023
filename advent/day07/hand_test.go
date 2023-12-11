package day07

import (
	"fmt"
	"slices"
	"testing"

	slcmp "cmp"

	"github.com/google/go-cmp/cmp"
)

var compareOptions = []cmp.Option{
	cmp.AllowUnexported(hand{}),
}

func TestNewHands(t *testing.T) {
	got, _ := NewHands("test_input.txt")

	want := hands{
		hand{
			bid:           765,
			cardsString:   "32T3K",
			originalCards: []int{2, 1, 9, 2, 12},
			cards:         []int{2, 1, 9, 2, 12},
			strength:      26,
			handType:      OnePair,
		},
		hand{
			bid:           220,
			cardsString:   "KTJJT",
			originalCards: []int{12, 9, 10, 10, 9},
			cards:         []int{12, 9, 10, 10, 9},
			strength:      50,
			handType:      TwoPairs,
		},
		hand{
			bid:           28,
			cardsString:   "KK677",
			originalCards: []int{12, 12, 5, 6, 6},
			cards:         []int{12, 12, 5, 6, 6},
			strength:      41,
			handType:      TwoPairs,
		},
		hand{
			bid:           684,
			cardsString:   "T55J5",
			originalCards: []int{9, 4, 4, 10, 4},
			cards:         []int{9, 4, 4, 10, 4},
			strength:      31,
			handType:      ThreeOfAKind,
		},
		hand{
			bid:           483,
			cardsString:   "QQQJA",
			originalCards: []int{11, 11, 11, 10, 13},
			cards:         []int{11, 11, 11, 10, 13},
			strength:      56,
			handType:      ThreeOfAKind,
		},
	}

	if diff := cmp.Diff(want, got, compareOptions...); diff != "" {
		t.Errorf("TestNewHands() mismatch (-want +got):\n%s", diff)
	}
}

func TestNewHandsWithJoker(t *testing.T) {
	_, got := NewHands("test_input.txt")

	want := hands{
		hand{
			bid:           765,
			cardsString:   "32T3K",
			originalCards: []int{2, 1, 9, 2, 12},
			cards:         []int{2, 1, 9, 2, 12},
			strength:      26,
			handType:      OnePair,
		},
		hand{
			bid:           28,
			cardsString:   "KK677",
			originalCards: []int{12, 12, 5, 6, 6},
			cards:         []int{12, 12, 5, 6, 6},
			strength:      41,
			handType:      TwoPairs,
		},
		hand{
			bid:           684,
			cardsString:   "T55J5",
			originalCards: []int{9, 4, 4, 10, 4},
			cards:         []int{9, 4, 4, 0, 4},
			strength:      21,
			handType:      FourOfAKind,
		},
		hand{
			bid:           483,
			cardsString:   "QQQJA",
			originalCards: []int{11, 11, 11, 10, 13},
			cards:         []int{11, 11, 11, 0, 13},
			strength:      46,
			handType:      FourOfAKind,
		},
		hand{
			bid:           220,
			cardsString:   "KTJJT",
			originalCards: []int{12, 9, 10, 10, 9},
			cards:         []int{12, 9, 0, 0, 9},
			strength:      30,
			handType:      FourOfAKind,
		},
	}

	if diff := cmp.Diff(want, got, compareOptions...); diff != "" {
		t.Errorf("TestNewHands() mismatch (-want +got):\n%s", diff)
	}
}

func TestHandsResult(t *testing.T) {
	hands, _ := NewHands("test_input.txt")
	got := hands.Result()

	want := 6440

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("hands.Result() mismatch (-want +got):\n%s", diff)
	}
}

func TestHandsWithJokerResult(t *testing.T) {
	_, handsWithJoker := NewHands("test_input.txt")
	got := handsWithJoker.Result()

	want := 5905

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("hands.Result() (with joker) mismatch (-want +got):\n%s", diff)
	}
}

func TestNewHand(t *testing.T) {
	tests := []struct {
		inputBid         string
		inputCardsString string
		inputCards       []int
		want             hand
	}{
		{
			inputBid:         "333",
			inputCardsString: "J4J5J",
			inputCards:       []int{10, 3, 10, 4, 10},
			want: hand{
				bid:           333,
				cardsString:   "J4J5J",
				originalCards: []int{10, 3, 10, 4, 10},
				cards:         []int{10, 3, 10, 4, 10},
				strength:      37,
				handType:      ThreeOfAKind,
			},
		},
		{
			inputBid:         "765",
			inputCardsString: "JJJJJ",
			inputCards:       []int{10, 10, 10, 10, 10},
			want: hand{
				bid:           765,
				cardsString:   "JJJJJ",
				originalCards: []int{10, 10, 10, 10, 10},
				cards:         []int{10, 10, 10, 10, 10},
				strength:      50,
				handType:      FiveOfAKind,
			},
		},
	}

	for i, test := range tests {
		test := test

		t.Run(fmt.Sprintf("case %d", i+1), func(tt *testing.T) {
			tt.Parallel()

			got := NewHand(test.inputBid, test.inputCardsString, test.inputCards)

			if diff := cmp.Diff(test.want, got, compareOptions...); diff != "" {
				t.Errorf("NewHand() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestNewHandWithJoker(t *testing.T) {
	tests := []struct {
		inputBid         string
		inputCardsString string
		inputCards       []int
		want             hand
	}{
		{
			inputBid:         "333",
			inputCardsString: "J4J5J",
			inputCards:       []int{10, 3, 10, 4, 10},
			want: hand{
				bid:           333,
				cardsString:   "J4J5J",
				originalCards: []int{10, 3, 10, 4, 10},
				cards:         []int{0, 3, 0, 4, 0},
				strength:      7,
				handType:      FourOfAKind,
			},
		},
		{
			inputBid:         "765",
			inputCardsString: "JJJJJ",
			inputCards:       []int{10, 10, 10, 10, 10},
			want: hand{
				bid:           765,
				cardsString:   "JJJJJ",
				originalCards: []int{10, 10, 10, 10, 10},
				cards:         []int{0, 0, 0, 0, 0},
				strength:      0,
				handType:      FiveOfAKind,
			},
		},
	}

	for i, test := range tests {
		test := test

		t.Run(fmt.Sprintf("case %d", i+1), func(tt *testing.T) {
			tt.Parallel()

			got := NewHandWithJoker(test.inputBid, test.inputCardsString, test.inputCards)

			if diff := cmp.Diff(test.want, got, compareOptions...); diff != "" {
				t.Errorf("NewHand() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestSortHandsWithJoker(t *testing.T) {
	tests := []struct {
		input hands
		want  hands
	}{
		{
			input: hands{
				hand{
					bid:           5,
					cardsString:   "32T3K",
					originalCards: []int{2, 1, 9, 2, 12},
					cards:         []int{2, 1, 9, 2, 12},
					strength:      26,
					handType:      OnePair,
				},
				hand{
					bid:           2,
					cardsString:   "J345A",
					originalCards: []int{10, 2, 3, 4, 13},
					cards:         []int{0, 2, 3, 4, 13},
					strength:      22,
					handType:      OnePair,
				},
				hand{
					bid:           3,
					cardsString:   "2345J",
					originalCards: []int{1, 2, 3, 4, 10},
					cards:         []int{1, 2, 3, 4, 0},
					strength:      10,
					handType:      OnePair,
				},
				hand{
					bid:           1,
					cardsString:   "2345A",
					originalCards: []int{1, 2, 3, 4, 13},
					cards:         []int{1, 2, 3, 4, 13},
					strength:      23,
					handType:      HighCard,
				},
			},
			want: hands{
				hand{
					bid:           1,
					cardsString:   "2345A",
					originalCards: []int{1, 2, 3, 4, 13},
					cards:         []int{1, 2, 3, 4, 13},
					strength:      23,
					handType:      HighCard,
				},
				hand{
					bid:           2,
					cardsString:   "J345A",
					originalCards: []int{10, 2, 3, 4, 13},
					cards:         []int{0, 2, 3, 4, 13},
					strength:      22,
					handType:      OnePair,
				},
				hand{
					bid:           3,
					cardsString:   "2345J",
					originalCards: []int{1, 2, 3, 4, 10},
					cards:         []int{1, 2, 3, 4, 0},
					strength:      10,
					handType:      OnePair,
				},
				hand{
					bid:           5,
					cardsString:   "32T3K",
					originalCards: []int{2, 1, 9, 2, 12},
					cards:         []int{2, 1, 9, 2, 12},
					strength:      26,
					handType:      OnePair,
				},
			},
		},
		{
			input: hands{
				hand{
					bid:           43,
					cardsString:   "JAAAA",
					originalCards: []int{10, 13, 13, 13, 13},
					cards:         []int{0, 13, 13, 13, 13},
					strength:      52,
					handType:      FiveOfAKind,
				},
				hand{
					bid:           17,
					cardsString:   "T3T3J",
					originalCards: []int{9, 2, 9, 2, 10},
					cards:         []int{9, 2, 9, 2, 0},
					strength:      22,
					handType:      FullHouse,
				},
				hand{
					bid:           19,
					cardsString:   "Q2Q2Q",
					originalCards: []int{11, 1, 11, 1, 11},
					cards:         []int{11, 1, 11, 1, 11},
					strength:      35,
					handType:      FullHouse,
				},
				hand{
					bid:           59,
					cardsString:   "AAAAJ",
					originalCards: []int{13, 13, 13, 13, 10},
					cards:         []int{13, 13, 13, 13, 0},
					strength:      52,
					handType:      FiveOfAKind,
				},
			},
			want: hands{
				hand{
					bid:           17,
					cardsString:   "T3T3J",
					originalCards: []int{9, 2, 9, 2, 10},
					cards:         []int{9, 2, 9, 2, 0},
					strength:      22,
					handType:      FullHouse,
				},
				hand{
					bid:           19,
					cardsString:   "Q2Q2Q",
					originalCards: []int{11, 1, 11, 1, 11},
					cards:         []int{11, 1, 11, 1, 11},
					strength:      35,
					handType:      FullHouse,
				},
				hand{
					bid:           43,
					cardsString:   "JAAAA",
					originalCards: []int{10, 13, 13, 13, 13},
					cards:         []int{0, 13, 13, 13, 13},
					strength:      52,
					handType:      FiveOfAKind,
				},
				hand{
					bid:           59,
					cardsString:   "AAAAJ",
					originalCards: []int{13, 13, 13, 13, 10},
					cards:         []int{13, 13, 13, 13, 0},
					strength:      52,
					handType:      FiveOfAKind,
				},
			},
		},
	}

	for i, test := range tests {
		test := test

		t.Run(fmt.Sprintf("case %d", i+1), func(tt *testing.T) {
			tt.Parallel()

			got := sortHands(test.input)

			if diff := cmp.Diff(test.want, got, compareOptions...); diff != "" {
				t.Errorf("sortHands() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestSorting(t *testing.T) {
	type MyInt struct {
		Value int
	}

	randomNumbers := []int{5, 2, 3, 1, 4}

	originalSlice := make([]MyInt, len(randomNumbers))
	for i, v := range randomNumbers {
		originalSlice[i] = MyInt{Value: v + 1}
	}

	sliceA := make([]MyInt, len(originalSlice))
	copy(sliceA, originalSlice)
	sliceB := originalSlice

	sliceA[0].Value = 999

	slices.SortFunc(sliceA, func(a, b MyInt) int {
		return slcmp.Compare(a.Value, b.Value)
	})

	fmt.Println("originalSlice:")
	for _, v := range originalSlice {
		fmt.Printf("%#v ", v)
	}
	fmt.Println()

	fmt.Println("siceA:")
	for _, v := range sliceA {
		fmt.Printf("%#v ", v)
	}
	fmt.Println()

	fmt.Println("sliceB:")
	for _, v := range sliceB {
		fmt.Printf("%#v ", v)
	}
	fmt.Println()
}
