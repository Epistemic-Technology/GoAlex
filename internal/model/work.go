package model

type Affiliation struct {
	InstitutionIDs       []string `json:"institution_ids"`
	RawAffiliationString string   `json:"raw_affiliation_string"`
}

type Authorship struct {
	Affiliations          []Affiliation           `json:"affiliations"`
	Author                DehydratedAuthor        `json:"author"`
	AuthorPosition        string                  `json:"author_position"`
	Corresponding         bool                    `json:"is_corresponding"`
	Countries             []string                `json:"countries"`
	Institution           []DehydratedInstitution `json:"institutions"`
	RawAffiliationStrings []string                `json:"raw_affiliation_strings"`
	RawAuthorName         string                  `json:"raw_author_name"`
}

type IDs struct {
	DOI      string `json:"doi"`
	MAG      string `json:"mag"`
	OpenAlex string `json:"openalex"`
	PMCID    string `json:"pmcid"`
	PMID     string `json:"pmid"`
}

type Location struct {
	Accepted       bool             `json:"is_accepted"`
	LandingPageURL string           `json:"landing_page_url"`
	License        string           `json:"license"`
	OA             bool             `json:"is_oa"`
	PDF_URL        string           `json:"pdf_url"`
	Published      bool             `json:"is_published"`
	Source         DehydratedSource `json:"source"`
	Version        string           `json:"version"`
}

type MeSH struct {
	DescriptorUI   string `json:"descriptor_ui"`
	DescriptorName string `json:"descriptor_name"`
	QualifierUI    string `json:"qualifier_ui"`
	QualifierName  string `json:"qualifier_name"`
	MajorTopic     bool   `json:"is_major_topic"`
}

type OpenAccess struct {
	AnyRepositoryHasFulltext bool   `json:"any_repository_has_fulltext"`
	OA                       bool   `json:"is_oa"`
	OAStatus                 string `json:"oa_status"`
	OAURL                    string `json:"oa_url"`
}

type Work struct {
	AbstractInvertedIndex map[string][]int `json:"abstract_inverted_index"`
	Authorships           []Authorship     `json:"authorships"`
	APCList               APC              `json:"apc_list"`
	APCPaid               APC              `json:"apc_paid"`
	BestOALocation        Location         `json:"best_oa_location"`
	Biblio                struct {
		Volume    string `json:"volume"`
		Issue     string `json:"issue"`
		FirstPage string `json:"first_page"`
		LastPage  string `json:"last_page"`
	} `json:"biblio"`
	CitationNormalizedPercentile struct {
		Value        float64 `json:"value"`
		Top1Percent  bool    `json:"is_in_top_1_percent"`
		Top10Percent bool    `json:"is_in_top_10_percent"`
	} `json:"citation_normalized_percentile"`
	CitedByAPIURL               string                       `json:"cited_by_api_url"`
	CitedByCount                int                          `json:"cited_by_count"`
	Concepts                    []DehydratedConceptWithScore `json:"concepts"`
	CorrespondingAuthorIDs      []string                     `json:"corresponding_author_ids"`
	CorrespondingInstitutionIDs []string                     `json:"corresponding_institution_ids"`
	CountriesDistinctCount      int                          `json:"countries_distinct_count"`
	CountsByYear                []struct {
		CitedByCount int `json:"cited_by_count"`
		Year         int `json:"year"`
	} `json:"counts_by_year"`
	CreatedDate    string  `json:"created_date"`
	DisplayName    string  `json:"display_name"`
	DOI            string  `json:"doi"`
	Fulltext       bool    `json:"has_fulltext"`
	FulltextOrigin string  `json:"fulltext_origin"`
	FWCI           float32 `json:"fwci"`
	Grants         []struct {
		AwardID           string `json:"award_id"`
		Funder            string `json:"funder"`
		FunderDisplayName string `json:"funder_display_name"`
	} `json:"grants"`
	ID                        string              `json:"id"`
	IDs                       IDs                 `json:"ids"`
	IndexedIn                 []string            `json:"indexed_in"`
	InstitutionsDistinctCount int                 `json:"institutions_distinct_count"`
	Keywords                  []DehydratedKeyword `json:"keywords"`
	Language                  string              `json:"language"`
	License                   string              `json:"license"`
	Locations                 []Location          `json:"locations"`
	LocationCount             int                 `json:"location_count"`
	MeSH                      []MeSH              `json:"mesh"`
	OpenAccess                OpenAccess          `json:"open_access"`
	Paratext                  bool                `json:"is_paratext"`
	PrimaryLocation           Location            `json:"primary_location"`
	PrimaryTopic              TopicWithScore      `json:"primary_topic"`
	PublicationDate           string              `json:"publication_date"`
	PublicationYear           int                 `json:"publication_year"`
	ReferenceWorks            []string            `json:"reference_works"`
	RelatedWorks              []string            `json:"related_works"`
	Retracted                 bool                `json:"is_retracted"`
	SDGs                      []struct {
		DisplayName string  `json:"display_name"`
		ID          string  `json:"id"`
		Score       float64 `json:"score"`
	} `json:"sustainable_development_goals"`
	Title        string           `json:"title"`
	Topics       []TopicWithScore `json:"topics"`
	Type         string           `json:"type"`
	TypeCrossref string           `json:"type_crossref"`
	UpdatedDate  string           `json:"updated_date"`
}
