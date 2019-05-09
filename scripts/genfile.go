package main

import (
	"os"

	"github.com/danielstefank/deduplication/distance"
	"github.com/olekukonko/tablewriter"
)

func main() {

	names := []string{"Daniel", "daniel", "Deniel", "dsadmapsomdpoasmd", "ajsdja", "Ddaniel"}

	matrix := distance.CalculateMatrixNaive(names)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(append([]string{""}, names...))
	table.SetBorder(false)                                // Set Border to false
	table.AppendBulk(prependNamesToMatric(matrix, names)) // Add Bulk Data
	table.Render()
}

func prependNamesToMatric(matrix [][]string, names []string) [][]string {
	newMatrix := make([][]string, 0, 0)

	for idx, distances := range matrix {
		name := names[idx]
		newDistances := append([]string{name}, distances...)
		newMatrix = append(newMatrix, newDistances)
	}
	return newMatrix
}
