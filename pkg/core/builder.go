package core

import (
	"maps"

	"github.com/Sunhill666/goalex/internal/model"
)

// QueryBuilder provides a fluent interface for building queries to the OpenAlex API.
type QueryBuilder[T any] struct {
	client   *Client
	endpoint string
	params   *QueryParams
}

// Page sets the page number for pagination.
func (q *QueryBuilder[T]) Page(p int) *QueryBuilder[T] {
	if q.params.Pagination == nil {
		q.params.Pagination = &PaginationParams{}
	}
	q.params.Pagination.Page = p
	return q
}

// PerPage sets the number of results per page for pagination.
func (q *QueryBuilder[T]) PerPage(pp int) *QueryBuilder[T] {
	if q.params.Pagination == nil {
		q.params.Pagination = &PaginationParams{}
	}
	q.params.Pagination.PerPage = pp
	return q
}

// Filter adds a filter parameter to the query.
func (q *QueryBuilder[T]) Filter(field string, value any) *QueryBuilder[T] {
	if q.params.Filter == nil {
		q.params.Filter = make(map[string]any)
	}
	q.params.Filter[field] = value
	return q
}

// FilterMap adds multiple filter parameters to the query.
func (q *QueryBuilder[T]) FilterMap(filters map[string]any) *QueryBuilder[T] {
	if q.params.Filter == nil {
		q.params.Filter = make(map[string]any)
	}
	maps.Copy(q.params.Filter, filters)
	return q
}

// Search sets the search query parameter.
func (q *QueryBuilder[T]) Search(query string) *QueryBuilder[T] {
	q.params.Search = query
	return q
}

// SearchFilter adds search filters with optional no-stem option to the query.
func (q *QueryBuilder[T]) SearchFilter(search_filters map[string]string, no_stem bool) *QueryBuilder[T] {
	if q.params.Filter == nil {
		q.params.Filter = make(map[string]any)
	}
	for k, v := range search_filters {
		newKey := k + ".search"
		if no_stem {
			newKey += ".no_stem"
		}
		q.params.Filter[newKey] = v
	}
	return q
}

// Sort adds a sort parameter to the query.
func (q *QueryBuilder[T]) Sort(field string, desc bool) *QueryBuilder[T] {
	if q.params.Sort == nil {
		q.params.Sort = make(map[string]bool)
	}
	q.params.Sort[field] = desc
	return q
}

// SortMap adds multiple sort parameters to the query.
func (q *QueryBuilder[T]) SortMap(sort map[string]bool) *QueryBuilder[T] {
	if q.params.Sort == nil {
		q.params.Sort = make(map[string]bool)
	}
	maps.Copy(q.params.Sort, sort)
	return q
}

// Select adds field selection parameters to the query.
func (q *QueryBuilder[T]) Select(fields ...string) *QueryBuilder[T] {
	if len(fields) == 0 {
		return q
	}
	if q.params.Select == nil {
		q.params.Select = make([]string, 0)
	}
	q.params.Select = append(q.params.Select, fields...)
	return q
}

// Sample sets the sample size for random sampling.
func (q *QueryBuilder[T]) Sample(sample int) *QueryBuilder[T] {
	if sample <= 0 {
		return q
	} else {
		q.params.Sample = sample
	}
	return q
}

// Seed sets the random seed for reproducible sampling.
func (q *QueryBuilder[T]) Seed(seed int) *QueryBuilder[T] {
	q.params.Seed = seed
	return q
}

// Get retrieves a single entity by its ID.
func (q *QueryBuilder[T]) Get(id string) (*T, error) {
	return GetEntity[T](q.client, q.endpoint, id)
}

// GetRandom retrieves a random entity.
func (q *QueryBuilder[T]) GetRandom() (*T, error) {
	return GetEntity[T](q.client, q.endpoint, "random")
}

// GroupBy adds a group by parameter to the query with optional inclusion of unknown values.
func (q *QueryBuilder[T]) GroupBy(field string, includeUnknown bool) *QueryBuilder[T] {
	q.params.GroupBy = field
	if includeUnknown {
		q.params.GroupBy += ":include_unknown"
	}
	return q
}

// AutoComplete creates a new query builder for autocomplete suggestions.
func (q *QueryBuilder[T]) AutoComplete(query string) *QueryBuilder[model.Completion] {
	autoCompleteBuilder := &QueryBuilder[model.Completion]{
		client:   q.client,
		endpoint: EndPointAutoComplete + q.endpoint,
		params:   q.params,
	}
	autoCompleteBuilder.params.AutoComplete = query
	return autoCompleteBuilder
}

// List executes the query and returns a list of entities.
func (q *QueryBuilder[T]) List() ([]*T, error) {
	resp, err := ListEntities[T](q.client, q.endpoint, q.params)
	if err != nil {
		return nil, err
	}
	return resp.Results, nil
}

// ListGroupBy executes the query and returns grouped results.
func (q *QueryBuilder[T]) ListGroupBy() ([]*model.GroupBy, error) {
	resp, err := ListEntities[T](q.client, q.endpoint, q.params)
	if err != nil {
		return nil, err
	}
	return resp.GroupBy, nil
}

// Cursor executes the query using cursor-based pagination and returns results with next cursor.
func (q *QueryBuilder[T]) Cursor(cursor ...string) ([]*T, string, error) {
	if len(cursor) > 0 {
		q.params.Cursor = cursor[0]
	} else {
		q.params.Cursor = "*"
	}
	resp, err := ListEntities[T](q.client, q.endpoint, q.params)
	if err != nil {
		return nil, "", err
	}
	return resp.Results, resp.Meta.NextCursor, nil
}

// ListWithMeta executes the query and returns results with metadata.
func (q *QueryBuilder[T]) ListWithMeta() (*model.PaginatedResponse[T], error) {
	return ListEntities[T](q.client, q.endpoint, q.params)
}
