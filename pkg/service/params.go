package service

import (
	"fmt"
	"net/url"
)

type PaginationParams struct {
	Page    int
	PerPage int
	Cursor  string
}

func (p *PaginationParams) ToQuery() url.Values {
	q := url.Values{}
	if p.Page > 0 {
		q.Set("page", fmt.Sprintf("%d", p.Page))
	}
	if p.PerPage > 0 {
		q.Set("per-page", fmt.Sprintf("%d", p.PerPage))
	}
	if p.Cursor != "" {
		q.Set("cursor", p.Cursor)
	}
	return q
}

type QueryParams struct {
	Pagination *PaginationParams
}

func (q *QueryParams) ToQuery() url.Values {
	query := url.Values{}
	if q.Pagination != nil {
		for k, vs := range q.Pagination.ToQuery() {
			for _, v := range vs {
				query.Add(k, v)
			}
		}
	}
	return query
}
