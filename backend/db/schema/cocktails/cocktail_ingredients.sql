CREATE TABLE IF NOT EXISTS cocktail_ingredients (
    id TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    cocktail_id TEXT NOT NULL REFERENCES cocktails(id),
    ingredient_id TEXT NOT NULL REFERENCES ingredients(id),
    amount INTEGER,
    unit_of_measure TEXT,
    is_automatic BOOLEAN NOT NULL,
    category TEXT CHECK (category IN ('mixer', 'manual_ingredient', 'garnish'))
);