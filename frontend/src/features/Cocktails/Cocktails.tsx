import {
  chocolatini,
  ginAndTonic,
  moscowMule,
  oldFashioned,
  rumAndCoke,
  shirleyTemple,
  vodkaCran,
  vodkaSoda,
} from '../../assets/exportCocktailIcons';
import { info } from '../../assets/exportIcons';

import styles from './styles/Cocktails.module.scss';

const cocktailImages: Record<string, string> = {
  'Moscow Mule': moscowMule,
  Chocolatini: chocolatini,
  'Gin and Tonic': ginAndTonic,
  'Old Fashioned': oldFashioned,
  'Rum and Coke': rumAndCoke,
  'Shirley Temple': shirleyTemple,
  'Vodka Cran': vodkaCran,
  'Vodka Soda': vodkaSoda,
};

const getCocktailImage = (cocktailName: string): string =>
  cocktailImages[cocktailName];

const fakeCocktails = [
  {
    id: 1,
    name: 'Moscow Mule',
    description:
      'A classic cocktail made with whiskey, lemon juice, and sugar.',
    glass_type: 'Cocktail',
    ice_type: 'Cubed',
    requires_shaker: true,
    instructions_id: 1,
    difficulty: 'Medium',
  },
  {
    id: 2,
    name: 'Chocolatini',
    description:
      'A rich, luxurious cocktail made with vodka, chocolate liqueur, and cream.',
    glass_type: 'Martini',
    ice_type: null,
    requires_shaker: false,
    instructions_id: 2,
    difficulty: 'Medium',
  },
  {
    id: 3,
    name: 'Gin and Tonic',
    description: 'A refreshing cocktail made with gin, tonic water, and lime.',
    glass_type: 'Cocktail',
    ice_type: 'Cubed',
    requires_shaker: false,
    instructions_id: 1,
    difficulty: 'Easy',
  },
  {
    id: 4,
    name: 'Old Fashioned',
    description:
      'A classic cocktail made with whiskey, bitters, sugar, and water.',
    glass_type: 'Old Fashioned',
    ice_type: 'Cubed',
    requires_shaker: false,
    instructions_id: 2,
    difficulty: 'Hard',
  },
  {
    id: 5,
    name: 'Rum and Coke',
    description: 'A simple cocktail made with rum and cola, served over ice.',
    glass_type: 'Highball',
    ice_type: 'Cubed',
    requires_shaker: false,
    instructions_id: 1,
    difficulty: 'Easy',
  },
  {
    id: 6,
    name: 'Shirley Temple',
    description:
      'A non-alcoholic cocktail made with ginger ale, grenadine, and a splash of orange juice.',
    glass_type: 'Highball',
    ice_type: 'Cubed',
    requires_shaker: false,
    instructions_id: 2,
    difficulty: 'Medium',
  },
  {
    id: 7,
    name: 'Vodka Cran',
    description: 'A refreshing cocktail made with vodka and cranberry juice.',
    glass_type: 'Highball',
    ice_type: 'Cubed',
    requires_shaker: false,
    instructions_id: 1,
    difficulty: 'Easy',
  },
  {
    id: 8,
    name: 'Vodka Soda',
    description:
      'A simple cocktail made with vodka and soda water, served over ice.',
    glass_type: 'Highball',
    ice_type: 'Cubed',
    requires_shaker: false,
    instructions_id: 2,
    difficulty: 'Easy',
  },
];

const setDifficultyColor = (difficulty: string): string => {
  switch (difficulty) {
    case 'Easy':
      return styles.easy;
    case 'Medium':
      return styles.medium;
    case 'Hard':
      return styles.hard;
    default:
      return '';
  }
};

const Cocktails = () => {
  return (
    <section className={styles.cocktails}>
      <div className={styles.cocktailsHeader}>
        <h3>Cocktails </h3>
        <img
          src={info}
          alt="Information about cocktails"
          className={styles.infoIcon}
        />
      </div>
      <section className={styles.cocktailsList}>
        {fakeCocktails.map((cocktail) => (
          <div key={cocktail.id} className={styles.cocktailItem}>
            <img
              src={getCocktailImage(cocktail.name)}
              alt={cocktail.name}
              className={styles.cocktailImage}
            />
            <h4 className={styles.cocktailName}>
              <span
                className={`${styles.difficulty} ${setDifficultyColor(cocktail.difficulty)}`}
              ></span>
              {cocktail.name}
            </h4>
            {/* <h4
              className={`${styles.cocktailName} ${setDifficultyColor(cocktail.difficulty)}`}
            >
              <span
                className={`${styles.difficulty} ${setDifficultyColor(cocktail.difficulty)}`}
              ></span>
              {cocktail.name}
            </h4> */}
          </div>
        ))}
      </section>
    </section>
  );
};

export default Cocktails;
