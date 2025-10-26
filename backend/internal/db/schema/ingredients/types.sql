CREATE TABLE IF NOT EXISTS ingredient_types (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    group_id TEXT NOT NULL REFERENCES ingredient_groups(id)
);