package internal

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/pkg/errors"
)

// removeIgnoredFields is a helper function to remove ignored fields from a map or slice of maps.
// It recursively traverses the data structure and removes any fields specified in ignoredFields.
// Parameters:
// - data: The data structure (map or slice of maps) from which fields should be removed.
// - ignoredFields: A variadic list of field names to be removed from the data structure.
func removeIgnoredFields(data interface{}, ignoredFields ...string) {
	switch v := data.(type) {
	case map[string]interface{}:
		for _, field := range ignoredFields {
			delete(v, field)
		}
		for _, value := range v {
			removeIgnoredFields(value, ignoredFields...)
		}
	case []interface{}:
		for _, item := range v {
			removeIgnoredFields(item, ignoredFields...)
		}
	}
}

// EqualJSON is a helper function to compare two JSON strings.
func EqualJSON(expectedJSON, actualJSON string, ignoredFields ...string) (bool, error) {
	expectedData, err := parseJSON(expectedJSON)
	if err != nil {
		return false, errors.Wrap(err, "failed to parse expected JSON")
	}

	actualData, err := parseJSON(actualJSON)
	if err != nil {
		return false, errors.Wrap(err, "failed to parse actual JSON")
	}

	if len(ignoredFields) > 0 {
		removeIgnoredFields(expectedData, ignoredFields...)
		removeIgnoredFields(actualData, ignoredFields...)
	}

	return reflect.DeepEqual(expectedData, actualData), nil
}

// parseJSON parses a JSON string into a generic Go data structure (map or slice).
func parseJSON(input string) (interface{}, error) {
	var data interface{}
	err := json.Unmarshal([]byte(input), &data)
	return data, err
}

// jsonPretty formats a generic Go data structure into a pretty JSON string.
func JsonPretty(data interface{}) string {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Sprintf("<invalid JSON: %v>", err)
	}
	return string(b)
}
