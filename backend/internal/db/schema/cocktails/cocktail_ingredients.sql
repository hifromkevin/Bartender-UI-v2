CREATE TABLE IF NOT EXISTS cocktail_ingredients (
    id TEXT PRIMARY KEY,
    cocktail_id TEXT NOT NULL,
    ingredient_id TEXT NOT NULL,
    amount REAL,
    unit_of_measure TEXT,
    is_automatic BOOLEAN NOT NULL DEFAULT 0,
    category TEXT CHECK (category IN ('mixer', 'manual_ingredient', 'garnish')),
    has_auto_option BOOLEAN NOT NULL DEFAULT 0,
    alternative_group TEXT,
    FOREIGN KEY (cocktail_id) REFERENCES cocktails(id) ON DELETE CASCADE,
    FOREIGN KEY (ingredient_id) REFERENCES ingredients(id),
    UNIQUE (cocktail_id, ingredient_id, alternative_group) -- ensures no duplicate ingredients in same alternative group
);
-- TODO: No bools! Use INTEGER 0/1 instead