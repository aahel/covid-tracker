package types

import "time"

type CovidStat struct {
	Meta struct {
		LastUpdated time.Time `json:"last_updated"`
	} `json:"meta"`
	Total struct {
		Confirmed int `json:"confirmed"`
	} `json:"total"`
}

const (
	Latitude  = "latitude"
	Longitude = "longitude"
)

type CovidStatResp struct {
	StateCode     string    `json:"state_code,omitempty"`
	StateCount    int       `json:"state_count,omitempty"`
	LastUpdatedAt time.Time `json:"last_updated_at,omitempty"`
	IndiaCount    int64     `json:"india_count,omitempty"`
}

type ReverseGeoCodingResp struct {
	Data []*ReverseGeoCoding `json:"data,omitempty"`
}

type ReverseGeoCoding struct {
	RegionCode string `json:"region_code,omitempty"`
}

type TotalCases struct {
	TotalCases int64 `bson:"total_cases"`
}
