CREATE TABLE IF NOT EXISTS ingredients (
    id TEXT PRIMARY KEY,
    ingredient_type_id TEXT REFERENCES ingredient_types(id),
    flavor_id TEXT REFERENCES flavors(id),
    flavor_descriptor_id TEXT REFERENCES descriptors(id), -- this is for things like "Spiced"
    style_id TEXT REFERENCES styles(id), -- things like "Dark" or "Silver"
    brand_id TEXT REFERENCES brands(id),
    quality_id TEXT REFERENCES qualities(id),
    descriptive_name TEXT UNIQUE,
    image_url TEXT,
    is_alcoholic INTEGER NOT NULL,
    is_organic INTEGER NOT NULL,
    is_seasonal INTEGER NOT NULL
);