package core

import (
	"github.com/Sunhill666/goalex/internal/model"
)

const (
	EndpointWorks        = "/works"
	EndpointAuthors      = "/authors"
	EndpointSources      = "/sources"
	EndpointInstitutions = "/institutions"
	EndpointTopics       = "/topics"
	EndpointKeywords     = "/keywords"
	EndpointPublishers   = "/publishers"
	EndpointFunders      = "/funders"
	EndpointConcepts     = "/concepts"
	EndPointAutoComplete = "/autocomplete"
)

func Query[T any](c *Client, endpoint string) *QueryBuilder[T] {
	return &QueryBuilder[T]{
		client:   c,
		endpoint: endpoint,
		params:   &QueryParams{},
	}
}

func (c *Client) Works() *QueryBuilder[model.Work] {
	return Query[model.Work](c, EndpointWorks)
}

func (c *Client) Authors() *QueryBuilder[model.Author] {
	return Query[model.Author](c, EndpointAuthors)
}

func (c *Client) Sources() *QueryBuilder[model.Source] {
	return Query[model.Source](c, EndpointSources)
}

func (c *Client) Institutions() *QueryBuilder[model.Institution] {
	return Query[model.Institution](c, EndpointInstitutions)
}

func (c *Client) Topics() *QueryBuilder[model.Topic] {
	return Query[model.Topic](c, EndpointTopics)
}

func (c *Client) Keywords() *QueryBuilder[model.Keyword] {
	return Query[model.Keyword](c, EndpointKeywords)
}

func (c *Client) Publishers() *QueryBuilder[model.Publisher] {
	return Query[model.Publisher](c, EndpointPublishers)
}

func (c *Client) Funders() *QueryBuilder[model.Funder] {
	return Query[model.Funder](c, EndpointFunders)
}

func (c *Client) Concepts() *QueryBuilder[model.Concept] {
	return Query[model.Concept](c, EndpointConcepts)
}
