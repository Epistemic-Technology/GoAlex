package tests

import (
	"testing"

	"github.com/Sunhill666/goalex"
)

func TestGoalexPackageExports(t *testing.T) {
	// Test that the main package properly exports core functionality

	t.Run("NewClient function", func(t *testing.T) {
		client := goalex.NewClient()
		if client == nil {
			t.Error("NewClient should return a non-nil client")
		}
	})

	t.Run("NewClient with options", func(t *testing.T) {
		client := goalex.NewClient(
			goalex.PolitePool("test@example.com"),
			goalex.Auth("test-token"),
		)
		if client == nil {
			t.Error("NewClient with options should return a non-nil client")
		}

		// Verify the options were applied
		if client != nil {
			if client.MailTo != "test@example.com" {
				t.Errorf("Expected MailTo to be 'test@example.com', got '%s'", client.MailTo)
			}
			if client.Token != "test-token" {
				t.Errorf("Expected Token to be 'test-token', got '%s'", client.Token)
			}
		}
	})

	t.Run("exported options", func(t *testing.T) {
		// Test that all options are properly exported
		options := []interface{}{
			goalex.PolitePool("test@example.com"),
			goalex.Auth("test-token"),
			goalex.WithTimeout(0),
			goalex.WithRetry(0, 0),
			goalex.WithHTTPClient(nil),
		}

		for i, opt := range options {
			if opt == nil {
				t.Errorf("Option %d should not be nil", i)
			}
		}
	})

	t.Run("client methods", func(t *testing.T) {
		client := goalex.NewClient()

		// Test that all query builders are available
		builders := []interface{}{
			client.Works(),
			client.Authors(),
			client.Sources(),
			client.Institutions(),
			client.Topics(),
			client.Keywords(),
			client.Publishers(),
			client.Funders(),
			client.Concepts(),
		}

		for i, builder := range builders {
			if builder == nil {
				t.Errorf("Builder %d should not be nil", i)
			}
		}
	})
}

// Example usage test that demonstrates the library's intended usage
func TestExampleUsage(t *testing.T) {
	// This test demonstrates how users would typically use the library

	// Create a client with polite pool
	client := goalex.NewClient(goalex.PolitePool("test@example.com"))

	// The client should be ready to use
	if client == nil {
		t.Fatal("Client should not be nil")
	}

	// Query builder should be chainable
	builder := client.Works().
		Filter("publication_year", 2020).
		Filter("is_oa", true).
		Search("machine learning").
		Sort("cited_by_count", true).
		Select("title", "doi", "publication_year").
		Page(1).
		PerPage(50)

	if builder == nil {
		t.Error("Query builder should not be nil after chaining")
	}

	// Note: We don't actually make API calls in this test since we don't have a mock server
	// In real usage, the user would call builder.List() or similar methods
}

func TestGoalexIntegration(t *testing.T) {
	// Test the integration between the main package and core package

	t.Run("client type alias", func(t *testing.T) {
		client := goalex.NewClient()

		// The client should have the expected methods from core.Client
		// This verifies that the type alias is working correctly
		if client.BaseURL == "" {
			t.Error("Client should have BaseURL field from core.Client")
		}

		if client.HTTPClient == nil {
			t.Error("Client should have HTTPClient field from core.Client")
		}
	})
}
