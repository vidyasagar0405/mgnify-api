package main
//
// import (
// 	"flag"
// 	"fmt"
// 	"io"
// 	"os"
// 	"strings"
// )
//
// type stringSlice []string
//
// func (s *stringSlice) String() string {
// 	return strings.Join(*s, ",")
// }
//
// func (s *stringSlice) Set(value string) error {
// 	*s = append(*s, value)
// 	return nil
// }
//
// var (
// 	outputFile string
// 	params     stringSlice
// 	showHelp   bool
// )
//
// func init() {
// 	flag.BoolVar(&showHelp, "h", false, "Show help")
// 	flag.BoolVar(&showHelp, "help", false, "Show help (alias for -h)")
// 	flag.StringVar(&outputFile, "o", "", "Output file (default: stdout)")
// 	flag.Var(&params, "p", "Query parameters as key=value pairs (can be used multiple times)")
// 	flag.Parse()
// }
//
// func main() {
// 	if showHelp {
// 		PrintHelp()
// 		os.Exit(0)
// 	}
//
// 	// Handle endpoint path from positional arguments
// 	args := flag.Args()
//
// 	if len(args) == 0 {
// 		PrintHelp()
// 		os.Exit(1)
// 	}
//
// 	// Initialize components
// 	apiClient := NewAPIClient()
// 	flattener := &JSONFlattener{}
//
// 	// Build endpoint path
// 	endpointPath := strings.Join(args, "/")
//
// 	// Parse query parameters
// 	queryParams := make(map[string]string)
// 	for _, param := range params {
// 		parts := strings.SplitN(param, "=", 2)
// 		if len(parts) != 2 {
// 			fmt.Printf("Invalid parameter format: %s. Use key=value\n", param)
// 			os.Exit(1)
// 		}
// 		queryParams[parts[0]] = parts[1]
// 	}
//
// 	// Fetch data from MGnify
// 	data, err := apiClient.FetchPaginatedData(endpointPath, queryParams)
// 	if err != nil {
// 		fmt.Printf("Error fetching data: %v\n", err)
// 		os.Exit(1)
// 	}
//
// 	// Flatten all records
// 	var flattened []map[string]interface{}
// 	for _, item := range data {
// 		flattened = append(flattened, flattener.Flatten(item))
// 	}
//
// 	// Set up output writer
// 	var writer io.Writer = os.Stdout
// 	if outputFile != "" {
// 		file, err := os.Create(outputFile)
// 		if err != nil {
// 			fmt.Printf("Error creating output file: %v\n", err)
// 			os.Exit(1)
// 		}
// 		defer file.Close()
// 		writer = file
// 	}
//
// 	// Write to CSV
// 	csvWriter := NewCSVWriter(writer)
// 	if err := csvWriter.WriteAll(flattened); err != nil {
// 		fmt.Printf("Error writing CSV: %v\n", err)
// 		os.Exit(1)
// 	}
// }
