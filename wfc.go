package main

import (
	"fmt"
	"math"
)

const (
	PatternHeight = 2
	PatternWidth  = 2
	WaveHeight    = 10
	WaveWidth     = 10
)

type Wave struct {
	Regions [WaveHeight][WaveWidth]Region
}

func NewWave(patterns []Pattern) Wave {
	regions := [WaveHeight][WaveWidth]Region{}
	for row := range WaveHeight {
		for col := range WaveWidth {
			regions[row][col] = NewRegion(patterns)
		}
	}

	return Wave{Regions: regions}
}

func (e Region) CalculateShannonEntropy() float64 {
	var sum float64
	for pattern, coefficient := range e.PatternCoefficients {
		if !coefficient || pattern.Probability == 0 {
			continue
		}
		result := pattern.Probability * math.Log2(pattern.Probability)
		sum += result
	}
	return -sum
}

func NewRegion(patterns []Pattern) Region {
	coefficients := make(map[Pattern]bool, len(patterns))
	for key := range coefficients {
		coefficients[key] = true
	}

	return Region{PatternCoefficients: coefficients}
}

type Pattern struct {
	Values            [PatternHeight][PatternWidth]int
	Probability       float64
	AdjacentNorth     [][PatternHeight][PatternWidth]int
	AdjacentNorthEast [][PatternHeight][PatternWidth]int
	AdjacentEast      [][PatternHeight][PatternWidth]int
	AdjacentSouthEast [][PatternHeight][PatternWidth]int
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

func ExtractPatternPairs() {

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
