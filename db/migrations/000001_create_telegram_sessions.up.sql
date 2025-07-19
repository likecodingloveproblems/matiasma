CREATE TABLE telegram_sessions (
    phone_number VARCHAR PRIMARY KEY,
    session_data BYTEA NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
