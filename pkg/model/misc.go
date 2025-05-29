package model

type APC struct {
	Currency   string `json:"currency"`
	Provenance string `json:"provenance"`
	Value      int    `json:"value"`
	ValueUSD   int    `json:"value_usd"`
}

type CountsByYear struct {
	CitedByCount int `json:"cited_by_count"`
	WorksCount   int `json:"works_count"`
	Year         int `json:"year"`
}

type GEO struct {
	City           string  `json:"city"`
	GeonamesCityID int     `json:"geonames_city_id"`
	Region         string  `json:"region"`
	CountryCode    string  `json:"country_code"`
	Country        string  `json:"country"`
	Latitude       float32 `json:"latitude"`
	Longitude      float32 `json:"longitude"`
}

type Role struct {
	ID         string `json:"id"`
	Role       string `json:"role"`
	WorksCount int    `json:"works_count"`
}

type SummaryStats struct {
	TwoYearMeanCitedness float64 `json:"2yr_mean_citedness"`
	HIndex               int     `json:"h_index"`
	I10Index             int     `json:"i10_index"`
}
