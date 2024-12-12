package api

import (
	"net/http"
	"surge/internal/utilities"
)

type ConfigurationResponse struct {
	Version        string `json:"version"`
	Name           string `json:"name"`
	SnowflakeStart string `json:"snowflake_start"`
}

var defaultVersion = "development build"

// EndpointConfiguration endpoint checks health of this service and returns information about public configuration
func (a *SurgeAPI) EndpointConfiguration(w http.ResponseWriter, r *http.Request) error {
	return writeResponseJSON(w, http.StatusOK, ConfigurationResponse{
		Version:        *utilities.Coalesce(a.version, &defaultVersion),
		Name:           "Surge API",
		SnowflakeStart: SnowflakeStartTime.String(),
	})
}
