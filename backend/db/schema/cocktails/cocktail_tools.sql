CREATE TABLE IF NOT EXISTS cocktail_tools (
    id TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    name TEXT NOT NULL,
    type TEXT,
    image_url TEXT
);