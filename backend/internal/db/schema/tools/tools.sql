CREATE TABLE IF NOT EXISTS tools (
    id TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    name TEXT NOT NULL,
    image_url TEXT
);