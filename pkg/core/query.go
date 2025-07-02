package core

import (
	"github.com/Sunhill666/goalex/internal/model"
)

// OpenAlex API endpoint constants
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

// Query creates a new QueryBuilder for the specified endpoint and entity type.
func Query[T any](c *Client, endpoint string) *QueryBuilder[T] {
	return &QueryBuilder[T]{
		client:   c,
		endpoint: endpoint,
		params:   &QueryParams{},
	}
}

// Works returns a QueryBuilder for querying works.
func (c *Client) Works() *QueryBuilder[model.Work] {
	return Query[model.Work](c, EndpointWorks)
}

// Authors returns a QueryBuilder for querying authors.
func (c *Client) Authors() *QueryBuilder[model.Author] {
	return Query[model.Author](c, EndpointAuthors)
}

// Sources returns a QueryBuilder for querying sources.
func (c *Client) Sources() *QueryBuilder[model.Source] {
	return Query[model.Source](c, EndpointSources)
}

// Institutions returns a QueryBuilder for querying institutions.
func (c *Client) Institutions() *QueryBuilder[model.Institution] {
	return Query[model.Institution](c, EndpointInstitutions)
}

// Topics returns a QueryBuilder for querying topics.
func (c *Client) Topics() *QueryBuilder[model.Topic] {
	return Query[model.Topic](c, EndpointTopics)
}

// Keywords returns a QueryBuilder for querying keywords.
func (c *Client) Keywords() *QueryBuilder[model.Keyword] {
	return Query[model.Keyword](c, EndpointKeywords)
}

// Publishers returns a QueryBuilder for querying publishers.
func (c *Client) Publishers() *QueryBuilder[model.Publisher] {
	return Query[model.Publisher](c, EndpointPublishers)
}

// Funders returns a QueryBuilder for querying funders.
func (c *Client) Funders() *QueryBuilder[model.Funder] {
	return Query[model.Funder](c, EndpointFunders)
}

// Concepts returns a QueryBuilder for querying concepts.
func (c *Client) Concepts() *QueryBuilder[model.Concept] {
	return Query[model.Concept](c, EndpointConcepts)
}
