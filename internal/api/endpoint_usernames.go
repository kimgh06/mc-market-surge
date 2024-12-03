package api

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"net/http"
	"strconv"
	"strings"
	"surge/internal/utilities"
)

func scanUsernameRows(rows *sql.Rows) ([]string, error) {
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	var items []string
	for rows.Next() {
		var i string
		if err := rows.Scan(&i); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

// EndpointUsernames resolves usernames from list of user id
func (a *SurgeAPI) EndpointUsernames(w http.ResponseWriter, r *http.Request) error {
	queryString := r.URL.Query().Get("id_list")
	split := strings.Split(queryString, ",")

	if split == nil || len(split) < 1 {
		return BadRequestError(ErrorCodeInvalidQuery, "invalid query id_list")
	}

	didItError := false
	splitCasted := utilities.Map(split, func(s string) string {
		casted, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			didItError = true
		}
		return strconv.FormatUint(casted, 10)
	})
	joined := strings.Join(splitCasted, ",")

	if didItError {
		return BadRequestError(ErrorCodeInvalidQuery, "invalid query id_list")
	}

	query := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).Select("username").From("auth.users").Where(fmt.Sprintf("id in (%s)", joined))

	rows, err := query.RunWith(a.db).Query()
	if err != nil {
		return InternalServerError("failed to query usernames: %+v", err)
	}

	results, err := scanUsernameRows(rows)
	if err != nil {
		return InternalServerError("failed to scan usernames from query result: %+v", err)
	}

	return writeResponseJSON(w, http.StatusOK, results)
}
