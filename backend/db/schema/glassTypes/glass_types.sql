CREATE TABLE IF NOT EXISTS glass_types (
    id TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    name INTEGER NOT NULL,
    image TEXT,
);