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
	Status           string `json:"status"`
	Database         string `json:"database"`
	Cache            string `json:"cache"`
	ExternalServices string `json:"external_services"`
}

type MetricsResponse struct {
	Uptime       string `json:"uptime"`
	RequestCount string `json:"request_count"`
	ErrorRate    string `json:"error_rate"`
}

func AggregateHealthStatus(status map[string]string) string {
	for _, serviceStatus := range status {
		if serviceStatus == "unhealthy" {
			return "unhealthy"
		}
	}
	return "healthy"
}
