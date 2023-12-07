package day05

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReaderProcess(t *testing.T) {
	reader := NewReader("test_input.txt")
	got := reader.Process()
	want := 35

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Reader.Process() mismatch (-want +got):\n%s", diff)
	}
}

func TestReaderProcessSeeds(t *testing.T) {
	tests := []struct {
		seedInput string
		want      [][]int
	}{
		{
			seedInput: "seeds: 79 14 55 13",
			want: [][]int{
				{55, 68},
				{79, 93},
			},
		},
		{
			seedInput: "seeds: 1 1 3 5 4 3",
			want: [][]int{
				{1, 2},
				{3, 8},
				{4, 7},
			},
		},
		{
			seedInput: "seeds: 60 20 30 10 35 35 20 30 80 20",
			want: [][]int{
				{20, 80},
				{80, 100},
			},
		},
		{
			seedInput: "seeds: 1132132257 323430997 2043754183 4501055 2539071613 1059028389 1695770806 60470169 2220296232 251415938 1673679740 6063698 962820135 133182317 262615889 327780505 3602765034 194858721 2147281339 37466509",
			want: [][]int{
				{262615889, 590396394},
				{962820135, 1096002452},
				{1132132257, 1455563254},
				{1673679740, 1679743438},
				{1695770806, 1756240975},
				{2043754183, 2048255238},
				{2147281339, 2184747848},
				{2220296232, 2471712170},
				{2539071613, 3598100002},
				{3602765034, 3797623755},
			},
		},
	}

	for i, test := range tests {
		test := test

		t.Run(fmt.Sprintf("case %d", i+1), func(tt *testing.T) {
			tt.Parallel()

			reader := NewReader("test_input.txt")
			reader.processSeeds(test.seedInput)

			got := reader.seedPairs

			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("Reader.processSeeds(%s) mismatch (-want +got):\n%s", test.seedInput, diff)
			}
		})
	}
}
