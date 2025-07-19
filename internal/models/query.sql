-- name: GetLatestSession :one
SELECT session_data FROM telegram_sessions 
ORDER BY created_at DESC 
LIMIT 1;

-- name: CreateSession :exec
INSERT INTO telegram_sessions (session_data, created_at)
VALUES ($1, $2);
