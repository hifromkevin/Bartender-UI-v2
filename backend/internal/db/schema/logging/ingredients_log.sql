CREATE TABLE IF NOT EXISTS ingredients_log (
    id TEXT NOT NULL,
    ingredient_id TEXT NOT NULL,
    amount REAL,
    unit_of_measure TEXT,
    PRIMARY KEY (id, ingredient_id),
    FOREIGN KEY (id) REFERENCES cocktail_logs(id),
    FOREIGN KEY (ingredient_id) REFERENCES ingredients(id)
);

CREATE INDEX IF NOT EXISTS idx_ingredients_log_ingredient_id ON ingredients_log(ingredient_id);
CREATE INDEX IF NOT EXISTS idx_ingredients_log_id ON ingredients_log(id);