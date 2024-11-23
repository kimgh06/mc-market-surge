package api

import (
	"context"
	"net/http"
)

func (a *SurgeAPI) requireServiceKey(w http.ResponseWriter, r *http.Request) (context.Context, error) {
	serviceKey := r.Header.Get("Surge-Service-Key")

	if a.config.ServiceKey.Value != serviceKey {
		return r.Context(), UnauthorizedError(ErrorCodeInvalidServiceKey, "This endpoint requires service key to access")
	}

	return r.Context(), nil
}
