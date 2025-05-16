CREATE TABLE IF NOT EXISTS ingredient_types (
    id TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    name TEXT NOT NULL UNIQUE,
    group_id TEXT NOT NULL REFERENCES ingredient_groups(id),
);