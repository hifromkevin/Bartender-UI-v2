CREATE TABLE IF NOT EXISTS ingredient_groups (
    id TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    name TEXT NOT NULL UNIQUE,
);