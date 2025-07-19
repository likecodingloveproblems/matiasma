-- name: GetUserLatestSession :one
SELECT session_data, created_at
FROM telegram_sessions
WHERE phone_number = $1
ORDER BY created_at DESC
LIMIT 1;

-- name: UpsertSession :exec
INSERT INTO telegram_sessions (phone_number, session_data, created_at)
VALUES ($1, $2, $3)
ON CONFLICT (phone_number)
    DO UPDATE SET
    session_data = EXCLUDED.session_data;
