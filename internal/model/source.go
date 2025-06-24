package model

type Source struct {
	DehydratedSource
	AbbreviatedTitle string                        `json:"abbreviated_title,omitempty"`
	AlternateTitles  []string                      `json:"alternate_titles,omitempty"`
	APCPrice         *APCPrice                     `json:"apc_price,omitempty"`
	APCUSD           int                           `json:"apc_usd,omitempty"`
	CitedByCount     int                           `json:"cited_by_count,omitempty"`
	CountryCode      string                        `json:"country_code,omitempty"`
	CountsByYear     []*CountsByYear               `json:"counts_by_year,omitempty"`
	CreatedDate      string                        `json:"created_date,omitempty"`
	HomePageURL      string                        `json:"homepage_url,omitempty"`
	IDs              *SourceIDs                    `json:"ids,omitempty"`
	Societies        []*Societies                  `json:"societies,omitempty"`
	SummaryStats     *SummaryStats                 `json:"summary_stats,omitempty"`
	Topics           []*TopicWithCount             `json:"topics,omitempty"`
	TopicShare       []*TopicShare                 `json:"topic_share,omitempty"`
	UpdatedDate      string                        `json:"updated_date,omitempty"`
	WorksAPIURL      string                        `json:"works_api_url,omitempty"`
	WorksCount       int                           `json:"works_count,omitempty"`
	XConcepts        []*DehydratedConceptWithScore `json:"x_concepts,omitempty"`
}

type SourceIDs struct {
	Fatcat    string   `json:"fatcat,omitempty"`
	ISSN      []string `json:"issn,omitempty"`
	ISSNL     string   `json:"issn_l,omitempty"`
	MAG       string   `json:"mag,omitempty"`
	OpenAlex  string   `json:"openalex,omitempty"`
	Wikipedia string   `json:"wikipedia,omitempty"`
}

type Societies struct {
	URL          string `json:"url,omitempty"`
	Organization string `json:"organization,omitempty"`
}

type DehydratedSource struct {
	RepositorySource
	IsCore   bool     `json:"is_core,omitempty"`
	IsInDOAJ bool     `json:"is_in_doaj,omitempty"`
	IsOA     bool     `json:"is_oa,omitempty"`
	ISSN     []string `json:"issn,omitempty"`
	ISSNL    string   `json:"issn_l,omitempty"`
	Type     string   `json:"type,omitempty"`
}

type RepositorySource struct {
	DisplayName             string   `json:"display_name,omitempty"`
	HostOrganization        string   `json:"host_organization,omitempty"`
	HostOrganizationLineage []string `json:"host_organization_lineage,omitempty"`
	HostOrganizationName    string   `json:"host_organization_name,omitempty"`
	ID                      string   `json:"id,omitempty"`
}
