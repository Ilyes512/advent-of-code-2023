package day3

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewSchematic(t *testing.T) {
	t.Parallel()

	got := NewSchematic()
	want := &schematic{
		numbers: make(map[coord]*partNumber),
	}

	if diff := cmp.Diff(want, got, cmp.AllowUnexported(schematic{})); diff != "" {
		t.Errorf("NewSchematic() mismatch (-want +got):\n%s", diff)
	}
}

func TestSchematicCalculate(t *testing.T) {
	tests := []struct {
		input []string
		want  []int
	}{
		{
			input: []string{"467..114.."},
			want:  []int{0, 0},
		},
		{
			input: []string{
				"467..114..",
				"...*......",
			},
			want: []int{467, 0},
		},
		{
			input: []string{
				"467..114..",
				"...*......",
				"..35..633.",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			},
			want: []int{4361, 467835},
		},
	}

	for i, test := range tests {
		test := test

		t.Run(fmt.Sprintf("case %d", i+1), func(tt *testing.T) {
			tt.Parallel()

			s := NewSchematic()
			for _, input := range test.input {
				s.AddRow(input)
			}
			got1, got2 := s.Calculate()

			got := []int{got1, got2}

			if diff := cmp.Diff(test.want, got); diff != "" {
				tt.Errorf("Schematic.Calculate() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestAppendNumber(t *testing.T) {
	tests := []struct {
		inputInt  int
		inputRune rune
		want      int
	}{
		{
			inputInt:  1,
			inputRune: '5',
			want:      15,
		},
		{
			inputInt:  23,
			inputRune: '7',
			want:      237,
		},
		{
			inputInt:  0,
			inputRune: '5',
			want:      5,
		},
	}

	for i, test := range tests {
		test := test

		t.Run(fmt.Sprintf("case %d", i+1), func(tt *testing.T) {
			tt.Parallel()

			got := appendNumber(test.inputInt, test.inputRune)

			if diff := cmp.Diff(test.want, got); diff != "" {
				tt.Errorf("appendNumber(%d, '%c') mismatch (-want +got):\n%s", test.inputInt, test.inputRune, diff)
			}
		})
	}
}
