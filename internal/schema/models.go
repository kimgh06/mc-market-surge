// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package schema

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type AuthRefreshToken struct {
	ID        int64
	UserID    uuid.NullUUID
	Token     sql.NullString
	Revoked   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AuthUser struct {
	ID                uuid.UUID
	Email             sql.NullString
	Username          sql.NullString
	EncryptedPassword sql.NullString
	CreatedAt         time.Time
	UpdatedAt         time.Time
	LastSignIn        sql.NullTime
}
