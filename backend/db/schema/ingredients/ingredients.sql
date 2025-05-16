CREATE TABLE IF NOT EXISTS ingredients (
    id TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    ingredient_type_id TEXT,
    flavor_id TEXT, // todo: make this an array for things like "Spiced Pear", 
    brand TEXT,
    descriptive_name TEXT,
    image_url TEXT,
    is_alcoholic BOOLEAN NOT NULL,
    is_organic BOOLEAN NOT NULL,
    is_seasonal BOOLEAN NOT NULL,
    quality_tier INTEGER CHECK (quality_tier IN (null, 1, 2, 3, 4, 5)),
    season TEXT CHECK (season IN ('Summer', 'Spring', 'Fall', 'Winter', NULL))
);

CREATE TRIGGER set_descriptive_name
AFTER INSERT ON ingredients
FOR EACH ROW
BEGIN
    UPDATE ingredients
    SET descriptive_name = 
        CASE 
            WHEN NEW.brand IS NOT NULL AND NEW.flavor_id IS NOT NULL THEN 
                NEW.brand || ' ' || (SELECT name FROM flavors WHERE id = NEW.flavor_id) || ' ' || (SELECT name FROM ingredient_types WHERE id = NEW.ingredient_type_id)
            WHEN NEW.brand IS NOT NULL THEN 
                NEW.brand || ' ' || (SELECT name FROM ingredient_types WHERE id = NEW.ingredient_type_id)
            WHEN NEW.flavor_id IS NOT NULL THEN 
                (SELECT name FROM flavors WHERE id = NEW.flavor_id) || ' ' || (SELECT name FROM ingredient_types WHERE id = NEW.ingredient_type_id)
            ELSE 
                (SELECT name FROM ingredient_types WHERE id = NEW.ingredient_type_id)
        END
    WHERE id = NEW.id
    AND (NEW.descriptive_name IS NULL OR NEW.descriptive_name = '');
END;