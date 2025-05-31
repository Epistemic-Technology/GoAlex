package core

import (
	"github.com/Sunhill666/goalex/internal/model"
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
