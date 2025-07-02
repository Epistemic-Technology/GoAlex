package model

type Topic struct {
	Description string      `json:"description,omitempty"`
	DisplayName string      `json:"display_name,omitempty"`
	Domain      *TopicField `json:"domain,omitempty"`
	Field       *TopicField `json:"field,omitempty"`
	ID          string      `json:"id,omitempty"`
	IDs         *TopicIDs   `json:"ids,omitempty"`
	Keywords    []string    `json:"keywords,omitempty"`
	Subfield    *TopicField `json:"subfield,omitempty"`
	UpdateDate  string      `json:"update_date,omitempty"`
	WorksCount  int         `json:"works_count,omitempty"`
}

type TopicIDs struct {
	OpenAlex  string `json:"openalex,omitempty"`
	Wikipedia string `json:"wikipedia,omitempty"`
}

type TopicWithCount struct {
	Topic
	Count int `json:"count,omitempty"`
}

type TopicWithScore struct {
	Topic
	Score float32 `json:"score,omitempty"`
}

type TopicShare struct {
	Topic
	Value float64 `json:"value,omitempty"`
}

type TopicField struct {
	ID          string `json:"id,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
}
