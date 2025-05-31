package model

type PaginatedResponse[T any] struct {
	Meta struct {
		Count       int `json:"count"`
		DBRespTime  int `json:"db_response_time_ms"`
		Page        int `json:"page"`
		PerPage     int `json:"per_page"`
		GroupsCount int `json:"groups_count"`
	} `json:"meta"`
	Results []T `json:"results"`
}
