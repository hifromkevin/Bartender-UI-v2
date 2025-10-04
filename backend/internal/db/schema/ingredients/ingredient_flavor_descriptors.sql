CREATE TABLE IF NOT EXISTS ingredient_flavor_descriptors (
    ingredient_id TEXT NOT NULL,
    flavor_descriptor_id TEXT NOT NULL,
    PRIMARY KEY (ingredient_id, flavor_descriptor_id),
    FOREIGN KEY (ingredient_id) REFERENCES ingredients(id),
    FOREIGN KEY (flavor_descriptor_id) REFERENCES flavor_descriptors(id)
);