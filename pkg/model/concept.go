package model

type DehydratedConcept struct {
	DisplayName string `json:"display_name"`
	ID          string `json:"id"`
	Level       int    `json:"level"`
	Wikipedia   string `json:"wikipedia"`
}

type DehydratedConceptWithScore struct {
	DehydratedConcept
	Score float32 `json:"score"`
}
