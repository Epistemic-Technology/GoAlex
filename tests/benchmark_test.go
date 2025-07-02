package tests

import (
	"net/http"
	"testing"

	"github.com/Sunhill666/goalex/pkg/core"
)

// Benchmark tests to measure performance characteristics

func BenchmarkClientCreation(b *testing.B) {
	for b.Loop() {
		client := core.New()
		_ = client
	}
}

func BenchmarkClientCreationWithOptions(b *testing.B) {
	for b.Loop() {
		client := core.New(
			core.PolitePool("test@example.com"),
			core.Auth("test-token"),
			core.WithRetry(3, 0),
		)
		_ = client
	}
}

func BenchmarkQueryBuilderChaining(b *testing.B) {
	client := core.New()

	for b.Loop() {
		builder := client.Works().
			Filter("publication_year", 2020).
			Filter("is_oa", true).
			Search("machine learning").
			Sort("cited_by_count", true).
			Select("title", "doi").
			Page(1).
			PerPage(50)
		_ = builder
	}
}

func BenchmarkQueryParamsGeneration(b *testing.B) {
	params := &core.QueryParams{
		Pagination: &core.PaginationParams{
			Page:    1,
			PerPage: 50,
		},
		Filter: map[string]any{
			"publication_year": 2020,
			"is_oa":            true,
			"type":             "article",
		},
		Search: "machine learning",
		Sort: map[string]bool{
			"publication_date": true,
			"cited_by_count":   false,
		},
		Select: []string{"title", "doi", "publication_year", "cited_by_count"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		query := params.ToQuery()
		_ = query
	}
}

func BenchmarkHTTPRequest(b *testing.B) {
	server := NewTestServer()
	defer server.Close()

	server.SetResponse(http.StatusOK, `{"message": "success"}`)
	client := NewTestClient(server.URL)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result map[string]interface{}
		err := client.Get("/test", &result)
		if err != nil {
			b.Fatalf("Unexpected error: %v", err)
		}
	}
}

func BenchmarkJSONDecoding(b *testing.B) {
	server := NewTestServer()
	defer server.Close()

	server.SetResponse(http.StatusOK, SampleWorkResponse)
	client := NewTestClient(server.URL)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		work, err := client.Works().Get("test-id")
		if err != nil {
			b.Fatalf("Unexpected error: %v", err)
		}
		_ = work
	}
}

func BenchmarkComplexQuery(b *testing.B) {
	server := NewTestServer()
	defer server.Close()

	server.SetResponse(http.StatusOK, SamplePaginatedResponse)
	client := NewTestClient(server.URL)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		works, err := client.Works().
			Filter("publication_year", 2020).
			Filter("is_oa", true).
			Search("machine learning").
			Sort("cited_by_count", true).
			Select("title", "doi", "publication_year").
			Page(1).
			PerPage(50).
			List()
		if err != nil {
			b.Fatalf("Unexpected error: %v", err)
		}
		_ = works
	}
}

// Memory allocation benchmarks
func BenchmarkQueryBuilderMemory(b *testing.B) {
	client := core.New()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		builder := client.Works().
			Filter("publication_year", 2020).
			Search("test").
			Sort("cited_by_count", true).
			Page(1).
			PerPage(25)
		_ = builder
	}
}

func BenchmarkParamsMemory(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		params := &core.QueryParams{
			Pagination: &core.PaginationParams{Page: 1, PerPage: 25},
			Filter:     map[string]any{"year": 2020},
			Search:     "test",
			Sort:       map[string]bool{"count": true},
			Select:     []string{"title", "doi"},
		}
		query := params.ToQuery()
		_ = query
	}
}
