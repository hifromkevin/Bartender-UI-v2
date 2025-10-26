CREATE TABLE IF NOT EXISTS cocktails (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    is_alcoholic BOOLEAN NOT NULL,
    image_url TEXT,
    requires_shaker BOOLEAN NOT NULL,
    glass_id TEXT,
    FOREIGN KEY (glass_id) REFERENCES glass_types(id)
);