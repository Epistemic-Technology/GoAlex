package core

import (
	"github.com/Sunhill666/goalex/internal/model"
	"maps"
)

type QueryBuilder[T any] struct {
	client   *Client
	endpoint string
	params   *QueryParams
}

func (q *QueryBuilder[T]) Page(p int) *QueryBuilder[T] {
	if q.params.Pagination == nil {
		q.params.Pagination = &PaginationParams{}
	}
	q.params.Pagination.Page = p
	return q
}

func (q *QueryBuilder[T]) PerPage(pp int) *QueryBuilder[T] {
	if q.params.Pagination == nil {
		q.params.Pagination = &PaginationParams{}
	}
	q.params.Pagination.PerPage = pp
	return q
}

func (q *QueryBuilder[T]) Filter(field string, value any) *QueryBuilder[T] {
	if q.params.Filter == nil {
		q.params.Filter = make(map[string]any)
	}
	q.params.Filter[field] = value
	return q
}

func (q *QueryBuilder[T]) FilterMap(filters map[string]any) *QueryBuilder[T] {
	if q.params.Filter == nil {
		q.params.Filter = make(map[string]any)
	}
	maps.Copy(q.params.Filter, filters)
	return q
}

func (q *QueryBuilder[T]) Search(query string) *QueryBuilder[T] {
	q.params.Search = query
	return q
}

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

func (q *QueryBuilder[T]) Sort(field string, desc bool) *QueryBuilder[T] {
	if q.params.Sort == nil {
		q.params.Sort = make(map[string]bool)
	}
	q.params.Sort[field] = desc
	return q
}

func (q *QueryBuilder[T]) SortMap(sort map[string]bool) *QueryBuilder[T] {
	if q.params.Sort == nil {
		q.params.Sort = make(map[string]bool)
	}
	maps.Copy(q.params.Sort, sort)
	return q
}

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

func (q *QueryBuilder[T]) Sample(sample int) *QueryBuilder[T] {
	if sample <= 0 {
		return q
	} else {
		q.params.Sample = sample
	}
	return q
}

func (q *QueryBuilder[T]) Seed(seed int) *QueryBuilder[T] {
	q.params.Seed = seed
	return q
}

func (q *QueryBuilder[T]) Get(id string) (*T, error) {
	return GetEntity[T](q.client, q.endpoint, id)
}

func (q *QueryBuilder[T]) List() ([]T, error) {
	resp, err := ListEntities[T](q.client, q.endpoint, q.params)
	if err != nil {
		return nil, err
	}
	return resp.Results, nil
}

func (q *QueryBuilder[T]) ListWithMeta() (*model.PaginatedResponse[T], error) {
	return ListEntities[T](q.client, q.endpoint, q.params)
}
