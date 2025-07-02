package tests

import (
	"net/http"
	"testing"
	"time"

	"github.com/Sunhill666/goalex/pkg/core"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name     string
		options  []core.Option
		expected func(*core.Client) bool
	}{
		{
			name:    "default client",
			options: nil,
			expected: func(c *core.Client) bool {
				return c.BaseURL == "https://api.openalex.org" &&
					c.Timeout == 10*time.Second &&
					c.MaxRetries == 3
			},
		},
		{
			name:    "client with polite pool",
			options: []core.Option{core.PolitePool("test@example.com")},
			expected: func(c *core.Client) bool {
				return c.MailTo == "test@example.com"
			},
		},
		{
			name:    "client with auth",
			options: []core.Option{core.Auth("test-token")},
			expected: func(c *core.Client) bool {
				return c.Token == "test-token"
			},
		},
		{
			name:    "client with timeout",
			options: []core.Option{core.WithTimeout(5 * time.Second)},
			expected: func(c *core.Client) bool {
				return c.Timeout == 5*time.Second
			},
		},
		{
			name:    "client with retry settings",
			options: []core.Option{core.WithRetry(5, 2*time.Second)},
			expected: func(c *core.Client) bool {
				return c.MaxRetries == 5 && c.RetryDelay == 2*time.Second
			},
		},
		{
			name: "client with custom HTTP client",
			options: []core.Option{core.WithHTTPClient(&http.Client{
				Timeout: 3 * time.Second,
			})},
			expected: func(c *core.Client) bool {
				return c.Timeout == 3*time.Second
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := core.New(tt.options...)
			if !tt.expected(client) {
				t.Errorf("Client configuration does not match expected values")
			}
		})
	}
}

func TestClientGet(t *testing.T) {
	server := NewTestServer()
	defer server.Close()

	tests := []struct {
		name           string
		path           string
		serverResponse func() (int, string)
		expectError    bool
		validateResult func(interface{}) bool
	}{
		{
			name: "successful request",
			path: "/test",
			serverResponse: func() (int, string) {
				return http.StatusOK, `{"message": "success", "value": 123}`
			},
			expectError: false,
			validateResult: func(result interface{}) bool {
				if data, ok := result.(*map[string]interface{}); ok {
					return (*data)["message"] == "success"
				}
				return false
			},
		},
		{
			name: "server error with retry",
			path: "/error",
			serverResponse: func() (int, string) {
				return http.StatusInternalServerError, `{"error": "internal server error"}`
			},
			expectError: true,
			validateResult: func(result interface{}) bool {
				return true // Error case, result doesn't matter
			},
		},
		{
			name: "not found error",
			path: "/notfound",
			serverResponse: func() (int, string) {
				return http.StatusNotFound, `{"error": "not found"}`
			},
			expectError: true,
			validateResult: func(result interface{}) bool {
				return true // Error case, result doesn't matter
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewTestClient(server.URL)
			server.SetResponse(tt.serverResponse())

			var result map[string]interface{}
			err := client.Get(tt.path, &result)

			if tt.expectError && err == nil {
				t.Error("Expected error but got nil")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
			if !tt.expectError && !tt.validateResult(&result) {
				t.Error("Result validation failed")
			}
		})
	}
}

func TestClientWithMailTo(t *testing.T) {
	server := NewTestServer()
	defer server.Close()

	client := NewTestClient(server.URL, core.PolitePool("test@example.com"))

	// Capture the request to verify mailto parameter
	server.ResponseHandler = func(req *http.Request) (int, string) {
		if req.URL.Query().Get("mailto") != "test@example.com" {
			t.Error("Expected mailto parameter in query string")
		}
		return http.StatusOK, `{"success": true}`
	}

	var result map[string]interface{}
	err := client.Get("/test", &result)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestClientWithToken(t *testing.T) {
	server := NewTestServer()
	defer server.Close()

	client := NewTestClient(server.URL, core.Auth("test-token"))

	// Capture the request to verify api_key parameter
	server.ResponseHandler = func(req *http.Request) (int, string) {
		if req.URL.Query().Get("api_key") != "test-token" {
			t.Error("Expected api_key parameter in query string")
		}
		return http.StatusOK, `{"success": true}`
	}

	var result map[string]interface{}
	err := client.Get("/test", &result)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
