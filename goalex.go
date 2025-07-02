package goalex

import "github.com/Sunhill666/goalex/pkg/core"

type Client = core.Client

var PolitePool = core.PolitePool
var Auth = core.Auth
var WithTimeout = core.WithTimeout
var WithRetry = core.WithRetry
var WithHTTPClient = core.WithHTTPClient

func NewClient(opts ...core.Option) *Client {
	return core.New(opts...)
}
