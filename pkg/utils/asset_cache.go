package utils

import "time"

type AssetCache struct {
	URL       string                 `json:"url"`
	Metadata  map[string]interface{} `json:"metadata"`
	UpdatedAt time.Time              `json:"updated_at"`
}
