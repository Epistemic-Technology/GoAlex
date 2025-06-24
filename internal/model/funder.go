package model

type Funder struct {
	AlternateTitles   []string        `json:"alternate_titles,omitempty"`
	CitedByCount      int             `json:"cited_by_count,omitempty"`
	CountryCode       string          `json:"country_code,omitempty"`
	CountsByYear      []*CountsByYear `json:"counts_by_year,omitempty"`
	CreatedDate       string          `json:"created_date,omitempty"`
	Description       string          `json:"description,omitempty"`
	DisplayName       string          `json:"display_name,omitempty"`
	GrantsCount       int             `json:"grants_count,omitempty"`
	HomePageURL       string          `json:"homepage_url,omitempty"`
	ID                string          `json:"id,omitempty"`
	IDs               *FunderIDs      `json:"ids,omitempty"`
	ImageThumbnailURL string          `json:"image_thumbnail_url,omitempty"`
	ImageURL          string          `json:"image_url,omitempty"`
	Roles             []*Role         `json:"roles,omitempty"`
	SummaryStats      *SummaryStats   `json:"summary_stats,omitempty"`
	UpdatedDate       string          `json:"updated_date,omitempty"`
	WorksCount        int             `json:"works_count,omitempty"`
}

type FunderIDs struct {
	Crossref string `json:"crossref,omitempty"`
	DOI      string `json:"doi,omitempty"`
	OpenAlex string `json:"openalex,omitempty"`
	ROR      string `json:"ror,omitempty"`
	Wikidata string `json:"wikidata,omitempty"`
}
