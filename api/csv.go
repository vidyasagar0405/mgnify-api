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
