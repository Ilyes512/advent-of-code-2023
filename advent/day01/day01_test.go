package day01

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCalculatePart1_without_translating_ascii_numbers(t *testing.T) {
	t.Parallel()

	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	document := NewDocument(input)
	want := 142

	got, err := document.Calculate()
	if err != nil {
		t.Fatalf("Unexpected error in Document.Calculate(): %v", err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Document.Calculate() mismatch (-want +got):\n%s", diff)
	}
}

func TestCalculatePart1_with_translating_ascii_numbers(t *testing.T) {
	t.Parallel()

	input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteenk`
	document := NewDocument(input)
	want := 281

	got, err := document.Translate().Calculate()
	if err != nil {
		t.Fatalf("Unexpected error in Document.Calculate(): %v", err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Document.Calculate() mismatch (-want +got):\n%s", diff)
	}
}

func TestGetNumberFromString(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: "1abc2", want: 12},
		{input: "pqr3stu8vwx", want: 38},
		{input: "a1b2c3d4e5f", want: 15},
		{input: "treb7uchet", want: 77},
	}

	for i, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("case %d", i+1), func(st *testing.T) {
			st.Parallel()

			got, err := GetNumberFromString(tc.input)
			if err != nil {
				t.Fatalf("Unexpected error in GetNumberFromString(): %v", err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				st.Errorf("GetNumberFromString(%q) mismatch (-want +got):\n%s", tc.input, diff)
			}
		})
	}
}

func TestGetNumberFromString_returns_an_error_if_string_contains_no_numbers(t *testing.T) {
	t.Parallel()

	input := "foobar"
	_, err := GetNumberFromString(input)
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
}

func TestTranslateAsciiNumbers(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "two1nine", want: "219"},
		{input: "eightwothree", want: "823"},
		{input: "abcone2threexyz", want: "123"},
		{input: "xtwone3four", want: "2134"},
		{input: "4nineeightseven2", want: "49872"},
		{input: "zoneight234", want: "18234"},
		{input: "two1nine", want: "219"},
		{input: "7pqrstsixteen", want: "76"},
	}

	for i, test := range tests {
		test := test
		t.Run(fmt.Sprintf("case %d", i+1), func(tt *testing.T) {
			tt.Parallel()

			got := TranslateAsciiNumbers(test.input)

			if diff := cmp.Diff(test.want, got); diff != "" {
				tt.Errorf("TranslateAsciiNumbers(%q) mismatch (-want +got):\n%s", test.input, diff)
			}
		})
	}
}

func TestGetAsciiNumberAtStart(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "two1nine", want: "2"},
		{input: "eightwothree", want: "8"},
		{input: "one2threexyz", want: "1"},
		{input: "twone3four", want: "2"},
		{input: "nineeightseven2", want: "9"},
		{input: "oneight234", want: "1"},
		{input: "two1nine", want: "2"},
		{input: "sixteen", want: "6"},
	}

	for i, test := range tests {
		test := test
		t.Run(fmt.Sprintf("case %d", i+1), func(tt *testing.T) {
			tt.Parallel()

			got, ok := GetAsciiNumberAtStart(test.input)
			if !ok {
				tt.Fatalf("expected GetAsciiNumberAtStart(%q) to return true", test.input)
			}

			if diff := cmp.Diff(test.want, got); diff != "" {
				tt.Errorf("TestGetAsciiNumberAtStart(%q) mismatch (-want +got):\n%s", test.input, diff)
			}
		})
	}
}
