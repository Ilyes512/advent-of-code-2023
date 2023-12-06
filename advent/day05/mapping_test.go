package day05

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMapRangeInRange(t *testing.T) {
	tests := []struct {
		mapRange    *mapRange
		inputSource int
		want        bool
	}{
		{
			mapRange: &mapRange{
				source: 5,
				dest:   15,
				len:    10,
			},
			inputSource: 0,
			want:        false,
		},
		{
			mapRange: &mapRange{
				source: 5,
				dest:   15,
				len:    10,
			},
			inputSource: 7,
			want:        true,
		},
		{
			mapRange: &mapRange{
				source: 88,
				dest:   18,
				len:    7,
			},
			inputSource: 0,
			want:        false,
		},
		{
			mapRange: &mapRange{
				source: 88,
				dest:   18,
				len:    7,
			},
			inputSource: 88,
			want:        true,
		},
		{
			mapRange: &mapRange{
				source: 88,
				dest:   18,
				len:    7,
			},
			inputSource: 95,
			want:        true,
		},
		{
			mapRange: &mapRange{
				source: 88,
				dest:   18,
				len:    7,
			},
			inputSource: 96,
			want:        false,
		},
		{
			mapRange: &mapRange{
				source: 88,
				dest:   18,
				len:    7,
			},
			inputSource: 87,
			want:        false,
		},
	}

	for i, test := range tests {
		test := test

		t.Run(fmt.Sprintf("case %d", i+1), func(tt *testing.T) {
			tt.Parallel()

			got := test.mapRange.inRange(test.inputSource)

			if diff := cmp.Diff(test.want, got); diff != "" {
				tt.Errorf("mapRange.inRange(%d) mismatch (-want +got):\n%s", test.inputSource, diff)
			}
		})
	}
}

var seedToSoilMapping = &mapping{
	source: "seed",
	dest:   "soil",
	ranges: []*mapRange{
		{
			source: 98,
			dest:   50,
			len:    2,
		},
		{
			source: 50,
			dest:   52,
			len:    48,
		},
	},
}

var soilToFertilizerMapping = &mapping{
	source: "soil",
	dest:   "fertilizer",
	ranges: []*mapRange{
		{
			source: 15,
			dest:   0,
			len:    37,
		},
		{
			source: 52,
			dest:   37,
			len:    2,
		},
		{
			source: 0,
			dest:   39,
			len:    15,
		},
	},
}

func TestMappingGetMappingFor(t *testing.T) {
	tests := []struct {
		mapping     *mapping
		inputSource int
		want        int
	}{
		{
			mapping:     seedToSoilMapping,
			inputSource: 79,
			want:        81,
		},
		{
			mapping:     seedToSoilMapping,
			inputSource: 14,
			want:        14,
		},
		{
			mapping:     seedToSoilMapping,
			inputSource: 55,
			want:        57,
		},
		{
			mapping:     seedToSoilMapping,
			inputSource: 13,
			want:        13,
		},

		{
			mapping:     soilToFertilizerMapping,
			inputSource: 81,
			want:        81,
		},
		{
			mapping:     soilToFertilizerMapping,
			inputSource: 14,
			want:        53,
		},
		{
			mapping:     soilToFertilizerMapping,
			inputSource: 57,
			want:        57,
		},
		{
			mapping:     soilToFertilizerMapping,
			inputSource: 13,
			want:        52,
		},
	}

	for i, test := range tests {
		test := test

		t.Run(fmt.Sprintf("case %d", i+1), func(tt *testing.T) {
			tt.Parallel()

			got := test.mapping.getMappingFor(test.inputSource)

			if diff := cmp.Diff(test.want, got); diff != "" {
				tt.Errorf("mapping.getMappingFor(%d) mismatch (-want +got):\n%s", test.inputSource, diff)
			}
		})
	}
}
