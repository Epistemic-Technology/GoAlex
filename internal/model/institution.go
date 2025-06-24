package model

type Institution struct {
	DehydratedInstitution
	AssociatedInstitutions  []*DehydratedInstitutionWithRelationship `json:"associated_institutions,omitempty"`
	CitedByCount            int                                      `json:"cited_by_count,omitempty"`
	CountsByYear            []*CountsByYear                          `json:"counts_by_year,omitempty"`
	CreatedDate             string                                   `json:"created_date,omitempty"`
	DisplayNameAcronyms     []string                                 `json:"display_name_acronyms,omitempty"`
	DisplayNameAlternatives []string                                 `json:"display_name_alternatives,omitempty"`
	GEO                     *GEO                                     `json:"geo,omitempty"`
	HomePageURL             string                                   `json:"homepage_url,omitempty"`
	IDs                     *InstitutionIDs                          `json:"ids,omitempty"`
	ImageThumbnailURL       string                                   `json:"image_thumbnail_url,omitempty"`
	ImageURL                string                                   `json:"image_url,omitempty"`
	International           *International                           `json:"international,omitempty"`
	IsSuperSystem           bool                                     `json:"is_super_system,omitempty"`
	Repositories            []*RepositorySource                      `json:"repositories,omitempty"`
	Roles                   []*Role                                  `json:"roles,omitempty"`
	SummaryStats            *SummaryStats                            `json:"summary_stats,omitempty"`
	Topics                  []*TopicWithCount                        `json:"topics,omitempty"`
	TopicShare              []*TopicShare                            `json:"topic_share,omitempty"`
	UpdatedDate             string                                   `json:"updated_date,omitempty"`
	WorksAPIURL             string                                   `json:"works_api_url,omitempty"`
	WorksCount              int                                      `json:"works_count,omitempty"`
	XConcepts               []*DehydratedConceptWithScore            `json:"x_concepts,omitempty"`
}

type InstitutionIDs struct {
	Grid      string `json:"grid,omitempty"`
	MAG       string `json:"mag,omitempty"`
	OpenAlex  string `json:"openalex,omitempty"`
	ROR       string `json:"ror,omitempty"`
	Wikipedia string `json:"wikipedia,omitempty"`
	Wikidata  string `json:"wikidata,omitempty"`
}

type DehydratedInstitution struct {
	CountryCode string   `json:"country_code,omitempty"`
	DisplayName string   `json:"display_name,omitempty"`
	ID          string   `json:"id,omitempty"`
	Lineage     []string `json:"lineage,omitempty"`
	ROR         string   `json:"ror,omitempty"`
	Type        string   `json:"type,omitempty"`
}

type DehydratedInstitutionWithRelationship struct {
	DehydratedInstitution
	Relationship string `json:"relationship,omitempty"`
}
