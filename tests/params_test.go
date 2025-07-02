package tests

import (
	"net/url"
	"strings"
	"testing"

	"github.com/Sunhill666/goalex/pkg/core"
)

func TestPaginationParams(t *testing.T) {
	tests := []struct {
		name     string
		params   *core.PaginationParams
		expected url.Values
	}{
		{
			name: "all pagination parameters",
			params: &core.PaginationParams{
				Page:    2,
				PerPage: 50,
				Cursor:  "next-cursor",
			},
			expected: url.Values{
				"page":     []string{"2"},
				"per-page": []string{"50"},
				"cursor":   []string{"next-cursor"},
			},
		},
		{
			name: "only page parameter",
			params: &core.PaginationParams{
				Page: 3,
			},
			expected: url.Values{
				"page": []string{"3"},
			},
		},
		{
			name:     "empty parameters",
			params:   &core.PaginationParams{},
			expected: url.Values{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.params.ToQuery()
			for key, expectedValues := range tt.expected {
				actualValues, exists := result[key]
				if !exists {
					t.Errorf("Expected key %s not found in result", key)
					continue
				}
				if len(actualValues) != len(expectedValues) {
					t.Errorf("Key %s: expected %d values, got %d", key, len(expectedValues), len(actualValues))
					continue
				}
				for i, expectedValue := range expectedValues {
					if actualValues[i] != expectedValue {
						t.Errorf("Key %s[%d]: expected %s, got %s", key, i, expectedValue, actualValues[i])
					}
				}
			}
			// Check no extra keys exist
			for key := range result {
				if _, exists := tt.expected[key]; !exists {
					t.Errorf("Unexpected key %s found in result", key)
				}
			}
		})
	}
}

func TestQueryParams(t *testing.T) {
	tests := []struct {
		name     string
		params   *core.QueryParams
		expected map[string]string
	}{
		{
			name: "complex query parameters",
			params: &core.QueryParams{
				Pagination: &core.PaginationParams{
					Page:    1,
					PerPage: 25,
				},
				Filter: map[string]any{
					"publication_year": 2020,
					"is_oa":            true,
				},
				Search: "machine learning",
				Sort: map[string]bool{
					"publication_date": true,  // desc
					"cited_by_count":   false, // asc
				},
				Select: []string{"title", "doi", "publication_year"},
				Sample: 100,
				Seed:   42,
			},
			expected: map[string]string{
				"page":     "1",
				"per-page": "25",
				"filter":   "publication_year:2020,is_oa:true",
				"search":   "machine learning",
				"sort":     "publication_date:desc,cited_by_count",
				"select":   "title,doi,publication_year",
				"sample":   "100",
				"seed":     "42",
			},
		},
		{
			name: "minimal parameters",
			params: &core.QueryParams{
				Search: "test",
			},
			expected: map[string]string{
				"search": "test",
			},
		},
		{
			name: "group by parameter",
			params: &core.QueryParams{
				GroupBy: "publication_year:include_unknown",
			},
			expected: map[string]string{
				"group_by": "publication_year:include_unknown",
			},
		},
		{
			name: "autocomplete parameter",
			params: &core.QueryParams{
				AutoComplete: "harvard",
			},
			expected: map[string]string{
				"q": "harvard",
			},
		},
		{
			name: "cursor parameter",
			params: &core.QueryParams{
				Cursor: "*",
			},
			expected: map[string]string{
				"cursor": "*",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.params.ToQuery()

			// Check expected parameters exist with correct values
			for key, expectedValue := range tt.expected {
				actualValue := result.Get(key)
				if key == "filter" || key == "sort" {
					// For filter and sort, just check they're not empty since order varies
					if actualValue == "" {
						t.Errorf("Parameter %s should not be empty", key)
					}
					// For complex parameters, verify all components are present
					if key == "sort" && expectedValue == "publication_date:desc,cited_by_count" {
						if !strings.Contains(actualValue, "publication_date:desc") {
							t.Errorf("Sort parameter should contain 'publication_date:desc', got %s", actualValue)
						}
						if !strings.Contains(actualValue, "cited_by_count") {
							t.Errorf("Sort parameter should contain 'cited_by_count', got %s", actualValue)
						}
						continue
					}
					if key == "filter" && expectedValue == "publication_year:2020,is_oa:true" {
						expectedParts := []string{"publication_year:2020", "is_oa:true"}
						for _, part := range expectedParts {
							if !strings.Contains(actualValue, part) {
								t.Errorf("Filter parameter should contain '%s', got %s", part, actualValue)
							}
						}
						continue
					}
				}
				if actualValue != expectedValue {
					t.Errorf("Parameter %s: expected %s, got %s", key, expectedValue, actualValue)
				}
			}

			// Check no unexpected parameters exist
			for key := range result {
				if _, exists := tt.expected[key]; !exists {
					t.Errorf("Unexpected parameter %s found: %s", key, result.Get(key))
				}
			}
		})
	}
}

func TestQueryParamsFilterOrder(t *testing.T) {
	// Test that filter parameters are handled correctly
	// Note: We can't guarantee order due to Go's map iteration randomness,
	// but we can verify the content is correct
	params := &core.QueryParams{
		Filter: map[string]any{
			"z_field": "value_z",
			"a_field": "value_a",
			"m_field": "value_m",
		},
	}

	result := params.ToQuery()
	filterParam := result.Get("filter")

	// Check that all expected fields are present
	if filterParam == "" {
		t.Error("Filter parameter should not be empty")
	}

	// Verify all fields are present (order may vary)
	expectedFields := []string{"z_field:value_z", "a_field:value_a", "m_field:value_m"}
	for _, expected := range expectedFields {
		if !strings.Contains(filterParam, expected) {
			t.Errorf("Expected filter parameter to contain '%s', got '%s'", expected, filterParam)
		}
	}
}

func TestQueryParamsSortOrder(t *testing.T) {
	// Test that sort parameters are handled correctly
	// Note: We can't guarantee order due to Go's map iteration randomness,
	// but we can verify the content is correct
	params := &core.QueryParams{
		Sort: map[string]bool{
			"z_field": true,
			"a_field": false,
			"m_field": true,
		},
	}

	result := params.ToQuery()
	sortParam := result.Get("sort")

	// Check that all expected fields are present
	if sortParam == "" {
		t.Error("Sort parameter should not be empty")
	}

	// Verify all fields are present (order may vary)
	expectedFields := []string{"z_field:desc", "a_field", "m_field:desc"}
	for _, expected := range expectedFields {
		if !strings.Contains(sortParam, expected) {
			t.Errorf("Expected sort parameter to contain '%s', got '%s'", expected, sortParam)
		}
	}
}
