CREATE TABLE IF NOT EXISTS ingredient_styles (
    id TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    name TEXT NOT NULL UNIQUE
);