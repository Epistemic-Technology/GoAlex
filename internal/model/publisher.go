package model

type Publisher struct {
	AlternateTitles   []string         `json:"alternate_titles,omitempty"`
	CitedByCount      int              `json:"cited_by_count,omitempty"`
	CountryCode       string           `json:"country_code,omitempty"`
	CountsByYear      []*CountsByYear  `json:"counts_by_year,omitempty"`
	CreatedDate       string           `json:"created_date,omitempty"`
	DisplayName       string           `json:"display_name,omitempty"`
	HierarchyLevel    int              `json:"hierarchy_level,omitempty"`
	ID                string           `json:"id,omitempty"`
	IDs               *PublisherIDs    `json:"ids,omitempty"`
	ImageThumbnailURL string           `json:"image_thumbnail_url,omitempty"`
	ImageURL          string           `json:"image_url,omitempty"`
	Lineage           []string         `json:"lineage,omitempty"`
	ParentPublisher   *ParentPublisher `json:"parent_publisher,omitempty"`
	Roles             []*Role          `json:"roles,omitempty"`
	SourceAPIURL      string           `json:"source_api_url,omitempty"`
	SummaryStats      *SummaryStats    `json:"summary_stats,omitempty"`
	UpdatedDate       string           `json:"updated_date,omitempty"`
	WorksCount        int              `json:"works_count,omitempty"`
}

type PublisherIDs struct {
	OpenAlex  string `json:"openalex,omitempty"`
	ROR       string `json:"ror,omitempty"`
	Wikipedia string `json:"wikipedia,omitempty"`
}

type ParentPublisher struct {
	ID          string `json:"id,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
}
