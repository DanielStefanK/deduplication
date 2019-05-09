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
	}

	app.Action = func(c *cli.Context) error {
		switch mode {
		case "naive":
			{
				naive(fileName, output)
			}
		case "block":
			{

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

func naive(filename string, output string) {
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
			newDistances = append(newDistances, strconv.Itoa(distance))
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
