CREATE TABLE IF NOT EXISTS stations (
    id TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    position INTEGER NOT NULL,
    ingredient_id TEXT REFERENCES ingredients(id)
);