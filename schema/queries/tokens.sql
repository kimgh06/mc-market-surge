-- name: CreateRefreshToken :one
insert into auth.refresh_tokens(id, user_id, token, revoked, created_at, updated_at)
values ($1, $2, $3, $4, now(), now())
returning *;

-- name: ListRefreshTokenByUser :many
select *
from auth.refresh_tokens
where user_id = $1;

-- name: GetRefreshToken :one
select *
from auth.refresh_tokens
where token = sqlc.arg('token')::varchar
  and revoked = false;

-- name: RevokeRefreshToken :exec
update auth.refresh_tokens
set revoked = true
where id = $1;

-- name: RevokeRefreshTokensOfUser :exec
update auth.refresh_tokens
set revoked = true
where user_id = $1;