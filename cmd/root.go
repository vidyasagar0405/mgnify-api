/*
Copyright © 2025 Vidyasagar Gopi vidyasagar0405@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"mgnify-api/api"

	"github.com/spf13/cobra"
)

var params []string
var outputFile string
var args []string

func init() {
	// Define persistent flags
	rootCmd.PersistentFlags().StringSliceVarP(&params, "param", "p", []string{}, "Params for the query")

	// Define local flags
	defaultOut := time.Now().Format("2006_01_02_15_04_05") + ".csv"
	rootCmd.Flags().StringVarP(&outputFile, "outputFile", "o", defaultOut, "Output file name")

    // Define Args
    args = rootCmd.Flags().Args()
}

var rootCmd = &cobra.Command{
	Use:   `mgnify-cli [args] [flags]...
  args - API endpoints
  flags - "-p/--params" and "-o/--outputFile"

  Examples:
    1. Basic usage with output file:
       mgnify-cli biomes abc123 studies -o biome_studies.csv

    2. Multiple query parameters:
       mgnify-cli -p experiment_type=metagenomic -p page_size=100 samples -o meta.csv

    3. Complex endpoint with parameters:
       mgnify-cli studies recent -p lineage="root:Host-associated" -o host_studies.csv
    `,
	Short: `Fetches data from the MGnify API, flattens nested JSON structures,
and exports the results to CSV. Handles pagination automatically.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Access flag values during command execution
		fmt.Println("Output File:", outputFile)
		fmt.Println("Params:", params)
        fmt.Println(args)

		// Initialize components
		apiClient := api.NewAPIClient()
		flattener := &api.JSONFlattener{}

		// Build endpoint path
		endpointPath := strings.Join(args, "/")

		// Parse query parameters
		queryParams := make(map[string]string)
		for _, param := range params {
			parts := strings.SplitN(param, "=", 2)
			if len(parts) != 2 {
				fmt.Printf("Invalid parameter format: %s. Use key=value\n", param)
				os.Exit(1)
			}
			queryParams[parts[0]] = parts[1]
		}

		// Fetch data from MGnify
		data, err := apiClient.FetchPaginatedData(endpointPath, queryParams)
		if err != nil {
			fmt.Printf("Error fetching data: %v\n", err)
			os.Exit(1)
		}

		// Flatten all records
		var flattened []map[string]interface{}
		for _, item := range data {
			flattened = append(flattened, flattener.Flatten(item))
		}

		// Set up output writer
		var writer io.Writer = os.Stdout
		if outputFile != "" {
			file, err := os.Create(outputFile)
			if err != nil {
				fmt.Printf("Error creating output file: %v\n", err)
				os.Exit(1)
			}
			defer file.Close()
			writer = file
		}

		// Write to CSV
		csvWriter := api.NewCSVWriter(writer)
		if err := csvWriter.WriteAll(flattened); err != nil {
			fmt.Printf("Error writing CSV: %v\n", err)
			os.Exit(1)
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}
