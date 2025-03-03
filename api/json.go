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

import "strconv"

// JSONFlattener handles flattening nested JSON structures
type JSONFlattener struct{}

// Flatten converts nested JSON structures to flat maps
func (f *JSONFlattener) Flatten(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	f.flattenRecursive(data, "", result)
	return result
}

func (f *JSONFlattener) flattenRecursive(input interface{}, prefix string, result map[string]interface{}) {
	switch v := input.(type) {
	case map[string]interface{}:
		for key, value := range v {
			newKey := f.createKey(prefix, key)
			f.flattenRecursive(value, newKey, result)
		}
	case []interface{}:
		for i, elem := range v {
			newKey := f.createKey(prefix, strconv.Itoa(i))
			f.flattenRecursive(elem, newKey, result)
		}
	default:
		if prefix != "" {
			result[prefix] = v
		}
	}
}

func (f *JSONFlattener) createKey(prefix, key string) string {
	if prefix == "" {
		return key
	}
	return prefix + "." + key
}

