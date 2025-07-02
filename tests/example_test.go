// Package tests demonstrates comprehensive testing for the GoAlex library.
// This example shows how users can test their applications that use GoAlex.
package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Sunhill666/goalex"
)

// ExampleTest demonstrates how to test an application function that uses GoAlex
func TestUserApplicationExample(t *testing.T) {
	// This test shows how a user would test their own code that uses GoAlex

	// Setup mock server
	server := NewTestServer()
	defer server.Close()

	// Configure the response for the test
	server.SetResponse(http.StatusOK, SamplePaginatedResponse)

	// Create a client configured for testing
	client := goalex.NewClient(goalex.PolitePool("test@example.com"))
	client.BaseURL = server.URL

	// Test a function that uses the client
	works, err := fetchRecentOAWorks(client, 2020)
	if err != nil {
		t.Fatalf("Failed to fetch works: %v", err)
	}

	if len(works) != 1 {
		t.Errorf("Expected 1 work, got %d", len(works))
	}

	if works[0].Title != "The state of OA: a large-scale analysis of the prevalence and impact of Open Access articles" {
		t.Errorf("Unexpected work title: %s", works[0].Title)
	}
}

// fetchRecentOAWorks is an example function that a user might write
// This demonstrates how they would structure their code to be testable
func fetchRecentOAWorks(client *goalex.Client, year int) ([]*WorkSummary, error) {
	// Use the GoAlex client to fetch works
	works, err := client.Works().
		Filter("publication_year", year).
		Filter("is_oa", true).
		Sort("cited_by_count", true).
		PerPage(50).
		List()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch works from OpenAlex: %w", err)
	}

	// Transform the results into our application's format
	var summaries []*WorkSummary
	for _, work := range works {
		summary := &WorkSummary{
			ID:    work.ID,
			Title: work.Title,
			Year:  work.PublicationYear,
		}
		summaries = append(summaries, summary)
	}

	return summaries, nil
}

// WorkSummary represents how a user application might structure work data
type WorkSummary struct {
	ID    string
	Title string
	Year  int
}

// TestErrorHandling demonstrates testing error scenarios
func TestErrorHandlingExample(t *testing.T) {
	server := NewTestServer()
	defer server.Close()

	// Test API error handling
	server.SetResponse(http.StatusNotFound, `{"error": "not found"}`)

	client := goalex.NewClient()
	client.BaseURL = server.URL

	works, err := fetchRecentOAWorks(client, 2020)
	if err == nil {
		t.Error("Expected error when API returns 404")
	}

	if works != nil {
		t.Error("Expected nil works on error")
	}
}

// TestClientConfiguration demonstrates testing different client configurations
func TestClientConfigurationExample(t *testing.T) {
	tests := []struct {
		name   string
		config func() *goalex.Client
		test   func(t *testing.T, client *goalex.Client)
	}{
		{
			name: "with polite pool",
			config: func() *goalex.Client {
				return goalex.NewClient(goalex.PolitePool("test@example.com"))
			},
			test: func(t *testing.T, client *goalex.Client) {
				if client.MailTo != "test@example.com" {
					t.Errorf("Expected MailTo to be set")
				}
			},
		},
		{
			name: "with authentication",
			config: func() *goalex.Client {
				return goalex.NewClient(goalex.Auth("test-token"))
			},
			test: func(t *testing.T, client *goalex.Client) {
				if client.Token != "test-token" {
					t.Errorf("Expected Token to be set")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := tt.config()
			tt.test(t, client)
		})
	}
}

// TestComplexQueryExample demonstrates testing complex query building
func TestComplexQueryExample(t *testing.T) {
	server := NewTestServer()
	defer server.Close()

	// Capture the request to verify query parameters
	server.ResponseHandler = func(req *http.Request) (int, string) {
		query := req.URL.Query()

		// Verify expected parameters are present
		if query.Get("search") != "machine learning" {
			t.Errorf("Expected search parameter")
		}

		if query.Get("per-page") != "100" {
			t.Errorf("Expected per-page parameter")
		}

		filter := query.Get("filter")
		if filter == "" {
			t.Error("Expected filter parameter")
		}

		return http.StatusOK, SamplePaginatedResponse
	}

	client := goalex.NewClient()
	client.BaseURL = server.URL

	// Execute a complex query
	_, err := client.Works().
		Search("machine learning").
		Filter("publication_year", 2020).
		Filter("type", "article").
		Sort("cited_by_count", true).
		PerPage(100).
		List()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
