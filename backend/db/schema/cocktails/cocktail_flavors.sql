CREATE TABLE IF NOT EXISTS cocktail_flavors (
    cocktail_id TEXT NOT NULL,
    flavor_id TEXT NOT NULL,
    PRIMARY KEY (cocktail_id, flavor_id),
    FOREIGN KEY (cocktail_id) REFERENCES cocktails(id) ON DELETE CASCADE,
    FOREIGN KEY (flavor_id) REFERENCES flavors(id)
);