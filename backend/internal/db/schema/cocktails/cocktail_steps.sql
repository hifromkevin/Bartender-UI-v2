CREATE TABLE IF NOT EXISTS cocktail_steps (
    id TEXT PRIMARY KEY,
    cocktail_id TEXT NOT NULL,
    step_number INTEGER NOT NULL,
    instruction TEXT NOT NULL,
    FOREIGN KEY (cocktail_id) REFERENCES cocktails(id) ON DELETE CASCADE,
    UNIQUE(cocktail_id, step_number)
);

CREATE INDEX IF NOT EXISTS idx_cocktail_steps_cocktail_id ON cocktail_steps(cocktail_id);

-- TODO: Consider making steps like Philz's menu:
 -- EX: Pre-made "Insert shaker into bartender" or "Add ice to shaker"
 -- MadLibs Style: "Add {amount} {unit_of_measure} of {ingredient_name}"