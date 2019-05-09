package distance

import (
	"strconv"

	"github.com/danielstefank/deduplication/distance/calculator"
)

func CalculateMatrixNaive(names []string) [][]string {
	matrix := make([][]string, 0, 0)

	for _, name := range names {
		matrix = append(matrix, calculateDistances(name, names))
	}

	return matrix
}

func calculateDistances(name string, names []string) []string {
	distances := make([]string, 0, 0)

	for _, coName := range names {
		distances = append(distances, strconv.Itoa(calculator.ComputeDistance(name, coName)))
	}

	return distances
}
