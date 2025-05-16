# Ingredient Enumerations

## DescriptiveName

The `DescriptiveName` field is populated as follows:

- If the user provides a value, it is used directly.
- If no value is provided:
  - If `Brand` is present, `DescriptiveName` is a combination of `Brand` + `IngredientType`.
  - If `Brand` is empty or `"null"`, `DescriptiveName` is just `IngredientType`.

### Example Scenarios

1. **User-Provided Value**:

   - `Brand`: `"Jose Cuervo"`
   - `IngredientType`: `"Tequila"`
   - `DescriptiveName`: `"Jose Cuervo Especial Gold Tequila"` (user-provided)

2. **Generated Value with Brand**:

   - `Brand`: `"Smirnoff"`
   - `IngredientType`: `"Vodka"`
   - `DescriptiveName`: `"Smirnoff Vodka"`

3. **Generated Value without Brand**:
   - `Brand`: `""`
   - `IngredientType`: `"Vodka"`
   - `DescriptiveName`: `"Vodka"`

## Quality Tiers

- **null**: Not Applicable (e.g., for certain `IngredientClassID` values)
- **1**: Well – Base Level
- **2**: Call – Mid Level (Name Brand)
- **3**: Top Shelf – High End
- **4**: Ultra - Luxury Tier
- **5**: Exclusive - Exclusive Luxury Tier

## Seasons

- **null**: Not Applicable
- **Winter**
- **Spring**
- **Summer**
- **Fall**
