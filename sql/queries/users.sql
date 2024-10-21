-- name: CreateUser :one
INSERT INTO
    "user" (id, created_at, updated_at, name,api_key)
VALUES
    ($1,$2,$3,$4,
        encode(sha256(random()::text::bytea), 'hex')
    )
RETURNING *;

-- name: GetUserByAPIKey :one
SELECT * FROM "user" WHERE api_key=$1;