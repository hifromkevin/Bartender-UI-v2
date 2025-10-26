CREATE TABLE IF NOT EXISTS stations (
    id TEXT PRIMARY KEY,
    position INTEGER NOT NULL,
    ingredient_id TEXT REFERENCES ingredients(id)
);