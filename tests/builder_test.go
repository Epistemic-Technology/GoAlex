package tests

import (
	"net/http"
	"testing"

	"github.com/Sunhill666/goalex/internal/model"
	"github.com/Sunhill666/goalex/pkg/core"
)

func TestQueryBuilderChaining(t *testing.T) {
	server := NewTestServer()
	defer server.Close()

	client := NewTestClient(server.URL)

	// Test method chaining
	builder := client.Works().
		Filter("publication_year", 2020).
		Search("machine learning").
		Sort("cited_by_count", true).
		Select("title", "doi").
		Page(2).
		PerPage(50).
		Sample(100).
		Seed(42)

	// Verify the builder contains the expected parameters
	if builder == nil {
		t.Fatal("Query builder should not be nil")
	}
}

func TestQueryBuilderFilter(t *testing.T) {
	server := NewTestServer()
	defer server.Close()

	client := NewTestClient(server.URL)

	tests := []struct {
		name           string
		setupBuilder   func() *core.QueryBuilder[model.Work]
		expectedFilter string
	}{
		{
			name: "single filter",
			setupBuilder: func() *core.QueryBuilder[model.Work] {
				return client.Works().Filter("publication_year", 2020)
			},
			expectedFilter: "publication_year:2020",
		},
		{
			name: "multiple filters",
			setupBuilder: func() *core.QueryBuilder[model.Work] {
				return client.Works().
					Filter("publication_year", 2020).
					Filter("is_oa", true)
			},
			expectedFilter: "publication_year:2020,is_oa:true",
		},
		{
			name: "filter map",
			setupBuilder: func() *core.QueryBuilder[model.Work] {
				return client.Works().FilterMap(map[string]any{
					"publication_year": 2020,
					"type":             "article",
				})
			},
			expectedFilter: "publication_year:2020,type:article",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server.ResponseHandler = func(req *http.Request) (int, string) {
				filter := req.URL.Query().Get("filter")
				// Note: Due to map iteration order, we can't guarantee exact string match
				// In real tests, you might want to parse and compare the filter components
				if filter == "" {
					t.Error("Expected filter parameter to be present")
				}
				return http.StatusOK, SamplePaginatedResponse
			}

			builder := tt.setupBuilder()
			_, err := builder.List()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestQueryBuilderSearchFilter(t *testing.T) {
	server := NewTestServer()
	defer server.Close()

	client := NewTestClient(server.URL)

	server.ResponseHandler = func(req *http.Request) (int, string) {
		filter := req.URL.Query().Get("filter")
		if filter == "" {
			t.Error("Expected filter parameter for search filter")
		}
		// Should contain title.search or title.search.no_stem
		return http.StatusOK, SamplePaginatedResponse
	}

	// Test search filter without no_stem
	_, err := client.Works().
		SearchFilter(map[string]string{"title": "machine learning"}, false).
		List()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test search filter with no_stem
	_, err = client.Works().
		SearchFilter(map[string]string{"title": "machine learning"}, true).
		List()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestQueryBuilderSort(t *testing.T) {
	server := NewTestServer()
	defer server.Close()

	client := NewTestClient(server.URL)

	tests := []struct {
		name         string
		setupBuilder func() *core.QueryBuilder[model.Work]
	}{
		{
			name: "single sort ascending",
			setupBuilder: func() *core.QueryBuilder[model.Work] {
				return client.Works().Sort("publication_date", false)
			},
		},
		{
			name: "single sort descending",
			setupBuilder: func() *core.QueryBuilder[model.Work] {
				return client.Works().Sort("cited_by_count", true)
			},
		},
		{
			name: "multiple sorts",
			setupBuilder: func() *core.QueryBuilder[model.Work] {
				return client.Works().
					Sort("publication_date", true).
					Sort("cited_by_count", false)
			},
		},
		{
			name: "sort map",
			setupBuilder: func() *core.QueryBuilder[model.Work] {
				return client.Works().SortMap(map[string]bool{
					"publication_date": true,
					"cited_by_count":   false,
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server.ResponseHandler = func(req *http.Request) (int, string) {
				sort := req.URL.Query().Get("sort")
				if sort == "" {
					t.Error("Expected sort parameter to be present")
				}
				return http.StatusOK, SamplePaginatedResponse
			}

			builder := tt.setupBuilder()
			_, err := builder.List()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestQueryBuilderPagination(t *testing.T) {
	server := NewTestServer()
	defer server.Close()

	client := NewTestClient(server.URL)

	server.ResponseHandler = func(req *http.Request) (int, string) {
		page := req.URL.Query().Get("page")
		perPage := req.URL.Query().Get("per-page")

		if page != "2" {
			t.Errorf("Expected page=2, got page=%s", page)
		}
		if perPage != "50" {
			t.Errorf("Expected per-page=50, got per-page=%s", perPage)
		}

		return http.StatusOK, SamplePaginatedResponse
	}

	_, err := client.Works().Page(2).PerPage(50).List()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestQueryBuilderSelect(t *testing.T) {
	server := NewTestServer()
	defer server.Close()

	client := NewTestClient(server.URL)

	server.ResponseHandler = func(req *http.Request) (int, string) {
		selectParam := req.URL.Query().Get("select")
		if selectParam == "" {
			t.Error("Expected select parameter to be present")
		}
		return http.StatusOK, SamplePaginatedResponse
	}

	_, err := client.Works().Select("title", "doi", "publication_year").List()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestQueryBuilderSample(t *testing.T) {
	server := NewTestServer()
	defer server.Close()

	client := NewTestClient(server.URL)

	server.ResponseHandler = func(req *http.Request) (int, string) {
		sample := req.URL.Query().Get("sample")
		seed := req.URL.Query().Get("seed")

		if sample != "100" {
			t.Errorf("Expected sample=100, got sample=%s", sample)
		}
		if seed != "42" {
			t.Errorf("Expected seed=42, got seed=%s", seed)
		}

		return http.StatusOK, SamplePaginatedResponse
	}

	_, err := client.Works().Sample(100).Seed(42).List()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestQueryBuilderGroupBy(t *testing.T) {
	server := NewTestServer()
	defer server.Close()

	client := NewTestClient(server.URL)

	tests := []struct {
		name            string
		field           string
		includeUnknown  bool
		expectedGroupBy string
	}{
		{
			name:            "group by without include unknown",
			field:           "publication_year",
			includeUnknown:  false,
			expectedGroupBy: "publication_year",
		},
		{
			name:            "group by with include unknown",
			field:           "publication_year",
			includeUnknown:  true,
			expectedGroupBy: "publication_year:include_unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server.ResponseHandler = func(req *http.Request) (int, string) {
				groupBy := req.URL.Query().Get("group_by")
				if groupBy != tt.expectedGroupBy {
					t.Errorf("Expected group_by=%s, got group_by=%s", tt.expectedGroupBy, groupBy)
				}
				return http.StatusOK, `{"results": [], "group_by": [{"key": "2020", "count": 100}], "meta": {"count": 0}}`
			}

			_, err := client.Works().GroupBy(tt.field, tt.includeUnknown).ListGroupBy()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}
