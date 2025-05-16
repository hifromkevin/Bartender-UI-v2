CREATE TABLE IF NOT EXISTS ingredient_usage (
    id TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    ingredient_id TEXT NOT NULL REFERENCES ingredients(id),
    ingredient_amount INTEGER,
    unit_of_measure TEXT,
    used_at TIMESTAMP NOT NULL
);