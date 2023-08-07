-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, cf_id, verification_status, email, hashed_password)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;
