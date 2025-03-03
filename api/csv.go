package api

import (
	"encoding/csv"
	"fmt"
	"io"
	"sort"
)

// CSVWriter handles writing flattened data to CSV
type CSVWriter struct {
	Writer *csv.Writer
}

// NewCSVWriter creates a new CSV writer
func NewCSVWriter(w io.Writer) *CSVWriter {
	return &CSVWriter{
		Writer: csv.NewWriter(w),
	}
}

// WriteAll writes all records to CSV with headers
func (cw *CSVWriter) WriteAll(records []map[string]interface{}) error {
	headers := cw.getHeaders(records)
	if err := cw.Writer.Write(headers); err != nil {
		return err
	}

	for _, record := range records {
		row := cw.createRow(headers, record)
		if err := cw.Writer.Write(row); err != nil {
			return err
		}
	}

	cw.Writer.Flush()
	return nil
}

func (cw *CSVWriter) getHeaders(records []map[string]interface{}) []string {
	headerSet := make(map[string]struct{})
	for _, record := range records {
		for key := range record {
			headerSet[key] = struct{}{}
		}
	}

	headers := make([]string, 0, len(headerSet))
	for key := range headerSet {
		headers = append(headers, key)
	}
	sort.Strings(headers)
	return headers
}

func (cw *CSVWriter) createRow(headers []string, record map[string]interface{}) []string {
	row := make([]string, len(headers))
	for i, header := range headers {
		val := record[header]
		if val == nil {
			row[i] = ""
		} else {
			row[i] = fmt.Sprintf("%v", val)
		}
	}
	return row
}

