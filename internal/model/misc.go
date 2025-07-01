package model

type APC struct {
	Currency   string `json:"currency,omitempty"`
	Provenance string `json:"provenance,omitempty"`
	Value      int    `json:"value,omitempty"`
	ValueUSD   int    `json:"value_usd,omitempty"`
}

type APCPrice struct {
	Price    int    `json:"price,omitempty"`
	Currency string `json:"currency,omitempty"`
}

type Completion struct {
	ID           string `json:"id,omitempty"`
	ShortID      string `json:"short_id,omitempty"`
	DisplayName  string `json:"display_name,omitempty"`
	Hint         string `json:"hint,omitempty"`
	CitedByCount int    `json:"cited_by_count,omitempty"`
	WorksCount   int    `json:"works_count,omitempty"`
	EntityType   string `json:"entity_type,omitempty"`
	ExternalID   string `json:"external_id,omitempty"`
	FilterKey    string `json:"filter_key,omitempty"`
}

type CountsByYear struct {
	CitedByCount int `json:"cited_by_count,omitempty"`
	WorksCount   int `json:"works_count,omitempty"`
	Year         int `json:"year,omitempty"`
}

type International struct {
	DisplayName map[string]string `json:"display_name,omitempty"`
	Description map[string]string `json:"description,omitempty"`
}

type GEO struct {
	City           string  `json:"city,omitempty"`
	GeonamesCityID string  `json:"geonames_city_id,omitempty"`
	Region         string  `json:"region,omitempty"`
	CountryCode    string  `json:"country_code,omitempty"`
	Country        string  `json:"country,omitempty"`
	Latitude       float32 `json:"latitude,omitempty"`
	Longitude      float32 `json:"longitude,omitempty"`
}

type GroupBy struct {
	Key            string `json:"key,omitempty"`
	KeyDisplayName string `json:"key_display_name,omitempty"`
	Count          int    `json:"count,omitempty"`
}

type PaginatedResponseMeta struct {
	Count       int `json:"count,omitempty"`
	DBRespTime  int `json:"db_response_time_ms,omitempty"`
	Page        int `json:"page,omitempty"`
	PerPage     int `json:"per_page,omitempty"`
	GroupsCount int `json:"groups_count,omitempty"`
}

type Role struct {
	ID         string `json:"id,omitempty"`
	Role       string `json:"role,omitempty"`
	WorksCount int    `json:"works_count,omitempty"`
}

type SummaryStats struct {
	TwoYearMeanCitedness float64 `json:"2yr_mean_citedness,omitempty"`
	HIndex               int     `json:"h_index,omitempty"`
	I10Index             int     `json:"i10_index,omitempty"`
}
