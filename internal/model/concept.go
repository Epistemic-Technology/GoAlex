package model

type Concept struct {
	Ancestors         []*DehydratedConcept          `json:"ancestors,omitempty"`
	CitedByCount      int                           `json:"cited_by_count,omitempty"`
	CountsByYear      []*CountsByYear               `json:"counts_by_year,omitempty"`
	CreatedDate       string                        `json:"created_date,omitempty"`
	Description       string                        `json:"description,omitempty"`
	IDs               *ConceptIDs                   `json:"ids,omitempty"`
	ImageThumbnailURL string                        `json:"image_thumbnail_url,omitempty"`
	ImageURL          string                        `json:"image_url,omitempty"`
	International     *International                `json:"international,omitempty"`
	Level             int                           `json:"level,omitempty"`
	RelatedConcepts   []*DehydratedConceptWithScore `json:"related_concepts,omitempty"`
	SummaryStats      *SummaryStats                 `json:"summary_stats,omitempty"`
	UpdatedDate       string                        `json:"updated_date,omitempty"`
	Wikidata          string                        `json:"wikidata,omitempty"`
	WorksAPIURL       string                        `json:"works_api_url,omitempty"`
	WorksCount        int                           `json:"works_count,omitempty"`
}

type ConceptIDs struct {
	MAG       string   `json:"mag,omitempty"`
	OpenAlex  string   `json:"openalex,omitempty"`
	UMLSCUI   []string `json:"umls_cui,omitempty"`
	UMLSAUI   []string `json:"umls_aui,omitempty"`
	Wikidata  string   `json:"wikidata,omitempty"`
	Wikipedia string   `json:"wikipedia,omitempty"`
}

type DehydratedConcept struct {
	DisplayName string `json:"display_name,omitempty"`
	ID          string `json:"id,omitempty"`
	Level       int    `json:"level,omitempty"`
	Wikidata    string `json:"wikidata,omitempty"`
}

type DehydratedConceptWithScore struct {
	DehydratedConcept
	Score float32 `json:"score,omitempty"`
}
