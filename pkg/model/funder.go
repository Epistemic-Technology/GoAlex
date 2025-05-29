package model

type Funder struct {
	AlternateTitles []string       `json:"alternate_titles"`
	CitedByCount    int            `json:"cited_by_count"`
	CountryCode     string         `json:"country_code"`
	CountsByYear    []CountsByYear `json:"counts_by_year"`
	CreatedDate     string         `json:"created_date"`
	Description     string         `json:"description"`
	DisplayName     string         `json:"display_name"`
	GrantsCount     int            `json:"grants_count"`
	HomePageURL     string         `json:"homepage_url"`
	ID              string         `json:"id"`
	IDs             struct {
		Crossref string `json:"crossref"`
		DOI      string `json:"doi"`
		OpenAlex string `json:"openalex"`
		ROR      string `json:"ror"`
		Wikidata string `json:"wikidata"`
	} `json:"ids"`
	ImageThumbnailURL string       `json:"image_thumbnail_url"`
	ImageURL          string       `json:"image_url"`
	Roles             []Role       `json:"roles"`
	SummaryStats      SummaryStats `json:"summary_stats"`
	UpdatedDate       string       `json:"updated_date"`
	WorksCount        int          `json:"works_count"`
}
