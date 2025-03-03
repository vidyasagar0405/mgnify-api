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

