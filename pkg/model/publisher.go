package model

type Publisher struct {
	AlternateTitles []string       `json:"alternate_titles"`
	CitedByCount    int            `json:"cited_by_count"`
	CountryCode     string         `json:"country_code"`
	CountsByYear    []CountsByYear `json:"counts_by_year"`
	CreatedDate     string         `json:"created_date"`
	DisplayName     string         `json:"display_name"`
	HierarchyLevel  int            `json:"hierarchy_level"`
	ID              string         `json:"id"`
	IDs             struct {
		OpenAlex  string `json:"openalex"`
		ROR       string `json:"ror"`
		Wikipedia string `json:"wikipedia"`
	} `json:"ids"`
	ImageThumbnailURL string   `json:"image_thumbnail_url"`
	ImageURL          string   `json:"image_url"`
	Lineage           []string `json:"lineage"`
	ParentPublisher   struct {
		ID          string `json:"id"`
		DisplayName string `json:"display_name"`
	} `json:"parent_publisher"`
	Roles        []Role       `json:"roles"`
	SourceAPIURL string       `json:"source_api_url"`
	SummaryStats SummaryStats `json:"summary_stats"`
	UpdatedDate  string       `json:"updated_date"`
	WorksCount   int          `json:"works_count"`
}
