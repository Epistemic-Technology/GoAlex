package model

type Author struct {
	DehydratedAuthor
	Affiliations            []*AuthorAffiliation          `json:"affiliations,omitempty"`
	CitedByCount            int                           `json:"cited_by_count,omitempty"`
	CountsByYear            []*CountsByYear               `json:"counts_by_year,omitempty"`
	CreatedDate             string                        `json:"created_date,omitempty"`
	DisplayNameAlternatives []string                      `json:"display_name_alternatives,omitempty"`
	IDs                     *AuthorIDs                    `json:"ids,omitempty"`
	LastKnownInstitutions   []*DehydratedInstitution      `json:"last_known_institutions,omitempty"`
	SummaryStats            *SummaryStats                 `json:"summary_stats,omitempty"`
	Topics                  []*TopicWithCount             `json:"topics,omitempty"`
	TopicShare              []*TopicShare                 `json:"topic_share,omitempty"`
	UpdatedDate             string                        `json:"updated_date,omitempty"`
	WorksAPIURL             string                        `json:"works_api_url,omitempty"`
	WorksCount              int                           `json:"works_count,omitempty"`
	XConcepts               []*DehydratedConceptWithScore `json:"x_concepts,omitempty"`
}

type AuthorAffiliation struct {
	Institution *DehydratedInstitution `json:"institution,omitempty"`
	Years       []int                  `json:"years,omitempty"`
}

type AuthorIDs struct {
	OpenAlex  string `json:"openalex,omitempty"`
	ORCID     string `json:"orcid,omitempty"`
	Scopus    string `json:"scopus,omitempty"`
	Twitter   string `json:"twitter,omitempty"`
	Wikipedia string `json:"wikipedia,omitempty"`
}

type DehydratedAuthor struct {
	DisplayName string `json:"display_name,omitempty"`
	ID          string `json:"id,omitempty"`
	ORCID       string `json:"orcid,omitempty"`
}
