package model

type Affiliation struct {
	InstitutionIDs       []string `json:"institution_ids,omitempty"`
	RawAffiliationString string   `json:"raw_affiliation_string,omitempty"`
}

type Authorship struct {
	Affiliations          []*Affiliation           `json:"affiliations,omitempty"`
	Author                DehydratedAuthor         `json:"author,omitempty"`
	AuthorPosition        string                   `json:"author_position,omitempty"`
	Countries             []string                 `json:"countries,omitempty"`
	Institution           []*DehydratedInstitution `json:"institutions,omitempty"`
	IsCorresponding       bool                     `json:"is_corresponding,omitempty"`
	RawAffiliationStrings []string                 `json:"raw_affiliation_strings,omitempty"`
	RawAuthorName         string                   `json:"raw_author_name,omitempty"`
}

type Biblio struct {
	Volume    string `json:"volume,omitempty"`
	Issue     string `json:"issue,omitempty"`
	FirstPage string `json:"first_page,omitempty"`
	LastPage  string `json:"last_page,omitempty"`
}

type CitationNormalizedPercentile struct {
	IsTop1Percent  bool    `json:"is_in_top_1_percent,omitempty"`
	IsTop10Percent bool    `json:"is_in_top_10_percent,omitempty"`
	Value          float64 `json:"value,omitempty"`
}

type Grants struct {
	AwardID           string `json:"award_id,omitempty"`
	Funder            string `json:"funder,omitempty"`
	FunderDisplayName string `json:"funder_display_name,omitempty"`
}

type IDs struct {
	DOI      string `json:"doi,omitempty"`
	MAG      string `json:"mag,omitempty"`
	OpenAlex string `json:"openalex,omitempty"`
	PMCID    string `json:"pmcid,omitempty"`
	PMID     string `json:"pmid,omitempty"`
}

type Location struct {
	IsAccepted     bool              `json:"is_accepted,omitempty"`
	IsOA           bool              `json:"is_oa,omitempty"`
	IsPublished    bool              `json:"is_published,omitempty"`
	LandingPageURL string            `json:"landing_page_url,omitempty"`
	License        string            `json:"license,omitempty"`
	PDF_URL        string            `json:"pdf_url,omitempty"`
	Source         *DehydratedSource `json:"source,omitempty"`
	Version        string            `json:"version,omitempty"`
}

type MeSH struct {
	DescriptorUI   string `json:"descriptor_ui,omitempty"`
	DescriptorName string `json:"descriptor_name,omitempty"`
	QualifierUI    string `json:"qualifier_ui,omitempty"`
	QualifierName  string `json:"qualifier_name,omitempty"`
	IsMajorTopic   bool   `json:"is_major_topic,omitempty"`
}

type OAStatus string

const (
	OAStatusClosed  OAStatus = "closed"
	OAStatusBronze  OAStatus = "bronze"
	OAStatusDiamond OAStatus = "diamond"
	OAStatusGold    OAStatus = "gold"
	OAStatusGreen   OAStatus = "green"
	OAStatusHybrid  OAStatus = "hybrid"
)

type OpenAccess struct {
	AnyRepoHasFulltext bool      `json:"any_repository_has_fulltext,omitempty"`
	IsOA               bool      `json:"is_oa,omitempty"`
	OAStatus           *OAStatus `json:"oa_status,omitempty"`
	OAURL              string    `json:"oa_url,omitempty"`
}

type SDGs struct {
	DisplayName string  `json:"display_name,omitempty"`
	ID          string  `json:"id,omitempty"`
	Score       float64 `json:"score,omitempty"`
}

type Work struct {
	Abstract                     string                        `json:"abstract,omitempty"`
	AbstractInvertedIndex        map[string][]int              `json:"abstract_inverted_index,omitempty"`
	Authorships                  []*Authorship                 `json:"authorships,omitempty"`
	APCList                      *APC                          `json:"apc_list,omitempty"`
	APCPaid                      *APC                          `json:"apc_paid,omitempty"`
	BestOALocation               *Location                     `json:"best_oa_location,omitempty"`
	Biblio                       *Biblio                       `json:"biblio,omitempty"`
	CitationNormalizedPercentile *CitationNormalizedPercentile `json:"citation_normalized_percentile,omitempty"`
	CitedByAPIURL                string                        `json:"cited_by_api_url,omitempty"`
	CitedByCount                 int                           `json:"cited_by_count,omitempty"`
	Concepts                     []*DehydratedConceptWithScore `json:"concepts,omitempty"`
	CorrespondingAuthorIDs       []string                      `json:"corresponding_author_ids,omitempty"`
	CorrespondingInstitutionIDs  []string                      `json:"corresponding_institution_ids,omitempty"`
	CountriesDistinctCount       int                           `json:"countries_distinct_count,omitempty"`
	CountsByYear                 []*CountsByYear               `json:"counts_by_year,omitempty"`
	CreatedDate                  string                        `json:"created_date,omitempty"`
	DisplayName                  string                        `json:"display_name,omitempty"`
	DOI                          string                        `json:"doi,omitempty"`
	FulltextOrigin               string                        `json:"fulltext_origin,omitempty"`
	FWCI                         float32                       `json:"fwci,omitempty"`
	Grants                       []*Grants                     `json:"grants,omitempty"`
	HasFulltext                  bool                          `json:"has_fulltext,omitempty"`
	ID                           string                        `json:"id,omitempty"`
	IDs                          *IDs                          `json:"ids,omitempty"`
	IndexedIn                    []string                      `json:"indexed_in,omitempty"`
	InstitutionsDistinctCount    int                           `json:"institutions_distinct_count,omitempty"`
	IsParatext                   bool                          `json:"is_paratext,omitempty"`
	IsRetracted                  bool                          `json:"is_retracted,omitempty"`
	Keywords                     []*DehydratedKeyword          `json:"keywords,omitempty"`
	Language                     string                        `json:"language,omitempty"`
	License                      string                        `json:"license,omitempty"`
	Locations                    []*Location                   `json:"locations,omitempty"`
	LocationCount                int                           `json:"location_count,omitempty"`
	MeSH                         []*MeSH                       `json:"mesh,omitempty"`
	OpenAccess                   *OpenAccess                   `json:"open_access,omitempty"`
	PrimaryLocation              *Location                     `json:"primary_location,omitempty"`
	PrimaryTopic                 *TopicWithScore               `json:"primary_topic,omitempty"`
	PublicationDate              string                        `json:"publication_date,omitempty"`
	PublicationYear              int                           `json:"publication_year,omitempty"`
	ReferenceWorks               []string                      `json:"reference_works,omitempty"`
	RelatedWorks                 []string                      `json:"related_works,omitempty"`
	SDGs                         []*SDGs                       `json:"sustainable_development_goals,omitempty"`
	Title                        string                        `json:"title,omitempty"`
	Topics                       []*TopicWithScore             `json:"topics,omitempty"`
	Type                         string                        `json:"type,omitempty"`
	TypeCrossref                 string                        `json:"type_crossref,omitempty"`
	UpdatedDate                  string                        `json:"updated_date,omitempty"`
}
