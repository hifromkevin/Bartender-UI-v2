CREATE TABLE IF NOT EXISTS cocktail_tools (
    cocktail_id TEXT NOT NULL,
    tool_id TEXT NOT NULL,
    PRIMARY KEY (cocktail_id, tool_id),
    FOREIGN KEY (cocktail_id) REFERENCES cocktails(id) ON DELETE CASCADE,
    FOREIGN KEY (tool_id) REFERENCES tools(id) ON DELETE CASCADE
);