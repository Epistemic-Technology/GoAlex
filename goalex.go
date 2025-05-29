package goalex

import (
	"github.com/Sunhill666/goalex/pkg/client"
	"github.com/Sunhill666/goalex/pkg/service"
)

type QueryParams = service.QueryParams
type PaginationParams = service.PaginationParams

var WithMailto = client.WithMailto

type Client struct {
	works   *service.WorksService
}

func NewClient(opts ...client.Option) *Client {
	c := client.New(opts...)
	return &Client{
		works:   service.NewWorksService(c),
	}
}

func (c *Client) Works() *service.WorksService {
	return c.works
}
