package main

import "fmt"

const (
	PatternHeight = 2
	PatternWidth  = 2
	WaveHeight    = 10
	WaveWidth     = 10
)

type Wave struct {
	Regions [WaveHeight][WaveWidth]Region
}
type Region struct {
	PossiblePatterns []Pattern
}

type Pattern struct {
	Values      [PatternHeight][PatternWidth]int
	Probability float64
}

func (p Pattern) String() string {
	return fmt.Sprintf("Values: %d %d %d %d; Probability: %f", p.Values[0][0], p.Values[0][1], p.Values[1][0], p.Values[1][1], p.Probability)
}

func ExtractPatternValue(input [][]int, row, col int) [PatternHeight][PatternWidth]int {
	var patternValue [PatternHeight][PatternWidth]int
	for i, r := range input[row : row+PatternHeight] {
		copy(patternValue[i][:], r[col:col+PatternWidth])
	}
	return patternValue
}

func ExtractPatterns(input [][]int) []Pattern {
	patternSet := map[[PatternHeight][PatternWidth]int]int{}
	var overallPatternCount int
	for row := range input {
		for col := range input[row] {
			if len(input) >= row+PatternHeight && len(input[col]) >= col+PatternWidth {
				patternValue := ExtractPatternValue(input, row, col)
				count := patternSet[patternValue]
				patternSet[patternValue] = count + 1
				overallPatternCount++
			}
		}
	}

	patterns := []Pattern{}
	for values, count := range patternSet {
		patterns = append(patterns, Pattern{
			Values:      values,
			Probability: float64(count) / float64(overallPatternCount),
		})
	}

	return patterns
}

func main() {
	input := [][]int{
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
	}
	fmt.Println("INPUT", input[1:][1:])
	patterns := ExtractPatterns(input)
	for _, p := range patterns {
		fmt.Println(p)
	}
}
