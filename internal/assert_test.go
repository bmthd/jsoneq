package internal_test

import (
	assert "jsonassert/pkg"
	"testing"
)

func TestAssertEqualJSON(t *testing.T) {
	tc := map[string]struct {
		fail          bool
		expectedJSON  string
		actualJSON    string
		ignoredFields []string
	}{
		"Ability to compare correct JSON structures.": {
			fail:         false,
			expectedJSON: `{"a": 1, "b": 2}`,
			actualJSON:   `{"a": 1, "b": 2}`,
		},
	}

	for name, tt := range tc {
		t.Run(name, func(t *testing.T) {
			if tt.fail {
				t.Fatalf("Test case is not implemented yet.")
			}

			assert.AssertEqualJSON(t, tt.expectedJSON, tt.actualJSON)
		})
	}
}
