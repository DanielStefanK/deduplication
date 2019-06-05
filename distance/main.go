package distance

import (
	"math"

	"github.com/danielstefank/deduplication/distance/calculator"
)

func CalculateMatrixNaive(names []string) [][]int {
	matrix := make([][]int, 0, 0)

	for _, name := range names {
		matrix = append(matrix, calculateDistances(name, names))
	}

	return matrix
}

func getMinusArray(length int) []int {
	arr := make([]int, 0, 0)

	for i := 0; i < length; i++ {
		arr = append(arr, -1)
	}

	return arr
}

func CalculateMatrixBlock(names []string, length int) [][]int {
	matrix := make([][]int, 0, 0)

	for i := 0; i < len(names); i = i + length {
		start := int(math.Min(float64((i)), float64(len(names)-1)))
		end := int(math.Min(float64((i)+length), float64(len(names)-1)))
		for j := start; j < end; j++ {
			name := names[j]
			block := calculateDistances(name, names[start:end])
			matrix = append(matrix, append(append(getMinusArray(start), block...), getMinusArray(len(names)-end)...))
		}
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
