// Package goalex provides a Go client for the OpenAlex API.
// OpenAlex is a free and open catalog of scholarly papers, authors, venues, and more.
package goalex

import (
	"github.com/Sunhill666/goalex/internal/model"
	"github.com/Sunhill666/goalex/pkg/core"
)

// Client represents an HTTP client for interacting with the OpenAlex API.
type Client = core.Client

// Work represents an individual paper in the OpenAlex API.
type Work = model.Work

// PolitePool configures the client to use a polite pool with the provided email address.
var PolitePool = core.PolitePool

// Auth configures the client to use the provided API token for authentication.
var Auth = core.Auth

// WithTimeout configures the client's timeout duration.
var WithTimeout = core.WithTimeout

// WithRetry configures the client's retry behavior with maximum retry attempts and delay.
var WithRetry = core.WithRetry

// WithHTTPClient configures the client to use a custom HTTP client.
var WithHTTPClient = core.WithHTTPClient

// NewClient creates a new Client with the provided options.
func NewClient(opts ...core.Option) *Client {
	return core.New(opts...)
}
