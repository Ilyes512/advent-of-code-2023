package day05

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReaderProcessPart1(t *testing.T) {
	reader := NewReader("test_input.txt")
	got := reader.ProcessPart1()
	want := 35

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Reader.ProcessPart1() mismatch (-want +got):\n%s", diff)
	}
}

func TestReaderProcessPart2(t *testing.T) {
	reader := NewReader("test_input.txt")
	got := reader.ProcessPart2()
	want := 46

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Reader.ProcessPart2() mismatch (-want +got):\n%s", diff)
	}
}

func TestReaderProcessSeeds(t *testing.T) {
	tests := []struct {
		seedInput string
		want      seedRanges
	}{
		{
			seedInput: "seeds: 79 14 55 13",
			want: seedRanges{
				{min: 79, max: 93},
				{min: 55, max: 68},
			},
		},
		{
			seedInput: "seeds: 1 1 3 5 4 3",
			want: seedRanges{
				{min: 1, max: 2},
				{min: 3, max: 8},
				{min: 4, max: 7},
			},
		},
		{
			seedInput: "seeds: 60 20 30 10 35 35 20 30 80 20",
			want: seedRanges{
				{min: 60, max: 80},
				{min: 30, max: 40},
				{min: 35, max: 70},
				{min: 20, max: 50},
				{min: 80, max: 100},
			},
		},
		{
			seedInput: "seeds: 1132132257 323430997 2043754183 4501055 2539071613 1059028389 1695770806 60470169 2220296232 251415938 1673679740 6063698 962820135 133182317 262615889 327780505 3602765034 194858721 2147281339 37466509",
			want: seedRanges{
				{min: 1132132257, max: 1455563254},
				{min: 2043754183, max: 2048255238},
				{min: 2539071613, max: 3598100002},
				{min: 1695770806, max: 1756240975},
				{min: 2220296232, max: 2471712170},
				{min: 1673679740, max: 1679743438},
				{min: 962820135, max: 1096002452},
				{min: 262615889, max: 590396394},
				{min: 3602765034, max: 3797623755},
				{min: 2147281339, max: 2184747848},
			},
		},
	}

	for i, test := range tests {
		test := test

		t.Run(fmt.Sprintf("case %d", i+1), func(tt *testing.T) {
			tt.Parallel()

			reader := NewReader("test_input.txt")
			reader.processSeeds(test.seedInput)

			got := reader.seedRanges

			if diff := cmp.Diff(test.want, got, compareOptions...); diff != "" {
				t.Errorf("Reader.processSeeds(%s) mismatch (-want +got):\n%s", test.seedInput, diff)
			}
		})
	}
}
