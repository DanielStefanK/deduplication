package distance

import (
	"github.com/danielstefank/deduplication/distance/calculator"
)

func CalculateMatrixNaive(names []string) [][]int {
	matrix := make([][]int, 0, 0)

	for _, name := range names {
		matrix = append(matrix, calculateDistances(name, names))
	}

	return matrix
}

func calculateDistances(name string, names []string) []int {
	distances := make([]int, 0, 0)

	for _, coName := range names {
		distances = append(distances, calculator.ComputeDistance(name, coName))
	}

	return distances
}
