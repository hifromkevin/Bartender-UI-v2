CREATE TABLE IF NOT EXISTS cocktail_logs (
    id TEXT PRIMARY KEY,
    cocktail_id TEXT NOT NULL,
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (cocktail_id) REFERENCES cocktails(id)
);

CREATE INDEX IF NOT EXISTS idx_cocktail_logs_cocktail_id ON cocktail_logs(cocktail_id);
CREATE INDEX IF NOT EXISTS idx_cocktail_logs_timestamp ON cocktail_logs(timestamp);

-- Possibly use this instead: 
--id TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(4)) || '-' || hex(randomblob(2)) || '-' || hex(randomblob(2)) || '-' || hex(randomblob(2)) || '-' || hex(randomblob(6))))
-- It generates a more literally UUID style ID.