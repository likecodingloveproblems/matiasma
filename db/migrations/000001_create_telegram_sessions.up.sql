CREATE TABLE telegram_sessions (
    id SERIAL PRIMARY KEY,
    session_data BYTEA NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
