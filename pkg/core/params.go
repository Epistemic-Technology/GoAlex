package core

import (
	"fmt"
	"maps"
	"net/url"
	"strings"
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
	Filter     map[string]any
	Search     string
	Sort       map[string]bool
	Select     []string
}

func (q *QueryParams) ToQuery() url.Values {
	query := url.Values{}
	if q.Pagination != nil {
		paginationQuery := q.Pagination.ToQuery()
		maps.Copy(query, paginationQuery)
	}
	if q.Filter != nil {
		var sb strings.Builder
		first := true
		for k, v := range q.Filter {
			if !first {
				sb.WriteString(",")
			}
			sb.WriteString(fmt.Sprintf("%s:%s", k, fmt.Sprint(v)))
			first = false
		}
		query.Set("filter", sb.String())
	}
	if q.Search != "" {
		query.Set("search", q.Search)
	}
	if q.Sort != nil {
		var sb strings.Builder
		first := true
		for k, v := range q.Sort {
			if !first {
				sb.WriteString(",")
			}
			if v {
				sb.WriteString(fmt.Sprintf("%s:desc", k))
			} else {
				sb.WriteString(k)
			}
			first = false
		}
		query.Set("sort", sb.String())
	}
	if q.Select != nil {
		query.Set("select", strings.Join(q.Select, ","))
	}
	return query
}
