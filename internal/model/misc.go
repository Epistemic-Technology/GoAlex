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
