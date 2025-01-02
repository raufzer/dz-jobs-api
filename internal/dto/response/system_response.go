package response

type DefaultResponse struct {
	Message       string `json:"message"`
	Documentation string `json:"documentation"`
	Version       string `json:"version"`
	Health        string `json:"health"`
	Metrics       string `json:"metrics"`
}
type VersionResponse struct {
	APIVersion       int    `json:"api_version"`
	BuildVersion     string `json:"build_version"`
	CommitHash       string `json:"commit_hash"`
	ReleaseDate      string `json:"release_date"`
	Environment      string `json:"environment"`
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

