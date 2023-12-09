package day06

import (
	"fmt"
	"testing"

	slcmp "cmp"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var compareOptions = []cmp.Option{
	cmpopts.SortSlices(func(a, b raceRecord) bool {
		return slcmp.Less(a.time, b.time)
	}),
	cmp.AllowUnexported(raceRecord{}),
}

func TestNewRecordsPart1(t *testing.T) {
	got := NewRecordsPart1("test_input.txt")

	want := records{
		{time: 7, distance: 9},
		{time: 15, distance: 40},
		{time: 30, distance: 200},
	}

	if diff := cmp.Diff(want, got, compareOptions...); diff != "" {
		t.Errorf("NewRecords() mismatch (-want +got):\n%s", diff)
	}
}

func TestNewRecordsPart2(t *testing.T) {
	got := NewRecordsPart2("test_input.txt")

	want := records{
		{time: 71530, distance: 940200},
	}

	if diff := cmp.Diff(want, got, compareOptions...); diff != "" {
		t.Errorf("NewRecords() mismatch (-want +got):\n%s", diff)
	}
}

func TestRecordsPart1(t *testing.T) {
	got := NewRecordsPart1("test_input.txt").Result()

	want := 288

	if diff := cmp.Diff(want, got, compareOptions...); diff != "" {
		t.Errorf("records.Part1() mismatch (-want +got):\n%s", diff)
	}
}

func TestRecordsPart2(t *testing.T) {
	got := NewRecordsPart2("test_input.txt").Result()

	want := 71503

	if diff := cmp.Diff(want, got, compareOptions...); diff != "" {
		t.Errorf("records.Part1() mismatch (-want +got):\n%s", diff)
	}
}

func TestRaceRecordGetNumberOfRecordBreakingRaces(t *testing.T) {
	tests := []struct {
		raceRecord raceRecord
		want       int
	}{
		{
			raceRecord: raceRecord{
				time:     7,
				distance: 9,
			},
			want: 4,
		},
		{
			raceRecord: raceRecord{
				time:     15,
				distance: 40,
			},
			want: 8,
		},
		{
			raceRecord: raceRecord{
				time:     30,
				distance: 200,
			},
			want: 9,
		},
	}

	for i, test := range tests {
		test := test

		t.Run(fmt.Sprintf("case %d", i+1), func(tt *testing.T) {
			tt.Parallel()

			got := test.raceRecord.GetNumberOfRecordBreakingRaces()

			if diff := cmp.Diff(test.want, got, compareOptions...); diff != "" {
				t.Errorf("raceRecord.GetNumberOfRecordBreakingRaces() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
