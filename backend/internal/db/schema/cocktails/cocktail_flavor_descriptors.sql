CREATE TABLE IF NOT EXISTS cocktail_flavor_descriptors (
    cocktail_id TEXT NOT NULL,
    flavor_descriptor_id TEXT NOT NULL,
    PRIMARY KEY (cocktail_id, flavor_descriptor_id),
    FOREIGN KEY (cocktail_id) REFERENCES cocktails(id) ON DELETE CASCADE,
    FOREIGN KEY (flavor_descriptor_id) REFERENCES flavor_descriptors(id)
);