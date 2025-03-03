package main

import "fmt"

func PrintHelp() {

	helpString := `MGnify API Data Exporter

Usage:
  mgnify-cli [FLAGS] ENDPOINT_PARTS...

Description:
  Fetches data from the MGnify API, flattens nested JSON structures,
  and exports the results to CSV. Handles pagination automatically.

Positional Arguments:
  ENDPOINT_PARTS...  API endpoint components to join with slashes
                     (e.g.: "biomes abc123 studies" → biomes/abc123/studies)

Flags:
  -o, --output <file>  Output CSV file (default: stdout)
  -p <key=value>       Query parameters (multiple allowed)
                       (e.g.: -p experiment_type=assembly -p include=publications)
  -h, --help           Show this help message

Examples:
  1. Basic usage with output file:
     mgnify-cli biomes abc123 studies -o biome_studies.csv

  2. Multiple query parameters:
     mgnify-cli -p experiment_type=metagenomic -p page_size=100 samples -o meta.csv

  3. Complex endpoint with parameters:
     mgnify-cli studies recent -p lineage="root:Host-associated" -o host_studies.csv

  4. Output to stdout (terminal):
     mgnify-cli runs -p biome_name=soil

Version: 1.0.0`

	fmt.Println(helpString)
}
