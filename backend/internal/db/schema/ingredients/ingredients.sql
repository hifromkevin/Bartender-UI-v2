CREATE TABLE IF NOT EXISTS ingredients (
    id TEXT PRIMARY KEY,
    ingredient_type_id TEXT, 
    flavor_id TEXT,
    flavor_descriptor_id TEXT, -- this is for things like "Spiced"
    style_id TEXT, -- things like "Dark" or "Silver"
    brand_id TEXT, 
    quality_id TEXT,
    descriptive_name TEXT UNIQUE,
    image_url TEXT,
    is_alcoholic BOOLEAN NOT NULL,
    is_organic BOOLEAN NOT NULL,
    is_seasonal BOOLEAN NOT NULL,
);