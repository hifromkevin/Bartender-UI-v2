CREATE TABLE IF NOT EXISTS cocktail_usage (
    id TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    cocktail_id TEXT NOT NULL REFERENCES cocktails(id),
    made_at TIMESTAMP NOT NULL
);