package core

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Client represents an HTTP client for interacting with the OpenAlex API.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	MailTo     string
	Token      string
	Timeout    time.Duration
	MaxRetries int
	RetryDelay time.Duration
}

// Option is a function type for configuring the Client.
type Option func(*Client)

// PolitePool configures the client to use a polite pool with the provided email address.
func PolitePool(email string) Option {
	return func(c *Client) {
		c.MailTo = email
	}
}

// Auth configures the client to use the provided API token for authentication.
func Auth(token string) Option {
	return func(c *Client) {
		c.Token = token
	}
}

// WithTimeout configures the client's timeout duration.
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.Timeout = timeout
		c.HTTPClient.Timeout = timeout
	}
}

// WithRetry configures the client's retry behavior with maximum retry attempts and delay.
func WithRetry(maxRetries int, retryDelay time.Duration) Option {
	return func(c *Client) {
		c.MaxRetries = maxRetries
		c.RetryDelay = retryDelay
	}
}

// WithHTTPClient configures the client to use a custom HTTP client.
func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.HTTPClient = client
		if client.Timeout > 0 {
			c.Timeout = client.Timeout
		}
	}
}

// New creates a new Client with the provided options.
func New(opts ...Option) *Client {
	c := &Client{
		BaseURL:    "https://api.openalex.org",
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		Timeout:    10 * time.Second,
		MaxRetries: 3,
		RetryDelay: time.Second,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func isRetryableError(err error) bool {
	if err == nil {
		return false
	}
	if urlErr, ok := err.(*url.Error); ok {
		return urlErr.Temporary() || urlErr.Timeout()
	}
	return false
}

func isRetryableStatusCode(statusCode int) bool {
	switch statusCode {
	case http.StatusTooManyRequests, // 429
		http.StatusInternalServerError, // 500
		http.StatusBadGateway,          // 502
		http.StatusServiceUnavailable,  // 503
		http.StatusGatewayTimeout:      // 504
		return true
	default:
		return false
	}
}

// Get performs a GET request to the specified path and decodes the response into out.
func (c *Client) Get(path string, out any) error {
	return c.GetWithContext(context.Background(), path, out)
}

// GetWithContext performs a GET request to the specified path with context support
// and decodes the response into out.
func (c *Client) GetWithContext(ctx context.Context, path string, out any) error {
	base, err := url.Parse(c.BaseURL)
	if err != nil {
		return fmt.Errorf("invalid base URL: %w", err)
	}

	rel, err := url.Parse(path)
	if err != nil {
		return fmt.Errorf("invalid path: %w", err)
	}

	u := base.ResolveReference(rel)

	q := u.Query()
	if c.MailTo != "" {
		q.Set("mailto", c.MailTo)
	}
	if c.Token != "" {
		q.Set("api_key", c.Token)
	}

	u.RawQuery = q.Encode()

	var lastErr error
	for attempt := 0; attempt <= c.MaxRetries; attempt++ {
		if attempt > 0 {
			// If not the first attempt, wait for a while before retrying
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(c.RetryDelay * time.Duration(attempt)):
			}
		}

		req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}

		resp, err := c.HTTPClient.Do(req)
		if err != nil {
			lastErr = err
			if isRetryableError(err) && attempt < c.MaxRetries {
				continue
			}
			return fmt.Errorf("request failed after %d attempts: %w", attempt+1, err)
		}

		if resp.StatusCode >= 400 {
			_ = resp.Body.Close()
			if isRetryableStatusCode(resp.StatusCode) && attempt < c.MaxRetries {
				lastErr = fmt.Errorf("HTTP %d: %s", resp.StatusCode, resp.Status)
				continue
			}
			return fmt.Errorf("HTTP %d: %s", resp.StatusCode, resp.Status)
		}

		defer func() { _ = resp.Body.Close() }()
		err = json.NewDecoder(resp.Body).Decode(out)
		if err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}

		return nil
	}

	return fmt.Errorf("request failed after %d attempts, last error: %w", c.MaxRetries+1, lastErr)
}
