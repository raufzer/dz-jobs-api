package response

type VersionResponse struct {
	APIVersion       int    `json:"api_version"`
	BuildVersion     string `json:"build_version"`
	CommitHash       string `json:"commit_hash"`
	ReleaseDate      string `json:"release_date"`
	Environment      string `json:"environment"`
	Status           string `json:"status"`
	DocumentationURL string `json:"documentation_url"`
	LastMigration    string `json:"last_migration"`
}

type HealthResponse struct {
	Status string `json:"status"`
}

type MetricsResponse struct {
	Uptime       string `json:"uptime"`
	RequestCount string `json:"request_count"`
	ErrorRate    string `json:"error_rate"`
}

