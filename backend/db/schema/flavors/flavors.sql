CREATE TABLE IF NOT EXISTS flavors (
    id TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    name INTEGER NOT NULL,
);