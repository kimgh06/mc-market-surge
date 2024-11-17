package api

import (
	"database/sql"
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// EndpointUsername checks if username is used already
func (a *SurgeAPI) EndpointUsername(w http.ResponseWriter, r *http.Request) error {
	username := chi.URLParam(r, "username")

	if username == "" {
		return BadRequestError(ErrorCodeInvalidUsername, "invalid username %+v", username)
	}

	_, err := a.queries.GetUserByUsername(r.Context(), username)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return writeResponseJSON(w, http.StatusOK, false)
		}
		return err
	}

	return writeResponseJSON(w, http.StatusOK, true)
}
