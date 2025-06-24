package model

type PaginatedResponse[T any] struct {
	Meta    *PaginatedResponseMeta `json:"meta,omitempty"`
	Results []*T                   `json:"results,omitempty"`
	GroupBy []*T                   `json:"group_by,omitempty"`
}

type PaginatedResponseMeta struct {
	Count       int `json:"count,omitempty"`
	DBRespTime  int `json:"db_response_time_ms,omitempty"`
	Page        int `json:"page,omitempty"`
	PerPage     int `json:"per_page,omitempty"`
	GroupsCount int `json:"groups_count,omitempty"`
}
