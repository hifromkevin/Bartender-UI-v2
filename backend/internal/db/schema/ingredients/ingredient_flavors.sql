CREATE TABLE IF NOT EXISTS ingredient_flavors (
    ingredient_id TEXT NOT NULL,
    flavor_id TEXT NOT NULL,
    PRIMARY KEY (ingredient_id, flavor_id),
    FOREIGN KEY (ingredient_id) REFERENCES ingredients(id),
    FOREIGN KEY (flavor_id) REFERENCES flavors(id)
);