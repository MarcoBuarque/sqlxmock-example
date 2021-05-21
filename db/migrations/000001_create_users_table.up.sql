CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY NOT NULL,
    first_name text NOT NULL,
    last_name text NOT NULL,
    email text,
    country text,
    created_at TIMESTAMP DEFAULT NOW()
);