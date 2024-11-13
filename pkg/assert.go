package pkg

import (
	"jsonassert/internal"
	"testing"
)

// AssertEqualJSON is a helper function to compare two JSON strings.
// Parameters:
// - t: The testing.T instance.
// - expectedJSON: The expected JSON string.
// - actualJSON: The actual JSON string.
// - ignoredFields: A variadic list of field names to be removed from the data structure.
func AssertEqualJSON(t *testing.T, expectedJSON, actualJSON string, ignoredFields ...string) {
	t.Helper()

	equal, err := internal.EqualJSON(expectedJSON, actualJSON, ignoredFields...)
	if err != nil {
		t.Fatalf("failed to compare JSON: %v", err)
	}

	if !equal {
		t.Errorf("JSON mismatch:\nExpected:\n%s\nActual:\n%s", internal.JsonPretty(expectedJSON), internal.JsonPretty(actualJSON))
	}
}
