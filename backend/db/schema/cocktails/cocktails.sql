CREATE TABLE IF NOT EXISTS cocktails (
    id TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    name TEXT NOT NULL,
    description TEXT,
    is_alcoholic BOOLEAN NOT NULL,
    image_url TEXT,
    last_used_at TIMESTAMP
);