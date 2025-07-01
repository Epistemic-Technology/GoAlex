package model

type PaginatedResponse[T any] struct {
	Meta    *PaginatedResponseMeta `json:"meta,omitempty"`
	Results []*T                   `json:"results,omitempty"`
	GroupBy []*GroupBy             `json:"group_by,omitempty"`
}
