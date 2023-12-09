package day05

import (
	"fmt"
	"testing"

	slcmp "cmp"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var compareOptions = []cmp.Option{
	cmpopts.SortSlices(func(a, b seedRange) bool {
		return slcmp.Less(a.min, b.min)
	}),
	cmp.AllowUnexported(seedRange{}),
}

var seedToSoilMapping = &mapTable{
	source: "seed",
	dest:   "soil",
	maps: []*mapping{
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

var soilToFertilizerMapping = &mapTable{
	source: "soil",
	dest:   "fertilizer",
	maps: []*mapping{
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
		mapTable    *mapTable
		inputSource int
		want        int
	}{
		{
			mapTable:    seedToSoilMapping,
			inputSource: 79,
			want:        81,
		},
		{
			mapTable:    seedToSoilMapping,
			inputSource: 14,
			want:        14,
		},
		{
			mapTable:    seedToSoilMapping,
			inputSource: 55,
			want:        57,
		},
		{
			mapTable:    seedToSoilMapping,
			inputSource: 13,
			want:        13,
		},

		{
			mapTable:    soilToFertilizerMapping,
			inputSource: 81,
			want:        81,
		},
		{
			mapTable:    soilToFertilizerMapping,
			inputSource: 14,
			want:        53,
		},
		{
			mapTable:    soilToFertilizerMapping,
			inputSource: 57,
			want:        57,
		},
		{
			mapTable:    soilToFertilizerMapping,
			inputSource: 13,
			want:        52,
		},
	}

	for i, test := range tests {
		test := test

		t.Run(fmt.Sprintf("case %d", i+1), func(tt *testing.T) {
			tt.Parallel()

			got := test.mapTable.getMappingFor(test.inputSource)

			if diff := cmp.Diff(test.want, got); diff != "" {
				tt.Errorf("mapTable.getMappingFor(%d) mismatch (-want +got):\n%s", test.inputSource, diff)
			}
		})
	}
}

func TestMappingInRange(t *testing.T) {
	tests := []struct {
		mapping     *mapping
		inputSource int
		want        bool
	}{
		{
			mapping: &mapping{
				source: 5,
				dest:   15,
				len:    10,
			},
			inputSource: 0,
			want:        false,
		},
		{
			mapping: &mapping{
				source: 5,
				dest:   15,
				len:    10,
			},
			inputSource: 7,
			want:        true,
		},
		{
			mapping: &mapping{
				source: 88,
				dest:   18,
				len:    7,
			},
			inputSource: 0,
			want:        false,
		},
		{
			mapping: &mapping{
				source: 88,
				dest:   18,
				len:    7,
			},
			inputSource: 88,
			want:        true,
		},
		{
			mapping: &mapping{
				source: 88,
				dest:   18,
				len:    7,
			},
			inputSource: 95,
			want:        false,
		},
		{
			mapping: &mapping{
				source: 88,
				dest:   18,
				len:    7,
			},
			inputSource: 96,
			want:        false,
		},
		{
			mapping: &mapping{
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

			got := test.mapping.inRange(test.inputSource)

			if diff := cmp.Diff(test.want, got); diff != "" {
				tt.Errorf("mapping.inRange(%d) mismatch (-want +got):\n%s", test.inputSource, diff)
			}
		})
	}
}

func TestMappingDiff(t *testing.T) {
	tests := []struct {
		mapping *mapping
		want    int
	}{
		{
			mapping: &mapping{
				source: 5,
				dest:   15,
				len:    8,
			},
			want: -10,
		},
		{
			mapping: &mapping{
				source: 15,
				dest:   5,
				len:    7,
			},
			want: 10,
		},
		{
			mapping: &mapping{
				source: 98,
				dest:   50,
				len:    2,
			},
			want: 48,
		},
	}

	for i, test := range tests {
		test := test

		t.Run(fmt.Sprintf("case %d", i+1), func(tt *testing.T) {
			tt.Parallel()

			got := test.mapping.diff()

			if diff := cmp.Diff(test.want, got); diff != "" {
				tt.Errorf("mapping.diff() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestSeedRangesTransform(t *testing.T) {
	tests := []struct {
		seedRanges    seedRanges
		inputMapTable *mapTable
		want          seedRanges
	}{
		{
			seedRanges: seedRanges{
				{
					min: 79,
					max: 92,
				},
			},
			inputMapTable: &mapTable{
				source: "source",
				dest:   "dest",
				maps: []*mapping{
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
			},
			want: seedRanges{
				{
					min: 81,
					max: 94,
				},
			},
		},
		{
			seedRanges: seedRanges{
				{
					min: 20,
					max: 85,
				},
			},
			inputMapTable: &mapTable{
				source: "source",
				dest:   "dest",
				maps: []*mapping{
					{
						source: 35,
						dest:   44,
						len:    8,
					},
				},
			},
			want: seedRanges{
				{
					min: 20,
					max: 34,
				},
				{
					min: 44,
					max: 51,
				},
				{
					min: 43,
					max: 85,
				},
			},
		},
		{
			seedRanges: seedRanges{
				{
					min: 20,
					max: 85,
				},
				{
					min: 60,
					max: 65,
				},
			},
			inputMapTable: &mapTable{
				source: "source",
				dest:   "dest",
				maps: []*mapping{
					{
						source: 35,
						dest:   44,
						len:    8,
					},
					{
						source: 60,
						dest:   80,
						len:    15,
					},
				},
			},
			want: seedRanges{
				{
					min: 20,
					max: 34,
				},
				{
					min: 44,
					max: 51,
				},
				{
					min: 43,
					max: 59,
				},
				{
					min: 80,
					max: 94,
				},
				{
					min: 75,
					max: 85,
				},
				{
					min: 80,
					max: 85,
				},
			},
		},
	}

	for i, test := range tests {
		test := test

		t.Run(fmt.Sprintf("case %d", i+1), func(tt *testing.T) {
			tt.Parallel()

			got := test.seedRanges.transform(test.inputMapTable)

			if diff := cmp.Diff(test.want, got, compareOptions...); diff != "" {
				tt.Errorf("seedRanges.transform() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
