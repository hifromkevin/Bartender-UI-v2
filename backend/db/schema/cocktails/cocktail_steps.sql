CREATE TABLE IF NOT EXISTS cocktail_steps (
    id TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    cocktail_id TEXT NOT NULL REFERENCES cocktails(id),
    step_number INTEGER NOT NULL,
    instruction TEXT NOT NULL
);
