package model

type Keyword struct {
	DehydratedKeyword
	CitedByCount int    `json:"cited_by_count,omitempty"`
	CreatedDate  string `json:"created_date,omitempty"`
	UpdatedDate  string `json:"updated_date,omitempty"`
	WorksCount   int    `json:"works_count,omitempty"`
}

type DehydratedKeyword struct {
	DisplayName string  `json:"display_name,omitempty"`
	ID          string  `json:"id,omitempty"`
	Score       float64 `json:"score,omitempty"`
}
