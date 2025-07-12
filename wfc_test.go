package main

import "testing"

func TestExtractPatternValue(t *testing.T) {
	testCases := []struct {
		name     string
		input    [][]int
		row, col int
		want     [2][2]int
	}{
		{
			name: "0,0 to 1,1",
			input: [][]int{
				{0, 3, 6},
				{1, 4, 7},
				{2, 5, 8},
			},
			row:  0,
			col:  0,
			want: [2][2]int{{0, 3}, {1, 4}},
		},
		{
			name: "0,1 to 1,2",
			input: [][]int{
				{0, 3, 6},
				{1, 4, 7},
				{2, 5, 8},
			},
			row:  0,
			col:  1,
			want: [2][2]int{{3, 6}, {4, 7}},
		},
		{
			name: "1,4 to 2,5",
			input: [][]int{
				{0, 3, 6},
				{1, 4, 7},
				{2, 5, 8},
			},
			row:  1,
			col:  0,
			want: [2][2]int{{1, 4}, {2, 5}},
		},
		{
			name: "4,7 to 5,8",
			input: [][]int{
				{0, 3, 6},
				{1, 4, 7},
				{2, 5, 8},
			},
			row:  1,
			col:  1,
			want: [2][2]int{{4, 7}, {5, 8}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := ExtractPatternValue(tc.input, tc.row, tc.col)
			if got != tc.want {
				t.Errorf("ExtractPatternValue(%v, %d, %d): Got:%v, Wanted:%v", tc.input, tc.row, tc.col, got, tc.want)
			}
		})
	}
}
