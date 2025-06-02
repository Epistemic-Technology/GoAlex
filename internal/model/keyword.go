package model

type Keyword struct {
	DehydratedKeyword
	CitedByCount int    `json:"cited_by_count"`
	CreatedDate  string `json:"created_date"`
	UpdatedDate  string `json:"updated_date"`
	WorksCount   int    `json:"works_count"`
}

type DehydratedKeyword struct {
	DisplayName string  `json:"display_name"`
	ID          string  `json:"id"`
	Score       float64 `json:"score"`
}
