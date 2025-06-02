package model

type Topic struct {
	Description string     `json:"description"`
	DisplayName string     `json:"display_name"`
	Domain      TopicField `json:"domain"`
	Field       TopicField `json:"field"`
	ID          string     `json:"id"`
	IDs         struct {
		OpenAlex  string `json:"openalex"`
		Wikipedia string `json:"wikipedia"`
	} `json:"ids"`
	Keywords   []string   `json:"keywords"`
	Subfield   TopicField `json:"subfield"`
	UpdateDate string     `json:"update_date"`
	WorksCount int        `json:"works_count"`
}

type TopicWithCount struct {
	Topic
	Count int `json:"count"`
}

type TopicWithScore struct {
	Topic
	Score float32 `json:"score"`
}

type TopicShare struct {
	Topic
	Value float64 `json:"value"`
}

type TopicField struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
}
