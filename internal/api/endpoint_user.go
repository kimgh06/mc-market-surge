package api

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
)

// EndpointUser returns user data from logged in session
func (a *SurgeAPI) EndpointUser(w http.ResponseWriter, r *http.Request) error {
	claims := getClaims(r.Context())
	if claims == nil {
		return InternalServerError("failed to read claims")
	}

	if claims.Subject == "" {
		return BadRequestError(ErrorCodeBadJWT, "token subject is empty or undefined")
	}

	userId, err := strconv.ParseUint(claims.Subject, 10, 64)
	if err != nil {
		return BadRequestError(ErrorCodeBadJWT, "token subject is not a uuid")
	}

	user, err := a.queries.GetUser(r.Context(), int64(userId))
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return ForbiddenError(ErrorCodeUserNotFound, "token subject user does not exist")
		}
		return err
	}

	return writeResponseJSON(w, http.StatusOK, NewUserResponse(user))
}
