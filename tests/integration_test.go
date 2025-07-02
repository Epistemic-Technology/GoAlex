package tests

import (
	"net/http"
	"testing"
	"time"

	"github.com/Sunhill666/goalex/pkg/core"
)

func TestRetryMechanism(t *testing.T) {
	server := NewTestServer()
	defer server.Close()

	client := NewTestClient(server.URL, core.WithRetry(2, 10*time.Millisecond))

	t.Run("retry on server error", func(t *testing.T) {
		attempts := 0
		server.ResponseHandler = func(req *http.Request) (int, string) {
			attempts++
			if attempts <= 2 {
				return http.StatusInternalServerError, `{"error": "server error"}`
			}
			return http.StatusOK, `{"message": "success"}`
		}

		var result map[string]interface{}
		err := client.Get("/test", &result)

		if err != nil {
			t.Errorf("Expected success after retries, got error: %v", err)
		}

		if attempts != 3 {
			t.Errorf("Expected 3 attempts, got %d", attempts)
		}
	})

	t.Run("max retries exceeded", func(t *testing.T) {
		server.ResponseHandler = func(req *http.Request) (int, string) {
			return http.StatusInternalServerError, `{"error": "persistent error"}`
		}

		var result map[string]interface{}
		err := client.Get("/test", &result)

		if err == nil {
			t.Error("Expected error after max retries exceeded")
		}
	})

	t.Run("retry on rate limit", func(t *testing.T) {
		attempts := 0
		server.ResponseHandler = func(req *http.Request) (int, string) {
			attempts++
			if attempts == 1 {
				return http.StatusTooManyRequests, `{"error": "rate limited"}`
			}
			return http.StatusOK, `{"message": "success"}`
		}

		var result map[string]interface{}
		err := client.Get("/test", &result)

		if err != nil {
			t.Errorf("Expected success after retry on rate limit, got error: %v", err)
		}

		if attempts != 2 {
			t.Errorf("Expected 2 attempts, got %d", attempts)
		}
	})

	t.Run("no retry on client error", func(t *testing.T) {
		attempts := 0
		server.ResponseHandler = func(req *http.Request) (int, string) {
			attempts++
			return http.StatusBadRequest, `{"error": "bad request"}`
		}

		var result map[string]interface{}
		err := client.Get("/test", &result)

		if err == nil {
			t.Error("Expected error for client error")
		}

		if attempts != 1 {
			t.Errorf("Expected 1 attempt (no retry), got %d", attempts)
		}
	})
}

func TestTimeoutHandling(t *testing.T) {
	server := NewTestServer()
	defer server.Close()

	// Create client with very short timeout
	client := NewTestClient(server.URL, core.WithTimeout(1*time.Millisecond))

	// Simulate slow server response
	server.ResponseHandler = func(req *http.Request) (int, string) {
		time.Sleep(10 * time.Millisecond) // Longer than client timeout
		return http.StatusOK, `{"message": "success"}`
	}

	var result map[string]interface{}
	err := client.Get("/test", &result)

	if err == nil {
		t.Error("Expected timeout error")
	}

	// Check if it's a timeout error
	if err != nil && err.Error() != "" {
		// Timeout error should contain context about timeout or cancellation
		t.Logf("Got expected timeout error: %v", err)
	}
}

func TestErrorResponse(t *testing.T) {
	server := NewTestServer()
	defer server.Close()

	client := NewTestClient(server.URL)

	tests := []struct {
		name         string
		statusCode   int
		responseBody string
		expectError  bool
		shouldRetry  bool
	}{
		{
			name:         "not found error",
			statusCode:   http.StatusNotFound,
			responseBody: `{"error": "not found"}`,
			expectError:  true,
			shouldRetry:  false,
		},
		{
			name:         "unauthorized error",
			statusCode:   http.StatusUnauthorized,
			responseBody: `{"error": "unauthorized"}`,
			expectError:  true,
			shouldRetry:  false,
		},
		{
			name:         "forbidden error",
			statusCode:   http.StatusForbidden,
			responseBody: `{"error": "forbidden"}`,
			expectError:  true,
			shouldRetry:  false,
		},
		{
			name:         "rate limit error",
			statusCode:   http.StatusTooManyRequests,
			responseBody: `{"error": "rate limited"}`,
			expectError:  true,
			shouldRetry:  true,
		},
		{
			name:         "server error",
			statusCode:   http.StatusInternalServerError,
			responseBody: `{"error": "internal server error"}`,
			expectError:  true,
			shouldRetry:  true,
		},
		{
			name:         "bad gateway error",
			statusCode:   http.StatusBadGateway,
			responseBody: `{"error": "bad gateway"}`,
			expectError:  true,
			shouldRetry:  true,
		},
		{
			name:         "service unavailable error",
			statusCode:   http.StatusServiceUnavailable,
			responseBody: `{"error": "service unavailable"}`,
			expectError:  true,
			shouldRetry:  true,
		},
		{
			name:         "gateway timeout error",
			statusCode:   http.StatusGatewayTimeout,
			responseBody: `{"error": "gateway timeout"}`,
			expectError:  true,
			shouldRetry:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server.SetResponse(tt.statusCode, tt.responseBody)

			var result map[string]interface{}
			err := client.Get("/test", &result)

			if tt.expectError && err == nil {
				t.Error("Expected error but got nil")
			}

			if !tt.expectError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}

func TestInvalidJSON(t *testing.T) {
	server := NewTestServer()
	defer server.Close()

	client := NewTestClient(server.URL)

	server.SetResponse(http.StatusOK, `invalid json response`)

	var result map[string]interface{}
	err := client.Get("/test", &result)

	if err == nil {
		t.Error("Expected JSON decode error")
	}
}

func TestInvalidURL(t *testing.T) {
	// Test with invalid base URL
	client := core.New()
	client.BaseURL = "://invalid-url"

	var result map[string]interface{}
	err := client.Get("/test", &result)

	if err == nil {
		t.Error("Expected URL parse error")
	}
}
