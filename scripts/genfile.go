package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/danielstefank/deduplication/distance"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

func main() {
	var mode string
	var fileName string
	var output string
	var length int
	var threshold int

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "mode,m",
			Value:       "naive",
			Usage:       "Mode can be naive, block",
			Destination: &mode,
		},
		cli.StringFlag{
			Name:        "file,f",
			Value:       "./data.txt",
			Usage:       "the file to read the data from",
			Destination: &fileName,
		},
		cli.StringFlag{
			Name:        "output,o",
			Value:       "./matrix.txt",
			Usage:       "the file where the data sould be written to",
			Destination: &output,
		},
		cli.IntFlag{
			Name:        "threshold,t",
			Value:       2,
			Usage:       "the threshold to find duplicates",
			Destination: &threshold,
		},
		cli.IntFlag{
			Name:        "length, l",
			Value:       10,
			Usage:       "the length of the blocks",
			Destination: &length,
		},
	}

	app.Action = func(c *cli.Context) error {
		switch mode {
		case "naive":
			{
				naive(fileName, output, threshold)
			}
		case "block":
			{
				block(fileName, output, length, threshold)
			}
		default:
			{
				fmt.Println("mode not found")
			}
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func naive(filename string, output string, threshold int) {
	names := readInFile(filename)
	matrix := distance.CalculateMatrixNaive(names)
	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	// table.SetHeader(append([]string{""}, names...))
	table.SetBorder(false) // Set Border to false
	//table.AppendBulk(prependNamesToMatrix(convertMatrixToString(matrix), names))
	table.AppendBulk(convertMatrixToString(matrix)) // Add Bulk Data
	table.Render()
	d1 := []byte(tableString.String())
	err := ioutil.WriteFile(output, d1, 0644)
	if err != nil {
		fmt.Println("error writing file")
		panic(err)
	}

	d2 := []byte(calcDuplicates(matrix, names, threshold))
	err = ioutil.WriteFile("distances.txt", d2, 0644)
}

func block(filename string, output string, length int, threshold int) {
	names := readInFile(filename)
	matrix := distance.CalculateMatrixBlock(names, length)
	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	// table.SetHeader(append([]string{""}, names...))
	table.SetBorder(false) // Set Border to false
	//table.AppendBulk(prependNamesToMatrix(convertMatrixToString(matrix), names))
	table.AppendBulk(convertMatrixToString(matrix)) // Add Bulk Data
	table.Render()
	d1 := []byte(tableString.String())
	err := ioutil.WriteFile(output, d1, 0644)

	if err != nil {
		fmt.Println("error writing file")
		panic(err)
	}

	d2 := []byte(calcDuplicates(matrix, names, threshold))
	err = ioutil.WriteFile("distances.txt", d2, 0644)

}

func prependNamesToMatrix(matrix [][]string, names []string) [][]string {
	newMatrix := make([][]string, 0, 0)

	for idx, distances := range matrix {
		name := names[idx]
		newDistances := append([]string{name}, distances...)
		newMatrix = append(newMatrix, newDistances)
	}

	return newMatrix
}

func convertMatrixToString(matrix [][]int) [][]string {
	newMatrix := make([][]string, 0, 0)
	for _, distances := range matrix {
		newDistances := make([]string, 0, 0)
		for _, distance := range distances {
			var i string
			if distance == -1 {
				i = ""
			} else {
				i = strconv.Itoa(distance)
			}
			newDistances = append(newDistances, i)
		}
		newMatrix = append(newMatrix, newDistances)
	}

	return newMatrix
}

func check(e error) {
}

func readInFile(name string) []string {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("file not found")
		panic(err)
	}
	return strings.Split(string(data), "\n")
}

func calcDuplicates(matrix [][]int, names []string, threshold int) string {
	var sb2 strings.Builder
	for i := 0; i < len(matrix); i++ {
		var sb strings.Builder
		for j := i; j < len(matrix); j++ {
			if matrix[i][j] <= threshold && matrix[i][j] > 0 {
				fmt.Fprintf(&sb, "%s, ", names[j])
			}
		}
		if sb.Len() != 0 {
			fmt.Fprintf(&sb2, "%s: %s\n", names[i], sb.String())
		}
	}
	return sb2.String()
}
