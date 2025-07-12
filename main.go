package main

import "fmt"

type Wave struct {
}

type Pattern struct {
	values      [4]int
	probability float64
}

func (p Pattern) String() string {
	return fmt.Sprintf("Values: %d %d %d %d; Probability: %f", p.values[0], p.values[1], p.values[2], p.values[3], p.probability)
}

func extractPatterns(input [][]int, n int) []Pattern {
	patternSet := map[[4]int]int{}
	var overallPatternCount int
	for row := range input {
		for col := range input[row] {
			if len(input) > col+1 && len(input[col]) > row+1 {
				pattern := [4]int{input[row][col], input[row][col+1], input[row+1][col], input[row+1][col+1]}
				count := patternSet[pattern]
				patternSet[pattern] = count + 1
				overallPatternCount++
			}
		}
	}

	patterns := []Pattern{}
	for values, count := range patternSet {
		patterns = append(patterns, Pattern{
			values:      values,
			probability: float64(count) / float64(overallPatternCount),
		})
	}

	return patterns
}

func main() {
	n := 2
	input := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	patterns := extractPatterns(input, n)
	for _, p := range patterns {
		fmt.Println(p)
	}
}
